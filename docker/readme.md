Dockerfile
```
from golang:latest
WORKDIR /app
RUN mkdir -p /app/node
RUN git clone https://github.com/swanchain/go-computing-provider.git && cd go-computing-provider
RUN cd go-computing-provider && git checkout releases
RUN cd go-computing-provider &&  make
ENV CP_PATH=/app/node
RUN sed -i 's/sudo //g' /app/go-computing-provider/install.sh
RUN apt update
RUN apt install python3-full python3-pip -y 
ENV PIP_BREAK_SYSTEM_PACKAGES=1
RUN cd go-computing-provider &&  bash install.sh
ENV PATH=$PATH:/app/go-computing-provider
RUN apt-get clean
EXPOSE 8085


```

BUILD : 
`docker build -t swanchain .`

Result : 
```
[mohamed@l0c4lh0st node]$ dockebuild -t swanchain .
[+] Building 0.7s (15/15) FINISHED                                                                                                     docker:default
 => [internal] load .dockerignore                                                                                                                0.0s
 => => transferring context: 2B                                                                                                                  0.0s
 => [internal] load build definition from Dockerfile                                                                                             0.0s
 => => transferring dockerfile: 639B                                                                                                             0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                                                                 0.5s
 => [ 1/11] FROM docker.io/library/golang:latest@sha256:d5302d40dc5fbbf38ec472d1848a9d2391a13f93293a6a5b0b87c99dc0eaa6ae                         0.0s
 => CACHED [ 2/11] WORKDIR /app                                                                                                                  0.0s
 => CACHED [ 3/11] RUN mkdir -p /app/node                                                                                                        0.0s
 => CACHED [ 4/11] RUN git clone https://github.com/swanchain/go-computing-provider.git && cd go-computing-provider                              0.0s
 => CACHED [ 5/11] RUN cd go-computing-provider && git checkout releases                                                                         0.0s
 => CACHED [ 6/11] RUN cd go-computing-provider &&  make                                                                                         0.0s
 => CACHED [ 7/11] RUN sed -i 's/sudo //g' /app/go-computing-provider/install.sh                                                                 0.0s
 => CACHED [ 8/11] RUN apt update                                                                                                                0.0s
 => CACHED [ 9/11] RUN apt install python3-full python3-pip -y                                                                                   0.0s
 => CACHED [10/11] RUN cd go-computing-provider &&  bash install.sh                                                                              0.0s
 => CACHED [11/11] RUN apt-get clean                                                                                                             0.0s
 => exporting to image                                                                                                                           0.0s
 => => exporting layers                                                                                                                          0.0s
 => => writing image sha256:9ad4baf186988b8ff6b1b708e17d9cb9d687bcbd6717279d60d03630bab9e1a6                                                     0.0s
 => => naming to docker.io/library/swanchain  
```
usage : 

```
[mohamed@l0c4lh0st node]$ docker run --rm -it swanchain computing-provider
NAME:
   computing-provider - Swanchain decentralized computing network client

USAGE:
   computing-provider [global options] command [command options] [arguments...]

VERSION:
   0.4.5+git.57c9d53

COMMANDS:
   init        Initialize a new cp
   run         Start a cp process
   info        Print computing-provider info
   account     Manage account info of CP
   task        Manage tasks
   wallet      Manage wallets
   collateral  Manage the collateral amount
   ubi         Manage ubi tasks
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --repo value   repo directory for computing-provider client (default: "~/.swan/computing") [$CP_PATH]
   --help, -h     show help
   --version, -v  print the version

```
 mount volume for persistance :
docker run --rm -it -v ./node:/app/node swanchain computing-provider init --multi-address 127.0.0.1 --node-name dr4mohamed
`docker run --rm -it -v ./node:/app/node swanchain computing-provider wallet new`

edit config.toml and add cert and key and your public key and api MCS


`
docker run --rm -it -v ./docker-swan:/app/node -v ./letsencrypt:/etc/letsencrypt -v ./.kube:/root/.kube swanchain computing-provider run
`


/app/node   : persistance 
/etc/letsencrypt  :  cert domain
/root/.kube: k8s config file or service account


