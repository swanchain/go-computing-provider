# Edge Computing Provider (ECP)

Edge Computing Provider (ECP) specializes in processing data at the source of data generation, using minimal latency setups ideal for real-time applications. Currently supports ZK-SNARK proof generation for the Filecoin network.

## Overview

ECP handles specific, localized tasks directly on devices at the network's edge, such as IoT devices. It's designed for low-latency, real-time processing requirements.

### Current Supported Tasks

- **Filecoin C2 ZK-SNARK Proofs**: Generation of zero-knowledge proofs for Filecoin network
- **Future Support**: Aleo, Scroll, StarkNet, and other ZK proof types

## Prerequisites

### Hardware Requirements

- **CPU**: 4+ cores, 2.0 GHz+
- **RAM**: 8GB+ (16GB recommended)
- **Storage**: 100GB+ available space
- **GPU**: NVIDIA GPU with 8GB+ VRAM (for ZK proof generation)
- **Network**: Stable internet connection

### Software Requirements

- **Operating System**: Linux (Ubuntu 20.04+ recommended)
- **Go**: Version 1.21+
- **NVIDIA Drivers**: Latest stable version
- **CUDA**: Version 11.0+ (for GPU acceleration)

## Installation

### 1. Install Dependencies

```bash
# Install Go
wget -c https://golang.org/dl/go1.21.7.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc

# Install NVIDIA drivers and CUDA
# Follow NVIDIA's official installation guide for your GPU
```

### 2. Build Computing Provider

```bash
# Clone repository
git clone https://github.com/lagrangedao/go-computing-provider.git
cd go-computing-provider

# Build binary
make build
```

### 3. Initialize ECP

```bash
# Initialize repository
computing-provider init

# Set provider type to ECP
computing-provider account set-type --type ecp
```

## Configuration

### Basic ECP Configuration

```toml
# config.toml
[provider]
type = "ecp"
name = "my-ecp-provider"
description = "My Edge Computing Provider"

[ecp]
task_types = ["fil-c2"]
gpu_requirements = { memory = "8Gi", compute_capability = "7.0" }

[ecp.ubi]
enabled = true
param_path = "/path/to/ubi/params"
```

### UBI Configuration

For Filecoin C2 ZK-SNARK proofs, you need to download UBI parameters:

```bash
# Download UBI parameters
cd ubi
./setup.sh

# Or manually download parameters
./fetch-param-32.sh
./fetch-param-512.sh
```

### GPU Configuration

```toml
[ecp.gpu]
enabled = true
nvidia_runtime = "nvidia"
memory_limit = "8Gi"
compute_capability = "7.0"
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
# Enable Filecoin C2 tasks
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 1
```

### 4. Collateral Setup

```bash
# Check collateral requirements
computing-provider collateral info --ecp

# Add collateral
computing-provider collateral add --ecp --owner <OWNER_ADDRESS> --amount <AMOUNT>
```

## Running ECP

### Start ECP Service

```bash
# Start the ECP service
export CP_PATH=~/.swan/computing
nohup computing-provider run >> cp.log 2>&1 &
```

### Monitor ECP

```bash
# Check ECP status
computing-provider state

# List ECP tasks
computing-provider task list --ecp

# Monitor logs
tail -f cp.log
```

## Task Management

### List Tasks

```bash
# List all ECP tasks
computing-provider task list --ecp

# List with details
computing-provider task list --ecp --verbose

# Show recent tasks
computing-provider task list --ecp --tail 10
```

### Task Details

```bash
# Get task details
computing-provider task get <job_uuid> --ecp
```

### Task Status

ECP tasks have the following statuses:

- **Submitted**: Task submitted to network
- **Processing**: Task being processed
- **Completed**: Task completed successfully
- **Failed**: Task failed
- **Verified**: Task verified on-chain

## Performance Optimization

### GPU Optimization

```bash
# Check GPU status
nvidia-smi

# Monitor GPU usage
watch -n 1 nvidia-smi

# Check CUDA availability
nvcc --version
```

### Resource Monitoring

```bash
# Monitor system resources
htop

# Monitor disk usage
df -h

# Monitor network
iftop
```

## Troubleshooting

### Common Issues

1. **GPU Not Detected**
   ```bash
   # Check NVIDIA drivers
   nvidia-smi
   
   # Check CUDA installation
   nvcc --version
   ```

2. **UBI Parameters Missing**
   ```bash
   # Download parameters
   cd ubi
   ./setup.sh
   ```

3. **Task Failures**
   ```bash
   # Check task logs
   computing-provider task get <job_uuid> --ecp
   
   # Check system logs
   tail -f cp.log
   ```

4. **Network Issues**
   ```bash
   # Check network connectivity
   computing-provider network status
   
   # Check RPC endpoint
   curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' <RPC_URL>
   ```

### Debug Commands

```bash
# Check ECP configuration
computing-provider info

# Check wallet status
computing-provider wallet status

# Check collateral status
computing-provider collateral info --ecp
```

## Monitoring and Metrics

### System Metrics

```bash
# CPU usage
top

# Memory usage
free -h

# GPU usage
nvidia-smi -l 1
```

### Provider Metrics

```bash
# Task completion rate
computing-provider task list --ecp --show-completed

# Earnings
computing-provider collateral info --ecp
```

## Security Considerations

### Wallet Security

- Keep private keys secure
- Use separate addresses for different purposes
- Regularly backup keystore files

### System Security

- Keep system updated
- Use firewall rules
- Monitor system logs

## Exit Procedure

To exit as an ECP provider:

### Step 1: Set Exit Status

```bash
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 100
```

### Step 2: Withdraw Collateral

```bash
# Withdraw from Collateral account
computing-provider collateral withdraw --ecp --owner <OWNER_ADDRESS> --account <CP_ACCOUNT> <amount>

# Withdraw from Escrow account (requires 7-day waiting period)
computing-provider collateral withdraw-request --ecp --owner <OWNER_ADDRESS> <AMOUNT>

# After 7 days, confirm withdrawal
computing-provider collateral withdraw-confirm --ecp --owner <OWNER_ADDRESS>
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