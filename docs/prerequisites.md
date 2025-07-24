# Prerequisites

Before setting up your Computing Provider, ensure your system meets all the necessary requirements.

## System Requirements

### Operating System
- **Linux**: Ubuntu 20.04+ (recommended)
- **Architecture**: x86_64/amd64
- **Kernel**: Linux kernel 5.4+

### Hardware Requirements

#### Minimum Requirements
- **CPU**: 4 cores, 2.0 GHz
- **RAM**: 8GB
- **Storage**: 100GB available space
- **Network**: 100 Mbps internet connection

#### Recommended Requirements
- **CPU**: 8+ cores, 3.0 GHz+
- **RAM**: 16GB+
- **Storage**: 500GB+ SSD
- **Network**: 1 Gbps internet connection
- **GPU**: NVIDIA GPU with CUDA support (for AI tasks)

### Software Dependencies

#### Required Software
- **Go**: Version 1.21+
- **Docker**: Version 20.10+ (for container runtime)
- **Kubernetes**: Version 1.24+ (for FCP)
- **NVIDIA Drivers**: Latest stable (for GPU support)

#### Optional Software
- **Nginx**: For reverse proxy
- **Certbot**: For SSL certificate management
- **Prometheus**: For monitoring

## Network Requirements

### Public IP Address
- A static public IP address is required
- Port forwarding may be necessary for certain configurations

### Domain Name
- A domain name in the format `*.example.com`
- DNS records properly configured

### SSL Certificate
- Valid SSL certificate for your domain
- Can be obtained from Let's Encrypt or other providers

## Kubernetes Requirements (FCP Only)

### Container Runtime
Choose one of the following:

#### Option 1: Docker + cri-dockerd (Recommended)
```bash
# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install cri-dockerd
# Follow instructions at: https://github.com/Mirantis/cri-dockerd
```

#### Option 2: Docker + Containerd
```bash
# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Install Containerd
# Follow instructions at: https://github.com/containerd/containerd/blob/main/docs/getting-started.md
```

### Kubernetes Cluster
- **Version**: 1.24.0+
- **Network Plugin**: Calico, Flannel, or Weave Net
- **NVIDIA Plugin**: For GPU support
- **Ingress Controller**: nginx-ingress

## Security Requirements

### Firewall Configuration
- **Inbound Ports**: 80, 443, 6443 (Kubernetes API)
- **Outbound**: All traffic allowed

### User Permissions
- **Docker Group**: User must be in docker group
- **Kubernetes**: Proper RBAC configuration
- **File Permissions**: Appropriate file ownership

## Environment Setup

### Environment Variables
```bash
# Set the computing provider path
export CP_PATH=~/.swan/computing

# Add to your shell profile
echo "export CP_PATH=~/.swan/computing" >> ~/.bashrc
```

### Directory Structure
```bash
# Create necessary directories
mkdir -p ~/.swan/computing
mkdir -p /docker_repo  # For Docker registry (optional)
```

## Verification Commands

### Check System Requirements
```bash
# Check Go version
go version

# Check Docker
docker --version

# Check Kubernetes (if installed)
kubectl version --client

# Check available memory
free -h

# Check available disk space
df -h

# Check network connectivity
ping -c 3 google.com
```

### Check GPU Support (if applicable)
```bash
# Check NVIDIA drivers
nvidia-smi

# Check CUDA installation
nvcc --version
```

## Pre-installation Checklist

- [ ] Go 1.21+ installed and in PATH
- [ ] Docker installed and running
- [ ] Kubernetes cluster configured (FCP only)
- [ ] Public IP address available
- [ ] Domain name configured
- [ ] SSL certificate obtained
- [ ] Firewall configured
- [ ] User permissions set correctly
- [ ] Environment variables configured
- [ ] Required directories created

## Next Steps

Once all prerequisites are met:

1. [Install the Computing Provider](installation.md)
2. [Configure your environment](configuration.md)
3. [Set up your wallet](wallet.md) 