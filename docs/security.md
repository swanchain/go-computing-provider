# Security Best Practices

This guide outlines security best practices for running the Go Computing Provider securely.

## Overview

Security is crucial when running a computing provider, as you'll be handling financial transactions, private keys, and potentially sensitive computing tasks.

## Wallet Security

### Address Separation

The Computing Provider uses three different wallet addresses for security:

1. **Owner Address**: Controls account settings and permissions
2. **Worker Address**: Used for submitting proofs and paying gas fees
3. **Beneficiary Address**: Receives all earnings

### Best Practices

**Keep Private Keys Secure**
```bash
# Store private keys in secure location
chmod 600 ~/.swan/computing/keystore/*

# Use strong passwords for wallet encryption
computing-provider wallet init --password "strong-password-here"
```

**Separate Addresses**
```bash
# Use different addresses for different purposes
computing-provider account set-owner --address <OWNER_ADDRESS>
computing-provider account set-worker --address <WORKER_ADDRESS>
computing-provider account set-beneficiary --address <BENEFICIARY_ADDRESS>
```

**Regular Backups**
```bash
# Backup keystore files
cp -r ~/.swan/computing/keystore ~/.swan/computing/keystore.backup

# Store backups in secure location
# Consider encrypted storage for backups
```

### Security Commands

```bash
# Lock wallet when not in use
computing-provider wallet lock

# Unlock wallet when needed
computing-provider wallet unlock --password <PASSWORD>

# Change wallet password
computing-provider wallet change-password
```

## System Security

### Operating System Security

**Keep System Updated**
```bash
# Regular system updates
sudo apt update && sudo apt upgrade

# Security updates
sudo apt update && sudo apt upgrade --security
```

**User Management**
```bash
# Create dedicated user for computing provider
sudo adduser computing-provider

# Add user to necessary groups
sudo usermod -aG docker computing-provider
sudo usermod -aG sudo computing-provider

# Use dedicated user for running provider
sudo -u computing-provider computing-provider run
```

**File Permissions**
```bash
# Secure repository directory
chmod 700 ~/.swan/computing
chmod 600 ~/.swan/computing/config.toml
chmod 600 ~/.swan/computing/keystore/*

# Secure log files
chmod 600 cp.log
```

### Network Security

**Firewall Configuration**
```bash
# Enable firewall
sudo ufw enable

# Allow necessary ports
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 80/tcp    # HTTP
sudo ufw allow 443/tcp   # HTTPS
sudo ufw allow 6443/tcp  # Kubernetes API (FCP)

# Deny unnecessary ports
sudo ufw deny 8080/tcp   # Default provider port (use reverse proxy)
```

**SSL/TLS Configuration**
```bash
# Use SSL certificates for provider
# Configure in config.toml

[server]
ssl_enabled = true
ssl_cert = "/path/to/cert.pem"
ssl_key = "/path/to/key.pem"
```

**Reverse Proxy Setup**
```bash
# Use nginx as reverse proxy
sudo apt install nginx

# Configure nginx to proxy to provider
# This hides the provider behind a secure proxy
```

### Kubernetes Security (FCP)

**RBAC Configuration**
```bash
# Create service account
kubectl create serviceaccount computing-provider

# Create role
kubectl create role computing-provider \
  --verb=get,list,watch,create,update,patch,delete \
  --resource=pods,services,ingresses

# Bind role to service account
kubectl create rolebinding computing-provider \
  --role=computing-provider \
  --serviceaccount=default:computing-provider
```

**Network Policies**
```bash
# Apply network policies
kubectl apply -f network-policy.yaml

# Restrict pod-to-pod communication
# Only allow necessary network access
```

**Secrets Management**
```bash
# Store sensitive data in Kubernetes secrets
kubectl create secret generic provider-secrets \
  --from-literal=api-key=<API_KEY> \
  --from-literal=private-key=<PRIVATE_KEY>

# Use secrets in deployments
# Never store secrets in plain text
```

## Application Security

### Configuration Security

**Secure Configuration**
```toml
# config.toml
[server]
host = "127.0.0.1"  # Bind to localhost only
port = 8080
ssl_enabled = true

[auth]
enabled = true
jwt_secret = "your-secure-jwt-secret"
session_timeout = 3600
```

**Environment Variables**
```bash
# Use environment variables for sensitive data
export CP_WALLET_PASSWORD="secure-password"
export CP_API_KEY="your-api-key"

# Don't hardcode secrets in configuration files
```

### Access Control

**API Authentication**
```bash
# Enable authentication for provider API
# Use JWT tokens for API access
# Implement rate limiting
```

