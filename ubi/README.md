# Edge Computing Provider(ECP)

**ECP (Edge Computing Provider)** specializes in processing data at the source of data generation, using minimal latency setups ideal for real-time applications. This provider handles specific, localized tasks directly on devices at the network’s edge, such as IoT devices. 

At the current stage, ECP supports the generation of **FIL-C2** and **Aleo Proof**. More ZK proof types will be gradually endorsed, such as Scroll, starkNet, etc

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
wget https://github.com/swanchain/go-computing-provider/releases/download/v1.0.2/computing-provider
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
                    --task-types 1,2,4
```
**Note:** `--task-types`: Supports 5 task types (1: Fil-C2, 2: Mining, 3: AI, 4: Inference, 5: NodePort), separated by commas. For ECP, it needs to be set to 1,2,4.
- Collateral `SWAN` for ECP
```bash
computing-provider collateral add --ecp --from <YOUR_WALLET_ADDRESS>  <AMOUNT>   
```
> If you want to withdraw `SWAN` from ECP
>```bash
>computing-provider collateral withdraw --ecp --owner <YOUR_WALLET_ADDRESS> --account <YOUR_CP_ACCOUNT> <amount>
>```

- Deposit `SwanETH` to Sequencer Account
    ```bash
    computing-provider sequencer add --from <YOUR_WALLET_ADDRESS>  <amount>
    ```
> If you want to Withdraw `SwanETH` from Sequencer Account
>
>    ```bash
>    computing-provider sequencer withdraw --owner <YOUR_OWNER_WALLET_ADDRESS>  <amount>
>    ```
    
>**Note:** the gas cost is decided by the **[Dynamic Pricing Strategy](https://docs.swanchain.io/bulders/market-provider/web3-zk-computing-market/sequencer)**


## Config resource price

**Pricing:** Indicating acceptance of smart pricing orders, which may include orders priced lower than self-determined pricing. default "true"

Configure it in the `$CP_PATH/price.toml`:
```
[API]
Pricing = "true"   
```

1. Generate the pricing config with default values(Located at `$CP_PATH/price.toml`):
```
computing-provider --repo <YOUR_CP_PATH> price generate
```

2. Customize your resource prices, adjust resource prices based on how many swans are configured per hour

```
vi $CP_PATH/price.toml
```
example:
```
TARGET_CPU="0.2"            # SWAN/thread-hour
TARGET_MEMORY="0.1"         # SWAN/GB-hour
TARGET_HD_EPHEMERAL="0.005" # SWAN/GB-hour
TARGET_GPU_DEFAULT="1.6"    # SWAN/Default GPU unit a hour
TARGET_GPU_3080=""          # SWAN/3080 GPU unit a hour
```

3. View the configured price information:
```
computing-provider --repo <YOUR_CP_PATH> price view
```

CP Hardware Price Info:                                  
TARGET_CPU:             0.2 SWAN/thread-hour            
TARGET_MEMORY:          0.1 SWAN/GB-hour                
TARGET_HD_EPHEMERAL:    0.005 SWAN/GB-hour              
TARGET_GPU_DEFAULT:     1.6 SWAN/Default GPU unit a hour
TARGET_GPU_3080:         SWAN/GPU unit a hour     

## Config and Receive Inference task
The task type is configured as 4 (inference), and need to be configured as follows:
  * For container services with a single port, use `traefik`. Using `traefik` as the entry point for requests, you need to configure a domain(*.example.com) to resolve to the IP where CP is running. The port 9000 must be open for external access.
  * For container services with multiple ports, use the public IP + port. Need to configure `PortRange`( one-to-one mapping between host ports and the public network IP).

Configure it in the `$CP_PATH/config.toml`:
```
[API]
Domain = ""                                 # The domain name
AutoDeleteImage = false                     # Default false, automatically delete unused images
PortRange = ["40000-40050","40060",""40065] # Externally exposed port number for deploying multi-port image tasks  
```

Check the Status of Inference and Mining task

- Use the following command:
```
computing-provider task list --ecp
```
- Example output:
```
TASK UUID                               TASK NAME                               IMAGE NAME                              CONTAINER NAME                                  CONTAINER STATUS        REWARD  CREATE TIME  
75f9df4e-b6a5-40b0-b7ac-02fb1840dafa    iron02                                  swanchain254/iron_mainnet_f2pool:latest iron02-2b0d5                                    terminated              0.0000  2024-10-24 10:23:32
842dd7d3-e9f0-4795-af3b-104fa5527099    Zil1                                    swanchain254/zil_mainnet_f2pool:latest  Zil1-gb5sq                                      terminated              0.6195  2024-11-15 03:49:44
30f60c6f-085e-4cd3-b751-bf6081a6a2e2    rvn-f2-003                              swanchain254/rvn_mainnet_f2pool:latest  rvn-f2-003-082ei                                terminated              7.5016  2024-11-15 05:40:14
9b16aa08-66aa-4ffe-b7d6-7a6db512ea6c    rvn-004                                 swanchain254/rvn_mainnet_f2pool:latest  rvn-004-to2cr                                   terminated              5.9034  2024-11-15 05:44:21
4c95d6e6-34b2-49de-be7d-7ee326b641c7    Zil                                     swanchain254/zil_mainnet_f2pool:latest  Zil-81ydq                                       terminated              0.1155  2024-11-15 06:56:49
```

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

The ECP can submit proofs to the Sequencer service, which will then package and submit all proofs from the entire network over a period of time (**currently 24 hours**) in a single transaction. This way, ECP only needs to pay a minimal gas fee to the Sequencer (currently, the gas is decided by the [Dynamic Pricing Strategy](https://docs.swanchain.io/bulders/market-provider/web3-zk-computing-market/sequencer)). For more detailed information, see [here](https://docs.swanchain.io/swan-provider/market-provider-mp/zk-engine/sequencer).

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




