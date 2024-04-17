# Edge Computing Provider(ECP)

**ECP (Edge Computing Provider)** specializes in processing data at the source of data generation, using minimal latency setups ideal for real-time applications. This provider handles specific, localized tasks directly on devices at the network’s edge, such as IoT devices. 

At the current stage, ECP supports the generation of **ZK-Snark proof of Filecoin network**, and more ZK proof types will be gradually supported, such as Aleo, Scroll, starkNet, etc

## Prerequisites
 - Running the `setup.sh`
```bash
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/setup.sh | bash
```

 - Download the parameters for `zk-fil`:
```bash
export PARENT_PATH="<YOUR_ZK-FIL_PARAMS_PATH>"

# 512MiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-512.sh | bash

# 32GiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-32.sh | bash

```
## Install ECP and Init CP Account
- Download `computing-provider`
```bash
wget https://github.com/swanchain/go-computing-provider/releases/download/v0.4.6/computing-provider
```

- Generate a new wallet address and deposit the `swan-eth`, refer [here](https://docs.swanchain.io/swan-testnet/atom-accelerator-race/before-you-get-started/bridge-tokens):
```bash
./computing-provider wallet new
```
Output: 
```
 0x9024a875f44172591373b92...31d67AcCEa
```
> **Note:** By default, the CP's repo is `~/.swan/computing`, you can configure it by `export CP_PATH="<YOUR_CP_PATH>"`

- Initialize ECP Account
```bash
./computing-provider init --ownerAddress <YOUR_WALLET_ADDRESS>
```

## Start ECP service
```bash
#!/bin/bash
export FIL_PROOFS_PARAMETER_CACHE=$PARENT_PATH
export RUST_GPU_TOOLS_CUSTOM_GPU="GeForce RTX 4090:16384"   
        
nohup ./computing-provider ubi daemon \
	--multi-address=/ip4/<YOUR_PUBLIC_IP>/tcp/<YOUR_PORT> \
	--node-name=<YOUR_NODE_NAME> >> cp.log 2>&1 &
```
**Note:**
 -  `<YOUR_ZK-FIL_PARAMS_PATH>` is your parameters directory, 
 - `RUST_GPU_TOOLS_CUSTOM_GPU` is your GPU model and cores
 - `<YOUR_PUBLIC_IP>`, `<YOUR_PORT>` are your public IP and port , 
 - `<YOUR_NODE_NAME>` is your CP name which will show in the dashboard.