**Logging and Monitoring**
```bash
# Enable security logging
# Monitor for suspicious activity
# Log authentication attempts
# Log financial transactions
```

## Container Security (FCP)

### Image Security

**Use Trusted Images**
```bash
# Use official images from trusted sources
# Scan images for vulnerabilities
# Keep images updated

# Example: Use official NVIDIA CUDA images
FROM nvidia/cuda:11.8-base-ubuntu20.04
```

**Image Scanning**
```bash
# Scan images for vulnerabilities
docker scan <image-name>

# Use tools like Trivy or Clair
# Integrate scanning into CI/CD pipeline
```

### Runtime Security

**Container Security**
```bash
# Run containers as non-root user
# Use read-only root filesystem
# Limit container capabilities
# Use security profiles
```

**Resource Limits**
```yaml
# Kubernetes resource limits
resources:
  limits:
    cpu: "4"
    memory: "8Gi"
  requests:
    cpu: "2"
    memory: "4Gi"
```

## Monitoring and Auditing

### Security Monitoring

**System Monitoring**
```bash
# Monitor system logs
journalctl -f

# Monitor authentication logs
tail -f /var/log/auth.log

# Monitor network connections
netstat -tuln
```

**Provider Monitoring**
```bash
# Monitor provider logs
tail -f cp.log

# Monitor for suspicious activity
grep -i "error\|warning\|failed" cp.log

# Monitor financial transactions
computing-provider collateral info --ecp
```

### Audit Logging

**Enable Audit Logs**
```bash
# Enable Kubernetes audit logging
# Log all API requests
# Monitor for unauthorized access
```

**Regular Audits**
```bash
# Regular security audits
# Review access logs
# Check for unusual activity
# Verify configuration integrity
```

## Incident Response

### Security Incidents

**Detect Incidents**
```bash
# Monitor for unusual activity
# Check for unauthorized access
# Monitor for failed authentication
# Check for unusual network traffic
```

**Respond to Incidents**
```bash
# Immediate response
# Isolate affected systems
# Preserve evidence
# Notify relevant parties
```

**Recovery Procedures**
```bash
# Restore from secure backups
# Change compromised credentials
# Update security measures
# Document incident
```

## Compliance

### Regulatory Compliance

**Data Protection**
```bash
# Follow data protection regulations
# Encrypt sensitive data
# Implement access controls
# Regular security assessments
```

**Financial Compliance**
```bash
# Maintain accurate financial records
# Regular financial audits
# Transparent reporting
# Compliance with tax regulations
```

## Security Checklist

### Pre-deployment

- [ ] System updated with latest security patches
- [ ] Firewall configured and enabled
- [ ] SSL certificates installed
- [ ] Strong passwords set
- [ ] Wallet addresses separated
- [ ] File permissions secured
- [ ] User accounts properly configured

### Runtime

- [ ] Regular security updates applied
- [ ] Logs monitored for suspicious activity
- [ ] Backups performed regularly
- [ ] Access controls enforced
- [ ] Network security maintained
- [ ] Financial transactions monitored

### Maintenance

- [ ] Regular security audits performed
- [ ] Incident response plan tested
- [ ] Security policies updated
- [ ] Staff security training completed
- [ ] Compliance requirements met

## Security Tools

### Recommended Tools

**System Security**
- `fail2ban`: Intrusion prevention
- `rkhunter`: Rootkit detection
- `lynis`: Security auditing
- `clamav`: Antivirus scanning

**Network Security**
- `wireshark`: Network analysis
- `nmap`: Network scanning
- `tcpdump`: Packet capture
- `iptables`: Firewall management

**Container Security**
- `trivy`: Vulnerability scanning
- `falco`: Runtime security monitoring
- `opa`: Policy enforcement
- `kyverno`: Kubernetes policy engine

## Resources

### Security Documentation

- [Swan Chain Security Guidelines](https://docs.swanchain.io/security)
- [Kubernetes Security Best Practices](https://kubernetes.io/docs/concepts/security/)
- [Docker Security Best Practices](https://docs.docker.com/develop/security/)

### Security Communities

- [Swan Chain Discord](https://discord.gg/Jd2BFSVCKw)
- [Kubernetes Security SIG](https://github.com/kubernetes/community/tree/master/sig-security)
- [Docker Security](https://www.docker.com/security)

### Security Tools

- [OWASP Security Tools](https://owasp.org/www-community/Source_Code_Analysis_Tools)
- [NIST Cybersecurity Framework](https://www.nist.gov/cyberframework)
- [CIS Benchmarks](https://www.cisecurity.org/benchmarks/) 