# Configuration

This guide covers all configuration aspects of the Go Computing Provider.

## Configuration Files

The Computing Provider uses several configuration files located in your repository directory (`~/.swan/computing` by default).

### Main Configuration File

The main configuration file is `config.toml`. You can use the sample configuration as a starting point:

```bash
# Copy the sample configuration
cp config.toml.sample ~/.swan/computing/config.toml
```

## Configuration Structure

### Basic Configuration

```toml
# config.toml
[server]
host = "0.0.0.0"
port = 8080
ssl_enabled = true
ssl_cert = "/path/to/cert.pem"
ssl_key = "/path/to/key.pem"

[provider]
type = "fcp"  # or "ecp"
name = "my-computing-provider"
description = "My computing provider description"

[network]
chain_id = 1
rpc_url = "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
ws_url = "wss://mainnet.infura.io/ws/v3/YOUR_PROJECT_ID"

[wallet]
owner_address = "0x..."
worker_address = "0x..."
beneficiary_address = "0x..."
```

### Provider-Specific Configuration

#### FCP Configuration

```toml
[fcp]
kubernetes_config = "/path/to/kubeconfig"
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

#### ECP Configuration

```toml
[ecp]
task_types = ["fil-c2"]  # Supported task types
gpu_requirements = { memory = "8Gi", compute_capability = "7.0" }

[ecp.ubi]
enabled = true
param_path = "/path/to/ubi/params"
```

## Environment Variables

You can override configuration values using environment variables:

```bash
# Server configuration
export CP_SERVER_HOST="0.0.0.0"
export CP_SERVER_PORT="8080"
export CP_SERVER_SSL_ENABLED="true"

# Provider configuration
export CP_PROVIDER_TYPE="fcp"
export CP_PROVIDER_NAME="my-provider"

# Network configuration
export CP_NETWORK_CHAIN_ID="1"
export CP_NETWORK_RPC_URL="https://mainnet.infura.io/v3/YOUR_PROJECT_ID"

# Wallet configuration
export CP_WALLET_OWNER_ADDRESS="0x..."
export CP_WALLET_WORKER_ADDRESS="0x..."
export CP_WALLET_BENEFICIARY_ADDRESS="0x..."
```

## Wallet Configuration

### Address Types

The Computing Provider uses three different wallet addresses for security:

1. **Owner Address**: Controls account settings and permissions
2. **Worker Address**: Used for submitting proofs and paying gas fees
3. **Beneficiary Address**: Receives all earnings

### Setting Up Wallets

```bash
# Initialize wallet
computing-provider wallet init

# Import existing wallet
computing-provider wallet import --private-key <PRIVATE_KEY>

# Set addresses
computing-provider account set-owner --address <OWNER_ADDRESS>
computing-provider account set-worker --address <WORKER_ADDRESS>
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

## Network Configuration

### Supported Networks

- **Mainnet**: Chain ID 1
- **Testnet**: Chain ID 11155111 (Sepolia)
- **Local**: Chain ID 1337

### RPC Configuration

```toml
[network]
chain_id = 1
rpc_url = "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
ws_url = "wss://mainnet.infura.io/ws/v3/YOUR_PROJECT_ID"
gas_limit = 3000000
gas_price = "20000000000"  # 20 Gwei
```

## Resource Configuration

### Hardware Resources

```toml
[resources]
cpu_cores = 8
memory_gb = 16
storage_gb = 500
gpu_count = 1
gpu_memory_gb = 8
```

### Resource Limits

```toml
[limits]
max_concurrent_tasks = 10
max_cpu_per_task = "4"
max_memory_per_task = "8Gi"
max_gpu_per_task = 1
```

## Security Configuration

### SSL/TLS Configuration

```toml
[ssl]
enabled = true
cert_file = "/path/to/cert.pem"
key_file = "/path/to/key.pem"
ca_file = "/path/to/ca.pem"  # Optional
```

### Authentication

```toml
[auth]
enabled = true
jwt_secret = "your-jwt-secret"
session_timeout = 3600  # seconds
```

## Monitoring Configuration

### Metrics

```toml
[monitoring]
enabled = true
metrics_port = 9090
prometheus_enabled = true
health_check_interval = 30  # seconds
```

### Logging

```toml
[logging]
level = "info"  # debug, info, warn, error
format = "json"  # json, text
file = "/var/log/computing-provider.log"
max_size = 100  # MB
max_backups = 3
```

## Validation

Validate your configuration:

```bash
# Validate configuration file
computing-provider config validate

# Check configuration
computing-provider config show
```

## Configuration Examples

### Minimal FCP Configuration

```toml
[server]
host = "0.0.0.0"
port = 8080

[provider]
type = "fcp"
name = "my-fcp"

[network]
chain_id = 1
rpc_url = "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"

[wallet]
owner_address = "0x..."
worker_address = "0x..."
beneficiary_address = "0x..."
```

### Minimal ECP Configuration

```toml
[server]
host = "0.0.0.0"
port = 8080

[provider]
type = "ecp"
name = "my-ecp"

[network]
chain_id = 1
rpc_url = "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"

[wallet]
owner_address = "0x..."
worker_address = "0x..."
beneficiary_address = "0x..."

[ecp]
task_types = ["fil-c2"]
```

## Troubleshooting Configuration

### Common Issues

1. **Invalid TOML syntax**: Use a TOML validator
2. **Missing required fields**: Check the sample configuration
3. **Permission errors**: Ensure proper file permissions
4. **Network connectivity**: Verify RPC endpoints

### Configuration Commands

```bash
# Show current configuration
computing-provider config show

# Validate configuration
computing-provider config validate

# Generate sample configuration
computing-provider config generate
```

## Next Steps

After configuring your Computing Provider:

1. [Set up your wallet](wallet.md)
2. [Initialize your account](getting-started.md)
3. [Start the provider](getting-started.md) 