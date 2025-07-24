# Wallet Management

The `wallet` command allows you to manage wallet operations for your computing provider.

## Overview

```bash
computing-provider wallet <subcommand> [flags]
```

## Wallet Address Types

The Computing Provider uses three different wallet addresses for security:

1. **Owner Address**: Controls account settings and permissions
2. **Worker Address**: Used for submitting proofs and paying gas fees
3. **Beneficiary Address**: Receives all earnings

## Subcommands

### Initialize Wallet

Create a new wallet for your computing provider.

```bash
computing-provider wallet init [flags]
```

#### Flags

- `--password`: Set wallet password
- `--keystore`: Specify keystore directory

#### Examples

```bash
# Initialize wallet with default settings
computing-provider wallet init

# Initialize wallet with custom password
computing-provider wallet init --password mypassword

# Initialize wallet with custom keystore location
computing-provider wallet init --keystore /custom/keystore
```

### Import Wallet

Import an existing wallet using a private key.

```bash
computing-provider wallet import [flags]
```

#### Flags

- `--private-key`: Private key to import
- `--password`: Wallet password
- `--keystore`: Specify keystore directory

#### Examples

```bash
# Import wallet with private key
computing-provider wallet import --private-key 0x123...

# Import with custom password
computing-provider wallet import --private-key 0x123... --password mypassword
```

### List Wallets

List all available wallets.

```bash
computing-provider wallet list [flags]
```

#### Flags

- `--verbose, -v`: Show detailed information

#### Examples

```bash
# List all wallets
computing-provider wallet list

# List with detailed information
computing-provider wallet list --verbose
```

### Show Wallet

Display information about a specific wallet.

```bash
computing-provider wallet show [address] [flags]
```

#### Arguments

- `address`: Wallet address to display

#### Examples

```bash
# Show wallet information
computing-provider wallet show 0x123...

# Show wallet with balance
computing-provider wallet show 0x123... --balance
```

### Export Wallet

Export wallet information.

```bash
computing-provider wallet export [address] [flags]
```

#### Arguments

- `address`: Wallet address to export

#### Flags

- `--private-key`: Export private key
- `--keystore`: Export keystore file

#### Examples

```bash
# Export wallet keystore
computing-provider wallet export 0x123... --keystore

# Export private key (use with caution)
computing-provider wallet export 0x123... --private-key
```

### Set Addresses

Set the different address types for your provider.

#### Set Owner Address

```bash
computing-provider account set-owner --address <OWNER_ADDRESS>
```

#### Set Worker Address

```bash
computing-provider account set-worker --address <WORKER_ADDRESS>
```

#### Set Beneficiary Address

```bash
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

#### Examples

```bash
# Set owner address
computing-provider account set-owner --address 0x123...

# Set worker address
computing-provider account set-worker --address 0x456...

# Set beneficiary address
computing-provider account set-beneficiary --address 0x789...
```

### Check Balance

Check wallet balance.

```bash
computing-provider wallet balance [address] [flags]
```

#### Arguments

- `address`: Wallet address to check

#### Flags

- `--network`: Specify network (mainnet, testnet)

#### Examples

```bash
# Check balance
computing-provider wallet balance 0x123...

# Check balance on specific network
computing-provider wallet balance 0x123... --network testnet
```

## Wallet Security

### Best Practices

1. **Separate Addresses**: Use different addresses for different purposes
2. **Secure Storage**: Keep private keys secure and encrypted
3. **Backup**: Regularly backup your keystore files
4. **Access Control**: Limit access to wallet files

### Security Commands

```bash
# Change wallet password
computing-provider wallet change-password

# Lock wallet
computing-provider wallet lock

# Unlock wallet
computing-provider wallet unlock --password mypassword
```

## Keystore Management

### Keystore Location

Default keystore location: `~/.swan/computing/keystore`

### Keystore Operations

```bash
# List keystore files
computing-provider wallet keystore list

# Backup keystore
computing-provider wallet keystore backup --destination /backup/

# Restore keystore
computing-provider wallet keystore restore --source /backup/
```

## Transaction Management

### Sign Transaction

```bash
computing-provider wallet sign --tx <transaction_data>
```

### Send Transaction

```bash
computing-provider wallet send --to <recipient> --amount <amount>
```

### Examples

```bash
# Sign a transaction
computing-provider wallet sign --tx 0x123...

# Send ETH
computing-provider wallet send --to 0x456... --amount 0.1

# Send with custom gas
computing-provider wallet send --to 0x456... --amount 0.1 --gas-price 20000000000
```

## Network Support

### Supported Networks

- **Mainnet**: Ethereum mainnet
- **Testnet**: Sepolia testnet
- **Local**: Local development network

### Network-Specific Commands

```bash
# Check balance on mainnet
computing-provider wallet balance 0x123... --network mainnet

# Check balance on testnet
computing-provider wallet balance 0x123... --network testnet
```

## Troubleshooting

### Common Issues

1. **Wallet Not Found**
   - Check keystore directory
   - Verify wallet address
   - Check file permissions

2. **Incorrect Password**
   - Verify password
   - Check for typos
   - Reset password if needed

3. **Network Issues**
   - Check network connectivity
   - Verify RPC endpoint
   - Check network configuration

### Debug Commands

```bash
# Check wallet status
computing-provider wallet status

# Verify keystore integrity
computing-provider wallet verify

# Check network connectivity
computing-provider network status
```

## Related Commands

- [`account`](account.md) - Account management
- [`collateral`](collateral.md) - Collateral management
- [`network`](network.md) - Network configuration 