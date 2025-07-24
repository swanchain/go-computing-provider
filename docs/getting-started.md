# Getting Started

This guide will walk you through the complete setup process for your Computing Provider, from installation to running your first tasks.

## Quick Start Overview

1. [Install Dependencies](prerequisites.md)
2. [Install Computing Provider](installation.md)
3. [Configure Your Environment](configuration.md)
4. [Set Up Your Wallet](wallet.md)
5. [Choose Provider Type](#choose-provider-type)
6. [Start Your Provider](#start-your-provider)

## Choose Provider Type

The Computing Provider supports two types of providers:

### Edge Computing Provider (ECP)
- **Best for**: ZK-SNARK proof generation, low-latency tasks
- **Hardware**: GPU with 8GB+ VRAM, 8GB+ RAM
- **Setup**: Simpler setup, no Kubernetes required
- **Tasks**: Filecoin C2 proofs, future ZK proof types

### Fog Computing Provider (FCP)
- **Best for**: AI training, data processing, scalable workloads
- **Hardware**: 16GB+ RAM, 500GB+ storage, GPU recommended
- **Setup**: Requires Kubernetes cluster
- **Tasks**: AI model training, inference, custom workloads

## Initial Setup

### 1. Initialize Repository

```bash
# Initialize computing provider repository
computing-provider init

# Verify initialization
ls -la ~/.swan/computing/
```

### 2. Set Up Wallet

```bash
# Initialize wallet
computing-provider wallet init

# Import existing wallet (optional)
computing-provider wallet import --private-key <PRIVATE_KEY>

# List wallets
computing-provider wallet list
```

### 3. Configure Addresses

```bash
# Set owner address (controls account settings)
computing-provider account set-owner --address <OWNER_ADDRESS>

# Set worker address (pays gas fees, submits proofs)
computing-provider account set-worker --address <WORKER_ADDRESS>

# Set beneficiary address (receives earnings)
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

### 4. Configure Network

```bash
# For mainnet
computing-provider network set --chain-id 1

# For testnet (Sepolia)
computing-provider network set --chain-id 11155111
```

## Provider-Specific Setup

### ECP Setup

```bash
# Set provider type to ECP
computing-provider account set-type --type ecp

# Enable Filecoin C2 tasks
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 1

# Download UBI parameters (required for ZK proofs)
cd ubi
./setup.sh
```

### FCP Setup

```bash
# Set provider type to FCP
computing-provider account set-type --type fcp

# Enable FCP tasks
computing-provider account changeTaskTypes --ownerAddress <OWNER_ADDRESS> 1

# Verify Kubernetes cluster
kubectl cluster-info
kubectl get nodes
```

## Collateral Setup

Both provider types require collateral to participate in the network.

### Check Requirements

```bash
# For ECP
computing-provider collateral info --ecp

# For FCP
computing-provider collateral info --fcp
```

### Add Collateral

```bash
# For ECP
computing-provider collateral add --ecp --owner <OWNER_ADDRESS> --amount <AMOUNT>

# For FCP
computing-provider collateral add --fcp --owner <OWNER_ADDRESS> --amount <AMOUNT>
```

## Start Your Provider

### 1. Verify Configuration

```bash
# Check provider information
computing-provider info

# Check provider state
computing-provider state

# Check wallet status
computing-provider wallet status
```

### 2. Start the Service

```bash
# Set environment variable
export CP_PATH=~/.swan/computing

# Start the provider
nohup computing-provider run >> cp.log 2>&1 &

# Check if it's running
ps aux | grep computing-provider
```

### 3. Monitor Your Provider

```bash
# Check logs
tail -f cp.log

# Check status
computing-provider state

# List tasks
computing-provider task list --ecp  # or --fcp
```

## First Tasks

### ECP First Task

1. **Wait for UBI Tasks**: ECP providers automatically receive UBI (Universal Basic Income) tasks
2. **Monitor Task Status**: Check task list for new tasks
3. **Verify Completion**: Tasks are automatically verified on-chain

```bash
# Monitor for new tasks
watch -n 30 'computing-provider task list --ecp --tail 5'

# Check task details
computing-provider task get <job_uuid> --ecp
```

### FCP First Task

1. **Register with Orchestrator**: FCP providers register with the Swan Orchestrator
2. **Receive Tasks**: Tasks are assigned based on your capabilities
3. **Execute Tasks**: Kubernetes manages task execution

```bash
# Monitor for new tasks
watch -n 30 'computing-provider task list --fcp --tail 5'

# Check task details
computing-provider task get <job_uuid> --fcp
```

## Verification and Monitoring

### Check Provider Status

```bash
# Provider information
computing-provider info

# Provider state
computing-provider state

# Network status
computing-provider network status
```

### Monitor Performance

```bash
# System resources
htop
nvidia-smi  # if using GPU

# Task completion
computing-provider task list --ecp --show-completed  # or --fcp

# Earnings
computing-provider collateral info --ecp  # or --fcp
```

### Check Logs

```bash
# Provider logs
tail -f cp.log

# System logs
journalctl -u computing-provider -f

# Kubernetes logs (FCP only)
kubectl logs -f deployment/computing-provider
```

## Troubleshooting

### Common Startup Issues

1. **Configuration Errors**
   ```bash
   # Validate configuration
   computing-provider config validate
   
   # Check configuration
   computing-provider config show
   ```

2. **Wallet Issues**
   ```bash
   # Check wallet status
   computing-provider wallet status
   
   # Verify addresses
   computing-provider account info
   ```

3. **Network Issues**
   ```bash
   # Check network connectivity
   computing-provider network status
   
   # Test RPC endpoint
   curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' <RPC_URL>
   ```

4. **Resource Issues**
   ```bash
   # Check system resources
   free -h
   df -h
   nvidia-smi  # if using GPU
   ```

### Getting Help

- [Discord Community](https://discord.gg/Jd2BFSVCKw)
- [GitHub Issues](https://github.com/lagrangedao/go-computing-provider/issues)
- [Troubleshooting Guide](troubleshooting.md)

## Next Steps

After successfully starting your provider:

1. **Monitor Performance**: Track task completion rates and earnings
2. **Optimize Resources**: Adjust resource allocation based on usage
3. **Scale Up**: Consider adding more hardware for increased capacity
4. **Join Community**: Connect with other providers on Discord

## Provider Status Dashboard

Check your provider status at: [https://provider.swanchain.io](https://provider.swanchain.io)

This dashboard shows:
- Provider registration status
- Task completion statistics
- Earnings and collateral information
- Network health metrics 