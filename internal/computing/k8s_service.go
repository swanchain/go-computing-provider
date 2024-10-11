package computing

import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	calicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	calicoclientset "github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/retry"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	appV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	networkingv1 "k8s.io/api/networking/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientSet *kubernetes.Clientset
var k8sOnce sync.Once
var config *rest.Config
var version string

type K8sService struct {
	k8sClient *kubernetes.Clientset
	Version   string
	config    *rest.Config
}

func NewK8sService() *K8sService {
	var err error
	k8sOnce.Do(func() {
		kubeConfig := filepath.Join(homedir.HomeDir(), ".kube/config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return
		}
		config.QPS = 30
		config.Burst = 200
		clientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			logs.GetLogger().Errorf("Failed create k8s clientset, error: %v", err)
			return
		}

		versionInfo, err := clientSet.Discovery().ServerVersion()
		if err != nil {
			return
		}
		version = versionInfo.String()
	})

	return &K8sService{
		k8sClient: clientSet,
		Version:   version,
		config:    config,
	}
}

func (s *K8sService) CreateDeployment(ctx context.Context, nameSpace string, deploy *appV1.Deployment) (result *appV1.Deployment, err error) {
	return s.k8sClient.AppsV1().Deployments(nameSpace).Create(ctx, deploy, metaV1.CreateOptions{})
}

func (s *K8sService) DeleteDeployment(ctx context.Context, namespace, deploymentName string) error {
	return s.k8sClient.AppsV1().Deployments(namespace).Delete(ctx, deploymentName, metaV1.DeleteOptions{})
}

func (s *K8sService) DeletePod(ctx context.Context, namespace, spaceUuid string) error {
	return s.k8sClient.CoreV1().Pods(namespace).DeleteCollection(ctx, *metaV1.NewDeleteOptions(0), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("lad_app=%s", spaceUuid),
	})
}

func (s *K8sService) DeleteDeployRs(ctx context.Context, namespace, jobUuid string) error {
	return s.k8sClient.AppsV1().ReplicaSets(namespace).DeleteCollection(ctx, *metaV1.NewDeleteOptions(0), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("lad_app=%s", jobUuid),
	})
}

func (s *K8sService) GetDeploymentStatus(namespace, spaceUuid string) (string, error) {
	namespace = constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(namespace)
	podList, err := s.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("lad_app=%s", spaceUuid),
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	if len(podList.Items) > 0 {
		return string(podList.Items[0].Status.Phase), nil
	}
	return "", nil
}

func (s *K8sService) GetDeploymentImages(ctx context.Context, namespace, deploymentName string) ([]string, error) {
	deployment, err := s.k8sClient.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	var imageIds []string
	for _, container := range deployment.Spec.Template.Spec.Containers {
		imageIds = append(imageIds, container.Image)
	}
	return imageIds, nil
}

func (s *K8sService) GetServiceByName(ctx context.Context, namespace, serviceName string, opts metaV1.GetOptions) (result *coreV1.Service, err error) {
	return s.k8sClient.CoreV1().Services(namespace).Get(ctx, serviceName, opts)
}

func (s *K8sService) CreateService(ctx context.Context, nameSpace, spaceUuid string, containerPort int32) (result *coreV1.Service, err error) {
	service := &coreV1.Service{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      constants.K8S_SERVICE_NAME_PREFIX + spaceUuid,
			Namespace: nameSpace,
		},
		Spec: coreV1.ServiceSpec{
			Ports: []coreV1.ServicePort{
				{
					Name: "http",
					Port: containerPort,
				},
			},
			Selector: map[string]string{
				"lad_app": spaceUuid,
			},
		},
	}
	return s.k8sClient.CoreV1().Services(nameSpace).Create(ctx, service, metaV1.CreateOptions{})
}

