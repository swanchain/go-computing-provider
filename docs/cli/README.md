# Command Line Interface (CLI)

The Go Computing Provider provides a comprehensive command-line interface for managing all aspects of your computing provider operations.

## Overview

The CLI is accessed through the `computing-provider` command and provides subcommands for different operations:

```bash
computing-provider [global-flags] <command> [command-flags] [arguments]
```

## Global Flags

- `--repo <path>`: Repository directory for computing-provider client (default: `~/.swan/computing`)
- `--version`: Show version information
- `--help`: Show help information

## Commands Overview

### Core Commands
- [`init`](init.md) - Initialize a new computing provider repository
- [`run`](run.md) - Start the computing provider service
- [`info`](info.md) - Display provider information
- [`state`](state.md) - Show provider state

### Account Management
- [`account`](account.md) - Manage provider account settings
- [`wallet`](wallet.md) - Manage wallet operations

### Task Management
- [`task`](task.md) - Manage computing tasks
- [`ubi-task`](ubi-task.md) - Manage UBI tasks

### Financial Operations
- [`collateral`](collateral.md) - Manage provider collateral
- [`price`](price.md) - Manage pricing settings

### Network Operations
- [`network`](network.md) - Network configuration and management
- [`sequencer`](sequencer.md) - Sequencer operations

### Contract Operations
- [`contract`](contract.md) - Smart contract interactions

## Environment Variables

The CLI respects the following environment variables:

- `CP_PATH`: Repository directory path
- `CP_NETWORK`: Network configuration
- `CP_WALLET`: Wallet configuration

## Examples

### Basic Usage

```bash
# Initialize a new repository
computing-provider init

# Start the provider
computing-provider run

# Check provider status
computing-provider state

# List tasks
computing-provider task list --fcp
```

### With Custom Repository Path

```bash
# Use custom repository path
computing-provider --repo /custom/path init

# Or set environment variable
export CP_PATH=/custom/path
computing-provider init
```

### Getting Help

```bash
# General help
computing-provider --help

# Command-specific help
computing-provider task --help
computing-provider wallet --help

# Subcommand help
computing-provider task list --help
```

## Command Categories

### Setup Commands
- `init` - Initialize provider repository
- `config` - Configuration management

### Runtime Commands
- `run` - Start provider service
- `stop` - Stop provider service
- `restart` - Restart provider service

### Information Commands
- `info` - Provider information
- `state` - Provider state
- `status` - Service status

### Management Commands
- `account` - Account management
- `wallet` - Wallet operations
- `task` - Task management
- `collateral` - Collateral management

## Output Formats

The CLI supports different output formats:

### Table Format (Default)
```bash
computing-provider task list --fcp
```

### JSON Format
```bash
computing-provider task list --fcp --format json
```

### Verbose Output
```bash
computing-provider task list --fcp --verbose
```

## Error Handling

The CLI provides clear error messages and exit codes:

- `0`: Success
- `1`: General error
- `2`: Configuration error
- `3`: Network error
- `4`: Permission error

## Logging

Enable debug logging:

```bash
# Set log level
export CP_LOG_LEVEL=debug

# Or use flag
computing-provider --log-level debug run
```

## Configuration

The CLI reads configuration from:

1. Command line flags
2. Environment variables
3. Configuration file (`config.toml`)
4. Default values

## Security Considerations

- Private keys are stored securely in the repository
- Sensitive operations require confirmation
- Wallet operations are isolated

## Next Steps

- [Initialize your provider](init.md)
- [Configure your account](account.md)
- [Set up your wallet](wallet.md)
- [Start managing tasks](task.md) 