# Computing Provider
[![Discord](https://img.shields.io/discord/770382203782692945?label=Discord&logo=Discord)](https://discord.gg/Jd2BFSVCKw)
[![Twitter Follow](https://img.shields.io/twitter/follow/swan_chain)](https://twitter.com/swan_chain)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)

A computing provider is an individual or organization that participates in the decentralized computing network by offering computational resources such as processing power (CPU and GPU), memory, storage, and bandwidth.

As a resource provider, you can run a **ECP**(Edge Computing Provider) and **FCP**(Fog Computing Provider) to contribute yourcomputing resource.


 - **ECP (Edge Computing Provider)** specializes in processing data at the source of data generation, using minimal latency setups ideal for real-time applications. This provider handles specific, localized tasks directly on devices at the network’s edge, such as IoT devices. At the current stage, ECP supports the generation of **ZK-Snark proof of Filecoin network**, and more ZK proof types will be gradually supported, such as Aleo, Scroll, starkNet, etc. [Install Guideline](ubi/README.md)


 - **FCP (Fog Computing Provider)** Offers a layered network that extends cloud capabilities to the edge of the network, providing services such as AI model training and deployment. This provider utilizes infrastructure like Kubernetes (K8S) to support scalable, distributed computing tasks.  **FCP** will execute tasks assigned by Market Provider, like [Orchestrator](https://orchestrator.swanchain.io) on the [Swan chain](https://swanchain.io).


# Table of Content

-  As a ECP
	- [Run Edge Computing Provider](ubi/README.md)

- As a FCP
 	- [Prerequisites](#Prerequisites)
 	- [Install the Kubernetes](#Install-the-Kubernetes)
 		- [Install Container Runtime Environment](#install-Container-Runtime-Environment)
 		- [Optional - Setup a docker registry server](#Optional-setup-a-Docker-Registry-Server)
		- [Create a Kubernetes Cluster](#Create-a-Kubernetes-Cluster)
 		- [Install the Network Plugin](#Install-the-Network-Plugin)
		- [Install the NVIDIA Plugin](#Install-the-NVIDIA-Plugin)
		- [Install the Ingress-nginx Controller](#Install-the-Ingress-nginx-Controller)
 	- [Install and config the Nginx](#Install-and-config-the-nginx)
 	- [Install the Hardware resource-exporter](#Install-the-Hardware-resource-exporter)
 	- [Build and config the Computing Provider](#Build-and-config-the-Computing-Provider)
 	- [Install AI Inference Dependency(Optional)](#optional-Install-AI-Inference-Dependency)
 	- [Config and Receive UBI Tasks(Optional)](#optional-Config-and-Receive-UBI-Tasks)
	 - [Start the Computing Provider](#Start-the-Computing-Provider)
	 - [CLI of Computing Provider](#CLI-of-Computing-Provider)
 
## Prerequisites
Before you install the Computing Provider, you need to know there are some resources required:
 - Possess a public IP
 - Have a domain name (*.example.com)
 - Have an SSL certificate
 - `Go` version must 1.21+, you can refer here:

```bash
wget -c https://golang.org/dl/go1.21.7.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc
```

## Install the Kubernetes
The Kubernetes version should be `v1.24.0+`

###  Install Container Runtime Environment
If you plan to run a Kubernetes cluster, you need to install a container runtime into each node in the cluster so that Pods can run there, refer to [here](https://kubernetes.io/docs/setup/production-environment/container-runtimes/). And you just need to choose one option to install the `Container Runtime Environment`

#### Option 1: Install the `Docker` and `cri-dockerd` （**Recommended**）
To install the `Docker Container Runtime` and the `cri-dockerd`, follow the steps below:
* Install the `Docker`:
    - Please refer to the official documentation from [here](https://docs.docker.com/engine/install/).
* Install `cri-dockerd`:
    - `cri-dockerd` is a CRI (Container Runtime Interface) implementation for Docker. You can install it refer to [here](https://github.com/Mirantis/cri-dockerd).

#### Option 2: Install the `Docker` and the`Containerd`
* Install the `Docker`:
    - Please refer to the official documentation from [here](https://docs.docker.com/engine/install/).
* To install `Containerd` on your system:
  - `Containerd` is an industry-standard container runtime that can be used as an alternative to Docker. To install `containerd` on your system, follow the instructions on [getting started with containerd](https://github.com/containerd/containerd/blob/main/docs/getting-started.md).

### Optional-Setup a docker registry server
**If you are using the docker and you have only one node, the step can be skipped**.

If you have deployed a Kubernetes cluster with multiple nodes, it is recommended to set up a **private Docker Registry** to allow other nodes to quickly pull images within the intranet. 

* Create a directory `/docker_repo` on your docker server. It will be mounted on the registry container as persistent storage for our docker registry.
```bash
sudo mkdir /docker_repo
sudo chmod -R 777 /docker_repo
```
* Launch the docker registry container:
```bash
sudo docker run --detach \
  --restart=always \
  --name registry \
  --volume /docker_repo:/docker_repo \
  --env REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY=/docker_repo \
  --publish 5000:5000 \
  registry:2
```
![1](https://github.com/lagrangedao/go-computing-provider/assets/102578774/0c4cd53d-fb5f-43d9-b804-be83faf33986)


* Add the registry server to the node

 	- If you have installed the `Docker` and `cri-dockerd`(**Option 1**), you can update every node's configuration:


	```bash
	sudo vi /etc/docker/daemon.json
	```
	```
	## Add the following config
	"insecure-registries": ["<Your_registry_server_IP>:5000"]
	```
	Then restart the docker service
	```bash
	sudo systemctl restart docker
	```

 	- If you have installed the `containerd`(**Option 2**), you can update every node's configuration:

```bash
[plugins."io.containerd.grpc.v1.cri".registry]
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
    [plugins."io.containerd.grpc.v1.cri".registry.mirrors."<Your_registry_server_IP>:5000"]
      endpoint = ["http://<Your_registry_server_IP>:5000"]

[plugins."io.containerd.grpc.v1.cri".registry.configs]
  [plugins."io.containerd.grpc.v1.cri".registry.configs."<Your_registry_server_IP>:5000".tls]
      insecure_skip_verify = true                                                               
```

Then restart `containerd` service

```bash
sudo systemctl restart containerd
```
**<Your_registry_server_IP>:** the intranet IP address of your registry server.

Finally, you can check the installation by the command:
```bash
docker system info
```
![2](https://github.com/lagrangedao/go-computing-provider/assets/102578774/4cfc1981-3fca-415c-948f-86c496915cff)




### Create a Kubernetes Cluster
To create a Kubernetes cluster, you can use a container management tool like `kubeadm`. The below steps can be followed:

* [Install the kubeadm toolbox](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/).

* [Create a Kubernetes cluster with kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/) 


### Install the Network Plugin
Calico is an open-source **networking and network security solution for containers**, virtual machines, and native host-based workloads. Calico supports a broad range of platforms including **Kubernetes**, OpenShift, Mirantis Kubernetes Engine (MKE), OpenStack, and bare metal services.

To install Calico, you can follow the below steps, more information can be found [here](https://docs.tigera.io/calico/3.25/getting-started/kubernetes/quickstart).

**step 1**: Install the Tigera Calico operator and custom resource definitions
```
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.25.1/manifests/tigera-operator.yaml
```

**step 2**: Install Calico by creating the necessary custom resource
```
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.25.1/manifests/custom-resources.yaml
```
**step 3**: Confirm that all of the pods are running with the following command
```
watch kubectl get pods -n calico-system
```
**step 4**: Remove the taints on the control plane so that you can schedule pods on it.
```
kubectl taint nodes --all node-role.kubernetes.io/control-plane-
kubectl taint nodes --all node-role.kubernetes.io/master-
```
If you have installed it correctly, you can see the result shown in the figure by the command `kubectl get po -A`

![3](https://github.com/lagrangedao/go-computing-provider/assets/102578774/91ef353f-72af-41b2-82e8-061b92bfb999)

**Note:** 
 - If you are a single-host Kubernetes cluster, remember to remove the taint mark, otherwise, the task can not be scheduled to it.
```bash
kubectl taint node ${nodeName}  node-role.kubernetes.io/control-plane:NoSchedule-
```

### Install the NVIDIA Plugin
If your computing provider wants to provide a GPU resource, the NVIDIA Plugin should be installed, please follow the steps:

* [Install NVIDIA Driver](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/install-guide.html#nvidia-drivers).
>Recommend NVIDIA Linux drivers version should be 470.xx+

* [Install NVIDIA Device Plugin for Kubernetes](https://github.com/NVIDIA/k8s-device-plugin#quick-start).

If you have installed it correctly, you can see the result shown in the figure by the command 
`kubectl get po -n kube-system`

![4](https://github.com/lagrangedao/go-computing-provider/assets/102578774/8209c589-d561-43ad-adea-5ecb52618909)

### Install the Ingress-nginx Controller
The `ingress-nginx` is an ingress controller for Kubernetes using `NGINX` as a reverse proxy and load balancer. You can run the following command to install it:
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.7.1/deploy/static/provider/cloud/deploy.yaml
```
If you have installed it correctly, you can see the result shown in the figure by the command: 

 - Run `kubectl get po -n ingress-nginx`

![5](https://github.com/lagrangedao/go-computing-provider/assets/102578774/f3c0585a-df19-4971-91fe-d03365f4edee)

 - Run `kubectl get svc -n ingress-nginx`

![6](https://github.com/lagrangedao/go-computing-provider/assets/102578774/e3b3dadc-77c1-4dc0-843c-5b946e252b65)

### Install and config the Nginx
 -  Install `Nginx` service to the Server
```bash
sudo apt update
sudo apt install nginx
```
 -  Add a configuration for your Domain name
Assume your domain name is `*.example.com`
```
vi /etc/nginx/conf.d/example.conf
```

```bash
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}

server {
        listen 80;
        listen [::]:80;
        server_name *.example.com;                                           # need to your domain
        return 301 https://$host$request_uri;
        #client_max_body_size 1G;
}
server {
        listen 443 ssl;
        listen [::]:443 ssl;
        ssl_certificate  /etc/letsencrypt/live/example.com/fullchain.pem;     # need to config SSL certificate
        ssl_certificate_key  /etc/letsencrypt/live/example.com/privkey.pem;   # need to config SSL certificate

        server_name *.example.com;                                            # need to config your domain
        location / {
          proxy_pass http://127.0.0.1:<port>;  	# Need to configure the Intranet port corresponding to ingress-nginx-controller service port 80 
          proxy_set_header Host $http_host;
          proxy_set_header Upgrade $http_upgrade;
          proxy_set_header Connection $connection_upgrade;
       }
}
```

 - **Note:** 

	 - `server_name`: a generic domain name

	 - `ssl_certificate` and `ssl_certificate_key`: certificate for https.

	 - `proxy_pass`:  The port should be the Intranet port corresponding to `ingress-nginx-controller` service port 80

 - Reload the `Nginx` config
	```bash
	sudo nginx -s reload
	```
 - Map your "catch-all (wildcard) subdomain(*.example.com)" to a public IP address



### Install the Hardware resource-exporter
 The `resource-exporter` plugin is developed to collect the node resource constantly, computing provider will report the resource to the Lagrange Auction Engine to match the space requirement. To get the computing task, every node in the cluster must install the plugin. You just need to run the following command:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: DaemonSet
metadata:
  namespace: kube-system
  name: resource-exporter-ds
  labels:
    app: resource-exporter
spec:
  selector:
    matchLabels:
      app: resource-exporter
  template:
    metadata:
      labels:
        app: resource-exporter
    spec:
      containers:
      - name: resource-exporter
        image: filswan/resource-exporter:v11.2.8
        imagePullPolicy: IfNotPresent
EOF
```
If you have installed it correctly, you can see the result shown in the figure by the command:
`kubectl get po -n kube-system`

![7](https://github.com/lagrangedao/go-computing-provider/assets/102578774/38b0e15f-5ff9-4edc-a313-d0f6f4a0bda8)

## Build and config the Computing Provider

 - Build the Computing Provider 

	Firstly, clone the code to your local:
```bash
git clone https://github.com/swanchain/go-computing-provider.git
cd go-computing-provider
git checkout mainnet
```

Then build the Computing provider on the **Swan Mainnet** by following the below steps:

```bash
make clean && make mainnet
make install
```

> If you want to test the CP in the **testnet**, please build a testnet version:
> ```bash
> make clean && make testnet
> make install
> ```

## Initialize CP repo and Update Configuration 
1. Initialize repo
    ```
    computing-provider init --multi-address=/ip4/<YOUR_PUBLIC_IP>/tcp/<YOUR_PORT> --node-name=<YOUR_NODE_NAME>
    ```
    **Note:**
    - By default, the CP's repo is `~/.swan/computing`, you can configure it by `export CP_PATH="<YOUR_CP_PATH>"`
    - The CP service port (`8085` by default) must be mapped to the public IP address and port
      2. Update `config.toml`

          Edit the necessary configuration files according to your deployment requirements. 

          ```
             [API]
             Port = 8085                                    # The port number that the web server listens on
             MultiAddress = "/ip4/<public_ip>/tcp/<port>"   # The multiAddress for libp2p
             Domain = ""                                    # The domain name
             NodeName = ""                                  # The computing-provider node name
             WalletWhiteList = ""                           # CP only accepts user addresses from this whitelist for space deployment
             WalletBlackList = ""                           # CP reject user addresses from this blacklist for space deployment
 
             [UBI]
             UbiEnginePk = "0xB5aeb540B4895cd024c1625E146684940A849ED9"              # UBI Engine's public key, CP only accept the task from this UBI engine
             EnableSequencer = true                                                  # Batch submission of proof messages through Sequencer service
             AutoChainProof = false                                                  # Sequencer insufficient balance or service unavailable ，use chain to submit proof
       
             [LOG]
             CrtFile = "/YOUR_DOMAIN_NAME_CRT_PATH/server.crt"                       # Your domain name SSL .crt file path
             KeyFile = "/YOUR_DOMAIN_NAME_KEY_PATH/server.key"                       # Your domain name SSL .key file path
	
             [HUB]
             ServerUrl = "https://orchestrator-mainnet-api.swanchain.io"             # The Orchestrator's API address
             AccessToken = ""                                               	       # The Orchestrator's access token, Acquired from "https://orchestrator.swanchain.io", and switch to the `mainnet network` and use the owner address
             BalanceThreshold= 10                                                    # The cp’s collateral balance threshold
             OrchestratorPk = "0x4B98086A20f3C19530AF32D21F85Bc6399358e20"           # Orchestrator's public key, CP only accept the task from this Orchestrator
             VerifySign = true                                                       # Verify that the task signature is from Orchestrator
	
             [MCS]
             ApiKey = ""                                   # Acquired from "https://www.multichain.storage" -> setting -> Create API Key
             BucketName = ""                               # Acquired from "https://www.multichain.storage" -> bucket -> Add Bucket
             Network = "polygon.mainnet"                   # polygon.mainnet for mainnet, polygon.mumbai for testnet
	
             [Registry]
             ServerAddress = ""                            # The docker container image registry address, if only a single node, you can ignore
             UserName = ""                                 # The login username, if only a single node, you can ignore
             Password = ""                                 # The login password, if only a single node, you can ignore
	
             [RPC]
             SWAN_CHAIN_RPC = "https://mainnet-rpc01.swanchain.io"     # Swan chain RPC
          ```

**Note:**  
* Example `[api].WalletWhiteList` hosted on GitHub can be found [here](https://raw.githubusercontent.com/swanchain/market-providers/main/clients/whitelist.txt).
* Example `[api].WalletBlackList` hosted on GitHub can be found [here](https://raw.githubusercontent.com/swanchain/market-providers/main/clients/blacklist.txt).

## Initialize a Wallet and Deposit `SwanETH`
1.  Generate a new wallet address or import the previous wallet:

	```bash
	computing-provider wallet new
	```
	
	Example output:
	```
	0x7791f48931DB81668854921fA70bFf0eB85B8211
	```
	
	**or** import your wallet:
	```bash
	# Import wallet using the private key
	computing-provider wallet import <YOUR_PRIVATE_KEY_FILE>
	```
	**Note:** `<YOUR_PRIVATE_KEY_FILE>` is a file that contains the private key

2.  Deposit `SwanETH` to the wallet address:
	```bash
	computing-provider wallet send --from <YOUR_WALLET_ADDRESS> 0x7791f48931DB81668854921fA70bFf0eB85B8211 0.01
	```
	**Note:** If you don't have `SwanETH` and `SWANC`, please follow [the guideline](https://docs.swanchain.io/swan-mainnet/getting-started-guide) to [bridge ETH to Swan Mainnet](https://bridge.swanchain.io), and [claim the SWANC](https://faucet.swanchain.io) as collaterals.

## Initialization CP Account
Deploy a CP account contract:
```bash
computing-provider account create --ownerAddress <YOUR_OWNER_WALLET_ADDRESS> \
	--workerAddress <YOUR_WORKER_WALLET_ADDRESS> \
	--beneficiaryAddress <YOUR_BENEFICIARY_WALLET_ADDRESS>  \
	--task-types 3
```
**Note:** `--task-types`: Supports 4 task types (`1`: Fil-C2-512M, `2`: Aleo, `3`: AI, `4`: Fil-C2-32G), separated by commas. For FCP, it needs to be set to 3.

**Output:**
```
Contract deployed! Address: 0x3091c9647Ea5248079273B52C3707c958a3f2658
Transaction hash: 0xb8fd9cc9bfac2b2890230b4f14999b9d449e050339b252273379ab11fac15926
```

## Collateral `SWANC` for FCP
```bash
 computing-provider collateral add --fcp --from <YOUR_WALLET_ADDRESS>  <amount>
```
**Note:** Currently one AI task requires 5 `SWANC`. Please deposit enough collaterals for the tasks


## Withdraw `SWANC` from FCP
```bash
 computing-provider collateral withdraw --fcp --owner <YOUR_WALLET_ADDRESS> --account <YOUR_CP_ACCOUNT> <amount>
```
**Note:** If you want to withdraw the funds from FCP, you can run the above command


## Start the Computing Provider
You can run `computing-provider` using the following command
```bash
export CP_PATH=<YOUR_CP_PATH>
nohup computing-provider run >> cp.log 2>&1 & 
```
---
## [**OPTIONAL**] Install AI Inference Dependency
It is necessary for the Computing Provider to deploy the AI inference endpoint. But if you do not want to support the feature, you can skip it.
```bash
export CP_PATH=<YOUR_CP_PATH>
./install.sh
```

## [**OPTIONAL**] Config and Receive UBI Tasks
### **Step 1: Prerequisites:** Perform Filecoin Commit2 (fil-c2) ZK tasks.
1. Download parameters (specify the path with PARENT_PATH variable):
	```bash
	# At least 200G storage is needed
	export PARENT_PATH="<V28_PARAMS_PATH>"
	
	# 512MiB parameters
	curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-512.sh | bash
	
	# 32GiB parameters
	curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-32.sh | bash
	```
2. Configure environment variables in `fil-c2.env` under CP repo (`$CP_PATH`):

    ```bash
    FIL_PROOFS_PARAMETER_CACHE=$PARENT_PATH
    RUST_GPU_TOOLS_CUSTOM_GPU="GeForce RTX 3080:8704" 
    ```

* Adjust the value of `RUST_GPU_TOOLS_CUSTOM_GPU` based on the GPU used by the CP's Kubernetes cluster for fil-c2 tasks.
* For more device choices, please refer to this page:[https://github.com/filecoin-project/bellperson](https://github.com/filecoin-project/bellperson)

### Step 2: Collateral `SWANC` for receiving ZK Task

```bash
computing-provider collateral add --ecp --from <YOUR_WALLET_ADDRESS>  <amount>
```
**Note:** Currently one zk-task requires 0.0005 SWANC.

Example output:

```
0x7791f48931DB81668854921fA70bFf0eB85B8211
```
### Step 3: Add the type of ZK task

```bash
computing-provider account changeTaskTypes --ownerAddress <YOUR_OWNER_WALLET_ADDRESS> 1,2,3,4
```
**Note:** `--task-types` Supports 4 task types:
 - `1`: FIL-C2-512M
 - `2`: Aleo
 - `3`: AI
 - `4`: FIL-C2-32G

If you need to run FCP and ECP at the same time, you need to set it to `1,2,3,4`


### **Step 4: Account Management**

Use `computing-provider account` subcommands to update CP details:

```
computing-provider account -h
NAME:
   computing-provider account - Manage account info of CP

USAGE:
   computing-provider account command [command options] [arguments...]

COMMANDS:
   create                    Create a cp account to chain
   changeMultiAddress        Update MultiAddress of CP (/ip4/<public_ip>/tcp/<port>)
   changeOwnerAddress        Update OwnerAddress of CP
   changeWorkerAddress       Update workerAddress of CP
   changeBeneficiaryAddress  Update beneficiaryAddress of CP
   changeTaskTypes           Update taskTypes of CP (1:Fil-C2-512M, 2:Aleo, 3: AI, 4:Fil-C2-32G), separated by commas
   help, h                   Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

### Step 6: Check the Status of ZK task;

To check the ZK task list, use the following command:

```
computing-provider ubi list --show-failed
```

Example output:

```
TASK ID TASK TYPE       ZK TYPE         TRANSACTION HASH                                                        STATUS  REWARD  CREATE TIME         
2       CPU             fil-c2-512M     0xb06b3a8c2b2b96b564777a3866e27ce7c61631f77e5de3196e93eb916b0d2575      success 2.0     2024-01-20 03:30:30
33      CPU             fil-c2-512M     0x7567435e83a4a019a6356da8cf33e64a071f2d3355fce5289b9c17cf0144f282      success 2.0     2024-01-18 15:58:21
13      CPU             fil-c2-512M     0x7b3081314891aad3788c84935c67f9be0a8acc6b4fc77c5aa6fdfda728877fde      success 2.0     2024-01-20 04:27:40
238     CPU             fil-c2-512M     0xb8eb1f7b3cfc8210fa5546adc528f230241110e5cc9b4900725a9da28895aad9      success 2.0     2024-01-18 17:08:21
```

## Restart the Computing Provider
You can run `computing-provider` using the following command
```bash
export CP_PATH=<YOUR_CP_PATH>
nohup computing-provider run >> cp.log 2>&1 & 
```

## CLI of Computing Provider
* Check the current list of tasks running on CP, display detailed information for tasks using `-v`
```
computing-provider task list 
```
* Retrieve detailed information for a specific task using `task_uuid`
```
computing-provider task get [task_uuid]
```
* Delete task by `task_uuid`
```
computing-provider task delete [task_uuid]
```

## Getting Help

For usage questions or issues reach out to the Swan team either in the [Discord channel](https://discord.gg/3uQUWzaS7U) or open a new issue here on GitHub.

## License

Apache