func (s *K8sService) CreateServiceByNodePort(ctx context.Context, nameSpace, taskUuid string, containerPort int32, nodePort int32, exclude22Port []int32) (result *coreV1.Service, err error) {
	var servicePort []coreV1.ServicePort

	servicePort = append(servicePort, coreV1.ServicePort{
		Name:     fmt.Sprintf("tcp-%d", containerPort),
		Port:     containerPort,
		NodePort: nodePort,
	})
	for _, port := range exclude22Port {
		servicePort = append(servicePort, coreV1.ServicePort{
			Name: fmt.Sprintf("tcp-%d", port),
			Port: port,
		})
	}

	service := &coreV1.Service{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      constants.K8S_SERVICE_NAME_PREFIX + taskUuid,
			Namespace: nameSpace,
		},
		Spec: coreV1.ServiceSpec{
			Type:  coreV1.ServiceTypeNodePort,
			Ports: servicePort,
			Selector: map[string]string{
				"hub-private": taskUuid,
			},
		},
	}
	return s.k8sClient.CoreV1().Services(nameSpace).Create(ctx, service, metaV1.CreateOptions{})
}

func (s *K8sService) DeleteService(ctx context.Context, namespace, serviceName string) error {
	return s.k8sClient.CoreV1().Services(namespace).Delete(ctx, serviceName, metaV1.DeleteOptions{})
}

