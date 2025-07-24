# Troubleshooting Guide

This guide helps you resolve common issues when running the Go Computing Provider.

## Quick Diagnostic Commands

```bash
# Check provider status
computing-provider state

# Check configuration
computing-provider info

# Check wallet status
computing-provider wallet status

# Check network connectivity
computing-provider network status

# Check system resources
htop
free -h
df -h
nvidia-smi  # if using GPU
```

## Common Issues

### 1. Provider Won't Start

#### Symptoms
- Provider fails to start
- Error messages about missing configuration
- Permission denied errors

#### Solutions

**Check Repository Path**
```bash
# Verify CP_PATH is set
echo $CP_PATH

# Check if repository exists
ls -la ~/.swan/computing/

# Initialize if missing
computing-provider init
```

**Check Configuration**
```bash
# Validate configuration
computing-provider config validate

# Check configuration file
cat ~/.swan/computing/config.toml

# Regenerate configuration if corrupted
cp config.toml.sample ~/.swan/computing/config.toml
```

**Check Permissions**
```bash
# Fix repository permissions
chmod -R 755 ~/.swan/computing/

# Check file ownership
ls -la ~/.swan/computing/
```

### 2. Wallet Issues

#### Symptoms
- "Wallet not found" errors
- "Invalid private key" errors
- Balance showing as zero

#### Solutions

**Check Wallet Status**
```bash
# List wallets
computing-provider wallet list

# Check wallet status
computing-provider wallet status

# Verify addresses
computing-provider account info
```

**Reinitialize Wallet**
```bash
# Backup existing wallet (if needed)
cp -r ~/.swan/computing/keystore ~/.swan/computing/keystore.backup

# Reinitialize wallet
computing-provider wallet init

# Set addresses again
computing-provider account set-owner --address <OWNER_ADDRESS>
computing-provider account set-worker --address <WORKER_ADDRESS>
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

**Check Network Configuration**
```bash
# Verify network settings
computing-provider network status

# Test RPC endpoint
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  <RPC_URL>
```

### 3. Network Connectivity Issues

#### Symptoms
- "Connection refused" errors
- Timeout errors
- RPC endpoint not responding

#### Solutions

**Check Network Configuration**
```bash
# Test internet connectivity
ping -c 3 google.com

# Test RPC endpoint
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  <RPC_URL>

# Check firewall settings
sudo ufw status
```

**Update RPC Endpoint**
```bash
# Set new RPC endpoint
computing-provider network set --rpc-url <NEW_RPC_URL>

# Test new endpoint
computing-provider network status
```

**Check DNS Resolution**
```bash
# Test DNS
nslookup <RPC_HOSTNAME>

# Use IP address if DNS fails
# Update config.toml with IP address
```

### 4. Task Execution Issues

#### ECP Issues

**UBI Parameters Missing**
```bash
# Check UBI parameters
ls -la ubi/

# Download parameters
cd ubi
./setup.sh

# Verify parameters
ls -la *.params
```

**GPU Not Detected**
```bash
# Check NVIDIA drivers
nvidia-smi

# Check CUDA installation
nvcc --version

# Check GPU availability
lspci | grep -i nvidia
```

**Task Failures**
```bash
# Check task details
computing-provider task get <job_uuid> --ecp

# Check system logs
tail -f cp.log

# Check resource usage
htop
nvidia-smi
```

#### FCP Issues

**Kubernetes Cluster Issues**
```bash
# Check cluster status
kubectl cluster-info

# Check node status
kubectl get nodes

# Check system pods
kubectl get pods -n kube-system

# Check events
kubectl get events --all-namespaces --sort-by='.lastTimestamp'
```

**GPU Not Available in Kubernetes**
```bash
# Check NVIDIA device plugin
kubectl get pods -n kube-system | grep nvidia

# Check GPU allocation
kubectl get nodes -o json | jq '.items[].status.allocatable."nvidia.com/gpu"'

# Reinstall NVIDIA plugin if needed
kubectl delete -f https://raw.githubusercontent.com/NVIDIA/k8s-device-plugin/v0.14.1/nvidia-device-plugin.yml
kubectl create -f https://raw.githubusercontent.com/NVIDIA/k8s-device-plugin/v0.14.1/nvidia-device-plugin.yml
```

**Pod Scheduling Issues**
```bash
# Check pod status
kubectl get pods --all-namespaces

# Check pod events
kubectl describe pod <pod-name> -n <namespace>

# Check resource quotas
kubectl get resourcequota --all-namespaces
```

### 5. Resource Exhaustion

#### Symptoms
- Tasks stuck in pending
- System becomes unresponsive
- Out of memory errors

#### Solutions

**Check Resource Usage**
```bash
# Monitor system resources
htop
free -h
df -h

# Check Kubernetes resources
kubectl top nodes
kubectl top pods --all-namespaces
```

**Adjust Resource Limits**
```bash
# Update configuration
# Edit ~/.swan/computing/config.toml

[fcp]
resource_limits = { cpu = "2", memory = "4Gi" }  # Reduce limits

# Restart provider
pkill computing-provider
nohup computing-provider run >> cp.log 2>&1 &
```

**Clean Up Resources**
```bash
# Delete completed tasks
computing-provider task list --fcp --show-completed
computing-provider task delete <job_uuid> --fcp

