# Run Command

The `run` command starts the Computing Provider service.

## Overview

```bash
computing-provider run [flags]
```

## Description

The `run` command starts the Computing Provider service, which:

- Connects to the Swan Chain network
- Registers the provider with the network
- Begins accepting and processing tasks
- Manages wallet operations and collateral
- Provides monitoring and metrics endpoints

## Flags

- `--config <path>`: Path to configuration file (default: `~/.swan/computing/config.toml`)
- `--log-level <level>`: Log level (debug, info, warn, error) (default: `info`)
- `--daemon`: Run as daemon process
- `--pid-file <path>`: Path to PID file when running as daemon

## Examples

### Basic Usage

```bash
# Start the provider
computing-provider run

# Start with custom config
computing-provider run --config /path/to/config.toml

# Start with debug logging
computing-provider run --log-level debug
```

### Running as Daemon

```bash
# Run as daemon process
computing-provider run --daemon

# Run as daemon with PID file
computing-provider run --daemon --pid-file /var/run/computing-provider.pid
```

### Using Environment Variables

```bash
# Set environment variables
export CP_PATH=~/.swan/computing
export CP_LOG_LEVEL=debug

# Start provider
computing-provider run
```

## Service Management

### Start Service

```bash
# Simple start
computing-provider run

# Background start
nohup computing-provider run >> cp.log 2>&1 &

# Using systemd (if configured)
sudo systemctl start computing-provider
```

### Stop Service

```bash
# Graceful stop
pkill computing-provider

# Force stop
pkill -9 computing-provider

# Using systemd
sudo systemctl stop computing-provider
```

### Restart Service

```bash
# Manual restart
pkill computing-provider
nohup computing-provider run >> cp.log 2>&1 &

# Using systemd
sudo systemctl restart computing-provider
```

## Monitoring

### Check Service Status

```bash
# Check if running
ps aux | grep computing-provider

# Check process ID
pgrep computing-provider

# Check port usage
netstat -tuln | grep 8080
```

### Monitor Logs

```bash
# Follow logs in real-time
tail -f cp.log

# Search for errors
grep -i error cp.log

# Search for specific patterns
grep "task" cp.log
```

### Health Checks

```bash
# Check provider status
computing-provider state

# Check provider info
computing-provider info

# Check network connectivity
computing-provider network status
```

## Configuration

### Required Configuration

Before running, ensure you have:

1. **Initialized Repository**
   ```bash
   computing-provider init
   ```

2. **Set Up Wallet**
   ```bash
   computing-provider wallet init
   computing-provider account set-owner --address <OWNER_ADDRESS>
   computing-provider account set-worker --address <WORKER_ADDRESS>
   computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
   ```

3. **Configured Network**
   ```bash
   computing-provider network set --chain-id 1  # Mainnet
   ```

4. **Added Collateral**
   ```bash
   computing-provider collateral add --ecp --owner <OWNER_ADDRESS> --amount <AMOUNT>
   # or
   computing-provider collateral add --fcp --owner <OWNER_ADDRESS> --amount <AMOUNT>
   ```

### Configuration File

The service reads configuration from `~/.swan/computing/config.toml`:

```toml
[server]
host = "0.0.0.0"
port = 8080
ssl_enabled = false

[provider]
type = "ecp"  # or "fcp"
name = "my-provider"

[network]
chain_id = 1
rpc_url = "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"

[wallet]
owner_address = "0x..."
worker_address = "0x..."
beneficiary_address = "0x..."
```

## Startup Process

### Initialization Steps

1. **Load Configuration**: Read and validate configuration file
2. **Initialize Database**: Set up local database for task tracking
3. **Connect to Network**: Establish connection to Swan Chain
4. **Register Provider**: Register provider with the network
5. **Start Task Manager**: Begin accepting and processing tasks
6. **Start Monitoring**: Begin collecting metrics and health data

### Startup Logs

```bash
# Monitor startup process
tail -f cp.log

# Expected startup sequence:
# INFO: Loading configuration
# INFO: Initializing database
# INFO: Connecting to network
# INFO: Provider registered successfully
# INFO: Task manager started
# INFO: Service ready
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

2. **Network Issues**
   ```bash
   # Check network connectivity
   computing-provider network status
   
   # Test RPC endpoint
   curl -X POST -H "Content-Type: application/json" \
     --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
     <RPC_URL>
   ```

3. **Wallet Issues**
   ```bash
   # Check wallet status
   computing-provider wallet status
   
   # Verify addresses
   computing-provider account info
   ```

4. **Permission Issues**
   ```bash
   # Check file permissions
   ls -la ~/.swan/computing/
   
   # Fix permissions
   chmod -R 755 ~/.swan/computing/
   ```

### Debug Mode

```bash
# Run with debug logging
computing-provider run --log-level debug

# This provides detailed information about:
# - Configuration loading
# - Network connections
# - Task processing
# - Error details
```

## Performance

### Resource Usage

Monitor resource usage during startup:

```bash
# Monitor CPU and memory
top -p $(pgrep computing-provider)

# Monitor disk usage
df -h

# Monitor network
iftop
```

### Optimization

```bash
# Adjust log level for performance
computing-provider run --log-level warn

# Use dedicated user for better security
sudo -u computing-provider computing-provider run

# Set resource limits
ulimit -n 65536  # Increase file descriptor limit
```

## Integration

### Systemd Service

Create a systemd service file for automatic management:

```ini
# /etc/systemd/system/computing-provider.service
[Unit]
Description=Computing Provider Service
After=network.target

[Service]
Type=simple
User=computing-provider
Environment=CP_PATH=/home/computing-provider/.swan/computing
ExecStart=/usr/local/bin/computing-provider run
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
```

### Docker Integration

```dockerfile
# Dockerfile
FROM golang:1.21-alpine

COPY computing-provider /usr/local/bin/
COPY config.toml /root/.swan/computing/

EXPOSE 8080

CMD ["computing-provider", "run"]
```

## Related Commands

- [`init`](init.md) - Initialize provider repository
- [`state`](state.md) - Check provider state
- [`info`](info.md) - Display provider information
- [`task`](task.md) - Manage tasks
- [`wallet`](wallet.md) - Manage wallet operations 