func (s *K8sService) CreateIngress(ctx context.Context, k8sNameSpace, spaceUuid, hostName string, port int32) (*networkingv1.Ingress, error) {
	var ingressClassName = "nginx"
	ingress := &networkingv1.Ingress{
		ObjectMeta: metaV1.ObjectMeta{
			Name: constants.K8S_INGRESS_NAME_PREFIX + spaceUuid,
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/use-regex": "true",
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: &ingressClassName,
			Rules: []networkingv1.IngressRule{
				{
					Host: hostName,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/*",
									PathType: func() *networkingv1.PathType { t := networkingv1.PathTypePrefix; return &t }(),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: constants.K8S_SERVICE_NAME_PREFIX + spaceUuid,
											Port: networkingv1.ServiceBackendPort{
												Number: port,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return s.k8sClient.NetworkingV1().Ingresses(k8sNameSpace).Create(ctx, ingress, metaV1.CreateOptions{})
}

func (s *K8sService) DeleteIngress(ctx context.Context, nameSpace, ingressName string) error {
	return s.k8sClient.NetworkingV1().Ingresses(nameSpace).Delete(ctx, ingressName, metaV1.DeleteOptions{})
}

func (s *K8sService) CreateConfigMap(ctx context.Context, k8sNameSpace, jobUuid, basePath, configName string) (*coreV1.ConfigMap, error) {
	configFilePath := filepath.Join(basePath, configName)

	fileNameWithoutExt := filepath.Base(configName[:len(configName)-len(filepath.Ext(configName))])

	iniData, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	configMap := &coreV1.ConfigMap{
		ObjectMeta: metaV1.ObjectMeta{
			Name: jobUuid + "-" + fileNameWithoutExt,
		},
		Data: map[string]string{
			configName: string(iniData),
		},
	}
	return s.k8sClient.CoreV1().ConfigMaps(k8sNameSpace).Create(ctx, configMap, metaV1.CreateOptions{})
}

func (s *K8sService) GetPods(namespace, jobUuid string) (bool, error) {
	listOption := metaV1.ListOptions{}
	if jobUuid != "" {
		listOption = metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("lad_app=%s", jobUuid),
		}
	}
	podList, err := s.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), listOption)
	if err != nil {
		logs.GetLogger().Error(err)
		return false, err
	}
	if podList != nil && len(podList.Items) > 0 {
		return true, nil
	}
	return false, nil
}

func (s *K8sService) CreateNetworkPolicy(ctx context.Context, namespace string) (*networkingv1.NetworkPolicy, error) {
	networkPolicy := &networkingv1.NetworkPolicy{
		ObjectMeta: metaV1.ObjectMeta{
			Name:      namespace + "-" + generateString(10),
			Namespace: namespace,
		},
		Spec: networkingv1.NetworkPolicySpec{
			PolicyTypes: []networkingv1.PolicyType{networkingv1.PolicyTypeIngress, networkingv1.PolicyTypeEgress},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					From: []networkingv1.NetworkPolicyPeer{
						{
							NamespaceSelector: &metaV1.LabelSelector{
								MatchLabels: map[string]string{
									"kubernetes.io/metadata.name": namespace,
								},
							},
						},
					},
				},
			},
			Egress: []networkingv1.NetworkPolicyEgressRule{
				{
					To: []networkingv1.NetworkPolicyPeer{
						{
							NamespaceSelector: &metaV1.LabelSelector{
								MatchLabels: map[string]string{
									"kubernetes.io/metadata.name": namespace,
								},
							},
						},
					},
				},
			},
		},
	}

	return s.k8sClient.NetworkingV1().NetworkPolicies(namespace).Create(ctx, networkPolicy, metaV1.CreateOptions{})
}

func (s *K8sService) CreateNameSpace(ctx context.Context, nameSpace *coreV1.Namespace, opts metaV1.CreateOptions) (result *coreV1.Namespace, err error) {
	return s.k8sClient.CoreV1().Namespaces().Create(ctx, nameSpace, opts)
}

func (s *K8sService) GetNameSpace(ctx context.Context, nameSpace string, opts metaV1.GetOptions) (result *coreV1.Namespace, err error) {
	return s.k8sClient.CoreV1().Namespaces().Get(ctx, nameSpace, opts)
}

func (s *K8sService) DeleteNameSpace(ctx context.Context, nameSpace string) error {
	return s.k8sClient.CoreV1().Namespaces().Delete(ctx, nameSpace, metaV1.DeleteOptions{})
}

func (s *K8sService) ListUsedImage(ctx context.Context, nameSpace string) ([]string, error) {
	list, err := s.k8sClient.CoreV1().Pods(nameSpace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var usedImages []string
	for _, item := range list.Items {
		for _, status := range item.Status.ContainerStatuses {
			usedImages = append(usedImages, status.Image)
		}
	}
	return usedImages, nil
}

func (s *K8sService) ListNamespace(ctx context.Context) ([]string, error) {
	list, err := s.k8sClient.CoreV1().Namespaces().List(ctx, metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaces []string
	for _, item := range list.Items {
		namespaces = append(namespaces, item.Name)
	}
	return namespaces, nil
}

func (s *K8sService) CheckServiceNodePort(targetNodePort int32) (bool, int32, error) {
	services, err := s.k8sClient.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return false, 0, err
	}

	var usedPorts = make(map[int32]struct{})
	for _, service := range services.Items {
		for _, port := range service.Spec.Ports {
			usedPorts[port.NodePort] = struct{}{}
		}
	}

	var resultPort int32
	var flag bool
	if targetNodePort == 0 {
		for i := 30000; i < 32768; i++ {
			if _, ok := usedPorts[int32(i)]; !ok {
				resultPort = int32(i)
				flag = true
				break
			}
		}
	} else {
		if _, ok := usedPorts[targetNodePort]; ok {
			flag = false
			err = fmt.Errorf("invalid value: %d, provided port is already allocated", targetNodePort)
		} else {
			flag = true
			resultPort = targetNodePort
		}

	}
	return flag, resultPort, err
}

func (s *K8sService) StatisticalSources(ctx context.Context) ([]*models.NodeResource, error) {
	activePods, err := s.GetAllActivePod(ctx)
	if err != nil {
		return nil, err
	}
	var nodeList []*models.NodeResource

	nodes, err := s.k8sClient.CoreV1().Nodes().List(ctx, metaV1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	nodeGpuInfoMap, err := s.GetResourceExporterPodLog(ctx)
	if err != nil {
		logs.GetLogger().Errorf("failed to collect cluster gpu info, if have available gpu, please check resource-exporter. error: %+v", err)
	}

	for _, node := range nodes.Items {
		nodeGpu, _, nodeResource := GetNodeResource(activePods, &node)
		if nodeGpuInfoMap != nil {
			collectGpu := make(map[string]collectGpuInfo)
			if gpu, ok := nodeGpuInfoMap[node.Name]; ok {
				for index, gpuDetail := range gpu.Gpu.Details {
					gpuName := strings.ReplaceAll(gpuDetail.ProductName, " ", "-")
					if v, ok := collectGpu[gpuName]; ok {
						v.count += 1
						collectGpu[gpuName] = v
					} else {
						collectGpu[gpuName] = collectGpuInfo{
							index,
							1,
							0,
						}
					}
				}

				for name, info := range collectGpu {
					runCount := int(nodeGpu[name])
					if runCount < info.count {
						info.remainNum = info.count - runCount
					} else {
						info.remainNum = 0
					}
					collectGpu[name] = info
				}

				var counter = make(map[string]int)
				newGpu := make([]models.GpuDetail, 0)
				for _, gpuDetail := range gpu.Gpu.Details {
					gpuName := strings.ReplaceAll(gpuDetail.ProductName, " ", "-")
					newDetail := gpuDetail
					g := collectGpu[gpuName]
					if g.remainNum > 0 && counter[gpuName] < g.remainNum {
						newDetail.Status = models.Available
						counter[gpuName] += 1
					} else {
						newDetail.Status = models.Occupied
					}
					newGpu = append(newGpu, newDetail)
				}
				nodeResource.Gpu = models.Gpu{
					DriverVersion: gpu.Gpu.DriverVersion,
					CudaVersion:   gpu.Gpu.CudaVersion,
					AttachedGpus:  gpu.Gpu.AttachedGpus,
					Details:       newGpu,
				}
			}
		}

		nodeList = append(nodeList, nodeResource)
	}

	gpuResourceCache.Range(func(applyGpu, num any) bool {
		needNum := num.(int)
	loop:
		for i, node := range nodeList {
			for j, gpu := range node.Gpu.Details {
				if needNum == 0 {
					break loop
				}
				if applyGpu == gpu.ProductName && needNum > 0 {
					if gpu.Status == models.Available {
						nodeList[i].Gpu.Details[j].Status = models.Occupied
						needNum--
					}
				}
			}
		}
		return true
	})

	return nodeList, nil
}

func (s *K8sService) GetResourceExporterPodLog(ctx context.Context) (map[string]models.CollectNodeInfo, error) {
	var num int64 = 1
	podLogOptions := coreV1.PodLogOptions{
		Container:  "",
		TailLines:  &num,
		Timestamps: false,
	}

	podList, err := s.k8sClient.CoreV1().Pods("kube-system").List(ctx, metaV1.ListOptions{
		LabelSelector: "app=resource-exporter",
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	result := make(map[string]models.CollectNodeInfo)
	for _, pod := range podList.Items {
		podLog, err := s.GetPodLogByPodName("kube-system", pod.Name, &podLogOptions)
		if err != nil {
			logs.GetLogger().Errorf("collect gpu deatil info, nodeName: %s, error: %+v", pod.Spec.NodeName, err)
			continue
		}

		var nodeInfo models.CollectNodeInfo
		if err := json.Unmarshal([]byte(podLog), &nodeInfo); err != nil {
			s.k8sClient.CoreV1().Pods("kube-system").Delete(ctx, pod.Name, metaV1.DeleteOptions{})
			continue
		}
		result[pod.Spec.NodeName] = nodeInfo
	}
	return result, nil
}

func (s *K8sService) GetPodLogByPodName(namespace, podName string, podLogOptions *coreV1.PodLogOptions) (string, error) {
	req := s.k8sClient.CoreV1().Pods(namespace).GetLogs(podName, podLogOptions)
	buf, err := readLog(req)
	if err != nil {
		logs.GetLogger().Errorf("get pod log failed, podName: %s, error: %+v", podName, err)
		return "", err
	}
	return buf.String(), nil
}

func (s *K8sService) AddNodeLabel(nodeName, key string) error {
	key = strings.ReplaceAll(key, " ", "-")

	node, err := s.k8sClient.CoreV1().Nodes().Get(context.Background(), nodeName, metaV1.GetOptions{})
	if err != nil {
		return err
	}
	node.Labels[key] = "true"
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, updateErr := s.k8sClient.CoreV1().Nodes().Update(context.Background(), node, metaV1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		return fmt.Errorf("failed update node label: %w", retryErr)
	}
	return nil
}

func (s *K8sService) WaitForPodRunningByHttp(namespace, jobUuid, serviceIp string) (string, error) {
	var podName string
	var podErr = errors.New("get pod status failed")

	retryErr := retry.OnError(wait.Backoff{
		Steps:    120,
		Duration: 10 * time.Second,
	}, func(err error) bool {
		return err != nil && err.Error() == podErr.Error()
	}, func() error {
		if _, err := http.Get(serviceIp); err != nil {
			return podErr
		}
		podList, err := s.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("lad_app==%s", jobUuid),
		})
		if err != nil {
			logs.GetLogger().Error(err)
			return podErr
		}
		podName = podList.Items[0].Name

		return nil
	})

	if retryErr != nil {
		return podName, fmt.Errorf("failed waiting for pods to be running: %v", retryErr)
	}
	return podName, nil
}

func (s *K8sService) WaitForPodRunningByTcp(namespace, jobUuid string) (string, error) {
	var podName string
	err := wait.PollImmediate(time.Second*5, time.Minute*10, func() (done bool, err error) {
		podList, err := s.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("hub-private==%s", jobUuid),
		})
		if err != nil {
			logs.GetLogger().Error(err)
			return false, err
		}
		for _, pod := range podList.Items {
			if pod.Status.Phase != coreV1.PodRunning {
				return false, nil
			}
		}
		if len(podList.Items) == 0 {
			return false, nil
		}
		podName = podList.Items[0].Name
		return true, nil
	})

	if err != nil {
		return "", fmt.Errorf("get pod status failed, error: %v", err)
	}

	return podName, nil
}

func (s *K8sService) PodDoCommand(namespace, podName, containerName string, podCmd []string) error {
	reader, writer := io.Pipe()
	req := s.k8sClient.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&coreV1.PodExecOptions{
			Container: containerName,
			Command:   podCmd,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(s.config, "POST", req.URL())
	if err != nil {
		return fmt.Errorf("failed to create spdy client: %w", err)
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:  reader,
		Stdout: writer,
		Stderr: writer,
		Tty:    true,
	})
	if err != nil {
		return fmt.Errorf("failed to create stream: %w", err)
	}

	return nil
}

func (s *K8sService) GetNodeGpuSummary(ctx context.Context) (map[string]map[string]int64, error) {
	nodeGpuInfoMap, err := s.GetResourceExporterPodLog(ctx)
	if err != nil {
		logs.GetLogger().Errorf("Collect cluster gpu info Failed, if have available gpu, please check resource-exporter. error: %+v", err)
		return map[string]map[string]int64{}, err
	}

	var nodeGpuSummary = make(map[string]map[string]int64)
	for nodeName, gpu := range nodeGpuInfoMap {
		collectGpu := make(map[string]int64)
		for _, gpuDetail := range gpu.Gpu.Details {
			gpuName := strings.ReplaceAll(gpuDetail.ProductName, " ", "-")
			if v, ok := collectGpu[gpuName]; ok {
				v += 1
				collectGpu[gpuName] = v
			} else {
				collectGpu[gpuName] = 1
			}
		}
		nodeGpuSummary[nodeName] = collectGpu
	}
	return nodeGpuSummary, nil
}

func (s *K8sService) GetAllActivePod(ctx context.Context) ([]coreV1.Pod, error) {
	allPods, err := clientSet.CoreV1().Pods("").List(ctx, metaV1.ListOptions{
		FieldSelector: "status.phase=Running",
	})
	if err != nil {
		return nil, err
	}
	return allPods.Items, nil
}

func (s *K8sService) GetAPIServerEndpoint() string {
	last := strings.LastIndex(s.config.Host, ":")
	return s.config.Host[:last]
}

func (s *K8sService) GetDeploymentActiveCount() (int, error) {
	namespaces, err := s.ListNamespace(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
		return 0, err
	}

	var total int
	for _, namespace := range namespaces {
		if strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
			deployments, err := s.k8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
			if err != nil {
				logs.GetLogger().Errorf("Error getting deployments in namespace %s: %v\n", namespace, err)
				continue
			}

			for _, deployment := range deployments.Items {
				creationTimestamp := deployment.ObjectMeta.CreationTimestamp.Time
				currentTime := time.Now()
				age := currentTime.Sub(creationTimestamp)
				if deployment.Status.AvailableReplicas > 0 && age.Hours() > 0 {
					total++
				}
			}
		}
	}
	return total, nil
}

func (s *K8sService) GetGlobalNetworkSet(gnsName string) (*calicov3.GlobalNetworkSet, error) {
	calicoCs, err := calicoclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create calico client, error: %v", err)
	}
	networkSetList, err := calicoCs.ProjectcalicoV3().GlobalNetworkSets().List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get globalNetworkSets list, error: %v", err)
	}

	var gns calicov3.GlobalNetworkSet
	for _, item := range networkSetList.Items {
		if item.Name == gnsName {
			gns = item
		}
	}
	return &gns, nil
}

func (s *K8sService) GetGlobalNetworkPolicy(gnpName string) (*calicov3.GlobalNetworkPolicy, error) {
	calicoCs, err := calicoclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create calico client, error: %v", err)
	}

	networkPolicyList, err := calicoCs.ProjectcalicoV3().GlobalNetworkPolicies().List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get globalNetworkSets list, error: %v", err)
	}

	var gnp calicov3.GlobalNetworkPolicy
	for _, item := range networkPolicyList.Items {
		if item.Name == gnpName {
			gnp = item
		} else {
			if models.ExistResource(item.Name) {
				continue
			}
		}
	}
	return &gnp, nil
}

func (s *K8sService) GetUsedNodePorts() (map[int32]struct{}, error) {
	services, err := s.k8sClient.CoreV1().Services(metaV1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	nodePorts := make(map[int32]struct{})
	for _, svc := range services.Items {
		for _, port := range svc.Spec.Ports {
			if port.NodePort != 0 {
				nodePorts[port.NodePort] = struct{}{}
			}
		}
	}
	return nodePorts, nil
}

func generateVolume() ([]coreV1.VolumeMount, []coreV1.Volume) {
	fileType := coreV1.HostPathFile
	return []coreV1.VolumeMount{
			{
				Name:      "shm",
				MountPath: "/dev/shm",
			},
			{
				Name:      "lxcfs-proc-cpuinfo",
				MountPath: "/proc/cpuinfo",
				ReadOnly:  true,
			},
			{
				Name:      "system-cpu-online",
				MountPath: "/sys/devices/system/cpu/online",
			},
			{
				Name:      "lxcfs-proc-diskstats",
				MountPath: "/proc/diskstats",
				ReadOnly:  true,
			},
			{
				Name:      "lxcfs-proc-meminfo",
				MountPath: "/proc/meminfo",
				ReadOnly:  true,
			},
			{
				Name:      "lxcfs-proc-stat",
				MountPath: "/proc/stat",
				ReadOnly:  true,
			},
			{
				Name:      "lxcfs-proc-swaps",
				MountPath: "/proc/swaps",
				ReadOnly:  true,
			},
			{
				Name:      "lxcfs-proc-uptime",
				MountPath: "/proc/uptime",
				ReadOnly:  true,
			},
		}, []coreV1.Volume{
			{
				Name: "shm",
				VolumeSource: coreV1.VolumeSource{
					EmptyDir: &coreV1.EmptyDirVolumeSource{
						Medium: coreV1.StorageMediumMemory,
					},
				},
			},
			{
				Name: "lxcfs-proc-cpuinfo",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/cpuinfo",
						Type: &fileType,
					},
				},
			},
			{
				Name: "system-cpu-online",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/sys/devices/system/cpu/online",
						Type: &fileType,
					},
				},
			},
			{
				Name: "lxcfs-proc-diskstats",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/diskstats",
						Type: &fileType,
					},
				},
			},
			{
				Name: "lxcfs-proc-meminfo",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/meminfo",
						Type: &fileType,
					},
				},
			},
			{
				Name: "lxcfs-proc-stat",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/stat",
						Type: &fileType,
					},
				},
			},
			{
				Name: "lxcfs-proc-swaps",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/swaps",
						Type: &fileType,
					},
				},
			},
			{
				Name: "lxcfs-proc-uptime",
				VolumeSource: coreV1.VolumeSource{
					HostPath: &coreV1.HostPathVolumeSource{
						Path: "/var/lib/lxcfs/proc/uptime",
						Type: &fileType,
					},
				},
			},
		}
}

func (s *K8sService) GetClusterRuntime() (string, error) {
	nodes, err := s.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get node list, error: %v", err)
	}

	var runtime string
	for _, node := range nodes.Items {
		runtime = node.Status.NodeInfo.ContainerRuntimeVersion
		break
	}
	return runtime, nil
}

func readLog(req *rest.Request) (*strings.Builder, error) {
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return nil, err
	}
	defer podLogs.Close()
	buf := new(strings.Builder)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func generateLabel(name string) map[string]string {
	if name != "" {
		return map[string]string{
			name: "true",
		}
	} else {
		return map[string]string{}
	}
}

type collectGpuInfo struct {
	index     int
	count     int
	remainNum int
}
