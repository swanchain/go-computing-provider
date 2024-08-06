package constants

const K8S_NAMESPACE_NAME_PREFIX = "ns-"
const K8S_CONTAINER_NAME_PREFIX = "pod-"
const K8S_PRIVATE_CONTAINER_PREFIX = "vm-"
const K8S_INGRESS_NAME_PREFIX = "ing-"
const K8S_SERVICE_NAME_PREFIX = "svc-"
const K8S_DEPLOY_NAME_PREFIX = "deploy-"

const CPU_AMD = "AMD"
const CPU_INTEL = "INTEL"

const SPACE_TYPE_PRIVATE = "private"
const SPACE_TYPE_PUBLIC = "public"

const (
	BidMode_All = iota
	BidMode_Auto
	BidMode_Private
	BidMode_None
)
