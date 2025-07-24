# Installation Guide

This guide will walk you through the installation process for the Go Computing Provider.

## Prerequisites

Before installing the Computing Provider, ensure you have the following:

### System Requirements
- **Operating System**: Linux (Ubuntu 20.04+ recommended)
- **Go Version**: 1.21+ 
- **Public IP**: Required for network connectivity
- **Domain Name**: `*.example.com` format
- **SSL Certificate**: For secure communications

### Hardware Requirements
- **CPU**: Multi-core processor (4+ cores recommended)
- **RAM**: Minimum 8GB, 16GB+ recommended
- **Storage**: 100GB+ available space
- **Network**: Stable internet connection

## Installing Go

If you don't have Go installed, follow these steps:

```bash
# Download and install Go 1.21+
wget -c https://golang.org/dl/go1.21.7.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

# Add Go to your PATH
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc

# Verify installation
go version
```

## Building from Source

### 1. Clone the Repository

```bash
git clone https://github.com/lagrangedao/go-computing-provider.git
cd go-computing-provider
```

### 2. Build the Binary

```bash
# Build the computing-provider binary
make build

# Or build manually
go build -o computing-provider cmd/computing-provider/main.go
```

### 3. Install System-wide (Optional)

```bash
# Copy to system PATH
sudo cp computing-provider /usr/local/bin/

# Verify installation
computing-provider --version
```

## Using the Install Script

The project includes an automated installation script:

```bash
# Make the script executable
chmod +x install.sh

# Run the installation
./install.sh
```

## Docker Installation (Alternative)

If you prefer using Docker:

```bash
# Build the Docker image
docker build -t computing-provider .

# Run the container
docker run -it --rm computing-provider --help
```

## Verification

After installation, verify that everything is working:

```bash
# Check version
computing-provider --version

# Check help
computing-provider --help

# Initialize a new repository
computing-provider init
```

## Next Steps

After successful installation:

1. [Configure your environment](configuration.md)
2. [Set up your wallet](wallet.md)
3. [Choose your provider type](getting-started.md)

## Troubleshooting

### Common Issues

**Go not found in PATH**
```bash
export PATH=$PATH:/usr/local/go/bin
```

**Permission denied errors**
```bash
sudo chmod +x computing-provider
```

**Build failures**
- Ensure Go version is 1.21+
- Check that all dependencies are installed
- Verify you have sufficient disk space

For more troubleshooting help, see the [Troubleshooting Guide](troubleshooting.md). 