# Fog Computing Provider (FCP)

Fog Computing Provider (FCP) offers a layered network that extends cloud capabilities to the edge of the network, providing services such as AI model training and deployment. This provider utilizes infrastructure like Kubernetes (K8S) to support scalable, distributed computing tasks.

## Overview

FCP executes tasks assigned by Market Providers, like the [Orchestrator](https://orchestrator.swanchain.io) on the [Swan Chain](https://swanchain.io). It's designed for high-performance, scalable computing workloads.

### Supported Tasks

- **AI Model Training**: Machine learning model training
- **AI Model Inference**: Model inference and prediction
- **Data Processing**: Large-scale data processing
- **Custom Workloads**: Custom container-based workloads

## Prerequisites

### Hardware Requirements

- **CPU**: 8+ cores, 3.0 GHz+
- **RAM**: 16GB+ (32GB recommended)
- **Storage**: 500GB+ SSD
- **GPU**: NVIDIA GPU with 8GB+ VRAM (for AI tasks)
- **Network**: 1 Gbps internet connection

### Software Requirements

- **Operating System**: Linux (Ubuntu 20.04+ recommended)
- **Go**: Version 1.21+
- **Docker**: Version 20.10+
- **Kubernetes**: Version 1.24+
- **NVIDIA Drivers**: Latest stable version

## Installation

### 1. Install Dependencies

```bash
# Install Go
wget -c https://golang.org/dl/go1.21.7.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install cri-dockerd (for Docker + Kubernetes)
# Follow instructions at: https://github.com/Mirantis/cri-dockerd
```

### 2. Install Kubernetes

#### Option 1: Docker + cri-dockerd (Recommended)

```bash
# Install cri-dockerd
git clone https://github.com/Mirantis/cri-dockerd.git
cd cri-dockerd
mkdir bin
go build -o bin/cri-dockerd cmd/cri-dockerd/main.go
sudo install -o root -g root -m 0755 bin/cri-dockerd /usr/bin/cri-dockerd
```

#### Option 2: Docker + Containerd

```bash
# Install containerd
# Follow instructions at: https://github.com/containerd/containerd/blob/main/docs/getting-started.md
```

### 3. Create Kubernetes Cluster

```bash
# Initialize Kubernetes cluster
sudo kubeadm init --pod-network-cidr=10.244.0.0/16

# Set up kubectl
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# Install network plugin (Calico)
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
```

### 4. Install NVIDIA Plugin

```bash
# Install NVIDIA device plugin
kubectl create -f https://raw.githubusercontent.com/NVIDIA/k8s-device-plugin/v0.14.1/nvidia-device-plugin.yml
```

### 5. Install Ingress Controller

```bash
# Install nginx-ingress
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/baremetal/deploy.yaml
```

### 6. Build Computing Provider

```bash
# Clone repository
git clone https://github.com/lagrangedao/go-computing-provider.git
cd go-computing-provider

# Build binary
make build
```

### 7. Initialize FCP

```bash
# Initialize repository
computing-provider init

# Set provider type to FCP
computing-provider account set-type --type fcp
```

## Configuration

### Basic FCP Configuration

```toml
# config.toml
[provider]
type = "fcp"
name = "my-fcp-provider"
description = "My Fog Computing Provider"

[fcp]
kubernetes_config = "/home/user/.kube/config"
namespace = "default"
resource_limits = { cpu = "4", memory = "8Gi" }
gpu_enabled = true
nvidia_runtime = "nvidia"

[fcp.storage]
type = "local"
path = "/data/storage"

[fcp.network]
ingress_enabled = true
ingress_class = "nginx"
```

### Kubernetes Configuration

```toml
[fcp.kubernetes]
config_path = "/home/user/.kube/config"
namespace = "default"
service_account = "computing-provider"
resource_quota = { cpu = "16", memory = "32Gi", gpu = "4" }
```

### Storage Configuration

```toml
[fcp.storage]
type = "local"  # or "nfs", "s3"
path = "/data/storage"
size = "500Gi"
access_mode = "ReadWriteOnce"
```

### Network Configuration

```toml
[fcp.network]
ingress_enabled = true
ingress_class = "nginx"
service_type = "NodePort"
external_ip = "YOUR_PUBLIC_IP"
```

## Setup Process

### 1. Wallet Setup

```bash
# Initialize wallet
computing-provider wallet init

# Set addresses
computing-provider account set-owner --address <OWNER_ADDRESS>
computing-provider account set-worker --address <WORKER_ADDRESS>
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

### 2. Network Configuration

```bash
# Set network (mainnet/testnet)
computing-provider network set --chain-id 1  # Mainnet
# or
computing-provider network set --chain-id 11155111  # Sepolia testnet
```

### 3. Task Type Configuration

```bash
# Enable FCP tasks
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 1
```

### 4. Collateral Setup

```bash
# Check collateral requirements
computing-provider collateral info --fcp

# Add collateral
computing-provider collateral add --fcp --owner <OWNER_ADDRESS> --amount <AMOUNT>
```

## Running FCP

### Start FCP Service

```bash
# Start the FCP service
export CP_PATH=~/.swan/computing
nohup computing-provider run >> cp.log 2>&1 &
```

### Monitor FCP

```bash
# Check FCP status
computing-provider state

# List FCP tasks
computing-provider task list --fcp

# Monitor logs
tail -f cp.log
```

## Task Management

### List Tasks

```bash
# List all FCP tasks
computing-provider task list --fcp

# List with details
computing-provider task list --fcp --verbose

# Show recent tasks
computing-provider task list --fcp --tail 10

# Show completed tasks
computing-provider task list --fcp --show-completed
```

### Task Details

```bash
# Get task details
computing-provider task get <job_uuid> --fcp
```

### Task Status

FCP tasks have the following statuses:

- **Pending**: Task waiting to be scheduled
- **Running**: Task currently executing
- **Completed**: Task finished successfully
- **Failed**: Task failed
- **Cancelled**: Task was cancelled
- **Timeout**: Task exceeded time limit

### Delete Tasks

```bash
# Delete a task
computing-provider task delete <job_uuid> --fcp

# Force delete
computing-provider task delete <job_uuid> --fcp --force
```

## Kubernetes Management

### Check Cluster Status

```bash
# Check node status
kubectl get nodes

# Check pod status
kubectl get pods --all-namespaces

# Check service status
kubectl get services --all-namespaces
```

### Resource Monitoring

```bash
# Check resource usage
kubectl top nodes
kubectl top pods --all-namespaces

# Check GPU usage
kubectl get nodes -o json | jq '.items[].status.allocatable."nvidia.com/gpu"'
```

### Troubleshooting Kubernetes

```bash
# Check pod logs
kubectl logs <pod-name> -n <namespace>

# Describe pod
kubectl describe pod <pod-name> -n <namespace>

# Check events
kubectl get events --all-namespaces --sort-by='.lastTimestamp'
```

## Performance Optimization

### Resource Allocation

```bash
# Monitor resource usage
htop
kubectl top nodes
kubectl top pods --all-namespaces

# Check GPU usage
nvidia-smi
kubectl get nodes -o json | jq '.items[].status.allocatable."nvidia.com/gpu"'
```

### Storage Optimization

```bash
# Monitor disk usage
df -h
kubectl get pv
kubectl get pvc --all-namespaces
```

### Network Optimization

```bash
# Check network connectivity
kubectl get ingress --all-namespaces
kubectl get services --all-namespaces
```

## Troubleshooting

### Common Issues

1. **Kubernetes Cluster Issues**
   ```bash
   # Check cluster status
   kubectl cluster-info
   
   # Check node status
   kubectl get nodes
   
   # Check system pods
   kubectl get pods -n kube-system
   ```

2. **GPU Not Available**
   ```bash
   # Check NVIDIA device plugin
   kubectl get pods -n kube-system | grep nvidia
   
   # Check GPU allocation
   kubectl get nodes -o json | jq '.items[].status.allocatable."nvidia.com/gpu"'
   ```

3. **Task Failures**
   ```bash
   # Check task logs
   computing-provider task get <job_uuid> --fcp
   
   # Check pod logs
   kubectl logs <pod-name> -n <namespace>
   
   # Check system logs
   tail -f cp.log
   ```

4. **Storage Issues**
   ```bash
   # Check persistent volumes
   kubectl get pv
   kubectl get pvc --all-namespaces
   
   # Check storage class
   kubectl get storageclass
   ```

### Debug Commands

```bash
# Check FCP configuration
computing-provider info

# Check wallet status
computing-provider wallet status

# Check collateral status
computing-provider collateral info --fcp

# Check network status
computing-provider network status
```

## Monitoring and Metrics

### System Metrics

```bash
# CPU usage
top
kubectl top nodes

# Memory usage
free -h
kubectl top nodes

# GPU usage
nvidia-smi
kubectl get nodes -o json | jq '.items[].status.allocatable."nvidia.com/gpu"'
```

### Provider Metrics

```bash
# Task completion rate
computing-provider task list --fcp --show-completed

# Earnings
computing-provider collateral info --fcp

# Resource utilization
kubectl top nodes
kubectl top pods --all-namespaces
```

## Security Considerations

### Kubernetes Security

- Use RBAC for access control
- Enable network policies
- Use secrets for sensitive data
- Keep Kubernetes updated

### System Security

- Keep system updated
- Use firewall rules
- Monitor system logs
- Use secure container images

## Exit Procedure

To exit as an FCP provider:

### Step 1: Set Exit Status

```bash
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 100
```

### Step 2: Withdraw Collateral

```bash
# Withdraw from Collateral account
computing-provider collateral withdraw --fcp --owner <OWNER_ADDRESS> --account <CP_ACCOUNT> <amount>

# Withdraw from Escrow account (requires 7-day waiting period)
computing-provider collateral withdraw-request --fcp --owner <OWNER_ADDRESS> <AMOUNT>

# After 7 days, confirm withdrawal
computing-provider collateral withdraw-confirm --fcp --owner <OWNER_ADDRESS>
```

## Support

- [Discord Community](https://discord.gg/Jd2BFSVCKw)
- [GitHub Issues](https://github.com/lagrangedao/go-computing-provider/issues)
- [Swan Chain Documentation](https://docs.swanchain.io)

## Related Documentation

- [Installation Guide](../installation.md)
- [Configuration](../configuration.md)
- [CLI Reference](../cli/README.md)
- [Task Management](../cli/task.md)
- [Wallet Management](../cli/wallet.md)
- [Kubernetes Setup](../kubernetes-setup.md) 