# Clean up Kubernetes resources
kubectl delete pods --field-selector=status.phase=Succeeded --all-namespaces
kubectl delete pods --field-selector=status.phase=Failed --all-namespaces
```

### 6. Collateral Issues

#### Symptoms
- "Insufficient collateral" errors
- Cannot add collateral
- Withdrawal failures

#### Solutions

**Check Collateral Status**
```bash
# Check collateral info
computing-provider collateral info --ecp  # or --fcp

# Check account balance
computing-provider wallet balance <WORKER_ADDRESS>
```

**Add Collateral**
```bash
# Add collateral
computing-provider collateral add --ecp --owner <OWNER_ADDRESS> --amount <AMOUNT>

# Verify addition
computing-provider collateral info --ecp
```

**Withdrawal Issues**
```bash
# Check withdrawal status
computing-provider collateral info --ecp

# Request withdrawal (7-day waiting period)
computing-provider collateral withdraw-request --ecp --owner <OWNER_ADDRESS> <AMOUNT>

# Confirm withdrawal after 7 days
computing-provider collateral withdraw-confirm --ecp --owner <OWNER_ADDRESS>
```

## Performance Issues

### Slow Task Execution

**Check System Performance**
```bash
# Monitor CPU usage
top -p $(pgrep computing-provider)

# Monitor memory usage
free -h

# Monitor disk I/O
iotop

# Monitor network
iftop
```

**Optimize Configuration**
```bash
# Increase resource limits
# Edit ~/.swan/computing/config.toml

[fcp]
resource_limits = { cpu = "8", memory = "16Gi" }

# Restart provider
pkill computing-provider
nohup computing-provider run >> cp.log 2>&1 &
```

### High Resource Usage

**Identify Resource Hogs**
```bash
# Find processes using most CPU
ps aux --sort=-%cpu | head -10

# Find processes using most memory
ps aux --sort=-%mem | head -10

# Check disk usage
du -sh /* | sort -hr | head -10
```

**Optimize Resource Allocation**
```bash
# Adjust Kubernetes resource quotas
kubectl patch resourcequota <quota-name> -p '{"spec":{"hard":{"cpu":"8","memory":"16Gi"}}}'

# Update provider configuration
# Edit ~/.swan/computing/config.toml
```

## Log Analysis

### Understanding Logs

**Provider Logs**
```bash
# View recent logs
tail -f cp.log

# Search for errors
grep -i error cp.log

# Search for warnings
grep -i warning cp.log

# Search for specific task
grep <job_uuid> cp.log
```

**System Logs**
```bash
# Check system logs
journalctl -u computing-provider -f

# Check kernel logs
dmesg | tail -20

# Check network logs
journalctl -u systemd-networkd -f
```

**Kubernetes Logs (FCP)**
```bash
# Check pod logs
kubectl logs <pod-name> -n <namespace>

# Check service logs
kubectl logs -l app=computing-provider -n <namespace>

# Check ingress logs
kubectl logs -n ingress-nginx deployment/ingress-nginx-controller
```

## Recovery Procedures

### Complete Reset

**Backup Important Data**
```bash
# Backup configuration
cp -r ~/.swan/computing ~/.swan/computing.backup

# Backup wallet
cp -r ~/.swan/computing/keystore ~/.swan/computing/keystore.backup
```

**Reset Provider**
```bash
# Stop provider
pkill computing-provider

# Remove repository
rm -rf ~/.swan/computing

# Reinitialize
computing-provider init
computing-provider wallet init

# Restore configuration
cp ~/.swan/computing.backup/config.toml ~/.swan/computing/
cp -r ~/.swan/computing.backup/keystore ~/.swan/computing/
```

### Kubernetes Reset (FCP)

**Reset Kubernetes Cluster**
```bash
# Reset cluster
sudo kubeadm reset

# Reinitialize cluster
sudo kubeadm init --pod-network-cidr=10.244.0.0/16

# Reinstall components
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
kubectl create -f https://raw.githubusercontent.com/NVIDIA/k8s-device-plugin/v0.14.1/nvidia-device-plugin.yml
```

## Getting Help

### Before Asking for Help

1. **Collect Information**
   ```bash
   # System information
   uname -a
   cat /etc/os-release
   
   # Provider information
   computing-provider info
   computing-provider state
   
   # Logs
   tail -100 cp.log
   ```

2. **Document the Issue**
   - What were you trying to do?
   - What error messages did you see?
   - What steps did you take?
   - What is your system configuration?

### Support Channels

- **Discord**: [Swan Chain Community](https://discord.gg/Jd2BFSVCKw)
- **GitHub**: [Issue Tracker](https://github.com/lagrangedao/go-computing-provider/issues)
- **Documentation**: [Swan Chain Docs](https://docs.swanchain.io)

### Useful Commands for Support

```bash
# Generate debug information
computing-provider info > debug_info.txt
computing-provider state >> debug_info.txt
tail -100 cp.log >> debug_info.txt
nvidia-smi >> debug_info.txt  # if using GPU
kubectl get nodes -o wide >> debug_info.txt  # if using FCP
```

## Prevention

### Regular Maintenance

1. **Monitor System Resources**
   ```bash
   # Set up monitoring
   watch -n 30 'htop -n 1'
   watch -n 30 'nvidia-smi'
   ```

2. **Regular Backups**
   ```bash
   # Backup configuration
   cp -r ~/.swan/computing ~/.swan/computing.backup.$(date +%Y%m%d)
   ```

3. **Update Software**
   ```bash
   # Update system
   sudo apt update && sudo apt upgrade
   
   # Update provider
   git pull
   make build
   ```

4. **Check Logs Regularly**
   ```bash
   # Monitor logs
   tail -f cp.log
   ``` 