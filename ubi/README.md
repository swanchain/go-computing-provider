# Edge Computing Provider(ECP)

**ECP (Edge Computing Provider)** specializes in processing data at the source of data generation, using minimal latency setups ideal for real-time applications. This provider handles specific, localized tasks directly on devices at the network’s edge, such as IoT devices. 

At the current stage, ECP supports the generation of **ZK-Snark proof of Filecoin network**, and more ZK proof types will be gradually endorsed, such as Aleo, Scroll, starkNet, etc

## Prerequisites
 - Need to map the ECP service port of the intranet to the public network, the default port is`9085`:
```
 <Intranet_IP>:<9085> <--> <Public_IP>:<PORT>
```
 - Running the `setup.sh`
```bash
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/setup.sh | bash
```

 - Download the v28 parameters for `ZK-FIL` task:
```bash
# At least 200G storage is needed
export PARENT_PATH="<V28_PARAMS_PATH>"

# 512MiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-512.sh | bash

# 32GiB parameters
curl -fsSL https://raw.githubusercontent.com/swanchain/go-computing-provider/releases/ubi/fetch-param-32.sh | bash

```
## Install ECP and Init CP Account
- Download `computing-provider`
```bash
wget https://github.com/swanchain/go-computing-provider/releases/download/v0.6.5/computing-provider
```

- Initialize ECP repo
```bash
 ./computing-provider init --multi-address=/ip4/<YOUR_PUBLIC_IP>/tcp/<YOUR_PORT> --node-name=<YOUR_NODE_NAME>
```
- Generate a new wallet address and deposit the `SwanETH`, refer [here](https://docs.swanchain.io/swan-mainnet/getting-started-guide):
```bash
./computing-provider wallet new
```
Output: 
```
 0x9024a875f44172591373b92...31d67AcCEa
```

 - **[OPTIONAL]** You can also import your own wallet by private key

```
./computing-provider wallet import private.key
```

>**Note:** 
>1. By default, the CP's repo is `~/.swan/computing`, you can configure it by `export CP_PATH="<YOUR_CP_PATH>"`
>
>2. `private.key` is a file that contains the private key



- Initialize ECP Account
```bash
./computing-provider account create \
                    --ownerAddress <YOUR_OWNER_ADDRESS> \
                    --workerAddress <YOUR_WORKER_ADDRESS> \
                    --beneficiaryAddress <YOUR_BENEFICIAERY_ADDRESS>  \
                    --task-types 1,4
```
**Note:** `--task-types`: Supports 4 task types (1: Fil-C2-512M, 2: Mining, 3: AI, 4: Fil-C2-32G), separated by commas. For ECP, it needs to be set to 1,2,4.
- Collateral `SWANC` for ECP
```bash
computing-provider collateral add --ecp --from <YOUR_WALLET_ADDRESS>  <AMOUNT>   
```
> If you want to withdraw `SWANC` from ECP
>```bash
>computing-provider collateral withdraw --ecp --owner <YOUR_WALLET_ADDRESS> --account <YOUR_CP_ACCOUNT> <amount>
>```

>**Note:** A minimum of 100 `SWANC` collateral is currently required to receive a ZK task.

- Deposit `SwanETH` to Sequencer Account
    ```bash
    computing-provider sequencer add --from <YOUR_WALLET_ADDRESS>  <amount>
    ```
> If you want to Withdraw `SwanETH` from Sequencer Account
>
>    ```bash
>    computing-provider sequencer withdraw --owner <YOUR_OWNER_WALLET_ADDRESS>  <amount>
>    ```
    
>**Note:** Currently one zk-task requires 0.00001 SwanETH, 

## Start ECP service
```bash
#!/bin/bash
export FIL_PROOFS_PARAMETER_CACHE=$PARENT_PATH
export RUST_GPU_TOOLS_CUSTOM_GPU="GeForce RTX 4090:16384"
        
nohup ./computing-provider ubi daemon >> cp.log 2>&1 &
```
**Note:**
-  `<FIL_PROOFS_PARAMETER_CACHE>` is your parameters directory,
- `RUST_GPU_TOOLS_CUSTOM_GPU` is your GPU model and cores, you should update it to your own GPU model. More examples can be found [here](https://github.com/filecoin-project/bellperson?tab=readme-ov-file#supported--tested-cards)
- `<YOUR_PUBLIC_IP>`, `<YOUR_PORT>` are your public IP and port ,
- `<YOUR_NODE_NAME>` is your CP name which will show in the dashboard, If not specified, the default is `hostname`.

## About Sequencer

### Why need Sequencer?
In past tests, we discovered that due to the frequent interactions required by ECP (Ethereum Compliance Proof) to submit proofs to the blockchain, ECP incurs significant gas costs. To reduce these gas costs, the Sequencer has emerged as a Layer 3 solution.

The ECP can submit proofs to the Sequencer service, which will then package and submit all proofs from the entire network over a period of time (**currently 24 hours**) in a single transaction. This way, ECP only needs to pay a minimal gas fee to the Sequencer (currently, **a single proof requires 0.00001 SwanETH**). For more detailed information, see [here](https://docs.swanchain.io/swan-provider/market-provider-mp/zk-engine/sequencer).

### How to Set it?
We **strongly recommend** enabling the Sequencer feature (enabled by default). The steps to enable it are as follows:

 - Modify the `config.toml` file.
```
[UBI]
EnableSequencer = true             # Submit the proof to Sequencer service(default: true)
AutoChainProof = false             # when sequencer doesn't have enough funds or the service is unavailable, automatically submit proof to the Swan chain 

```
 - Deposit SwanETH into the Sequencer account.
```
computing-provider sequencer add --from <YOUR_WALLET_ADDRESS>  <amount>
```

 - Restart the ECP

```
#!/bin/bash
export FIL_PROOFS_PARAMETER_CACHE=$PARENT_PATH
export RUST_GPU_TOOLS_CUSTOM_GPU="GeForce RTX 4090:16384"
        
nohup ./computing-provider ubi daemon >> cp.log 2>&1 &
```




