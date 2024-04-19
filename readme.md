# NNG Mining Farm

## Project Overview

nng farm is a POW-verified mining farm, which is a project that is beneficial to the ecological development of the ETH chain.
Anyone can gain benefits by mining NNG, but the benefits are not only the mining itself. 
But more Easter egg benefits are waiting to be mined.

## Features

- **Smart Contract**: Written in Solidity and deployed on the ETH/BSC/Base... network, supporting secure mining and distribution of tokens.
- **Backend Service**: Developed in Go, responsible for handling user requests, interacting with the smart contract, and managing mining operations.

## Technology Stack

- **Solidity**: Language for smart contract development.
- **Go**: Language for backend service development.

## Quick Start

### Prerequisites

- Go 1.18 or higher

### Setup Steps

1. **Clone the Repository**

```bash
git clone https://github.com/nngfarm/NNG_Mining.git
cd NNG_Mining
```

#### Build nng_mining

```shell
# for linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build/NNG_Mining ./main.go
# chmod -v u+x NNG_Mining
```

```shell
# for mac
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./build/NNG_Mining ./main.go
# chmod -v u+x NNG_Mining
```        

```shell
# for windows
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./build/NNG_Mining.exe ./main.go
```

#### Start NNG_Mining

```
NNG_Mining -rpc YOUR_CUSTOM_RPC -run [mining|pledge|claim] -thread DEFAULT_8 -recipient YOUR_REWARD_RECIPIENT_ADDRESS

# flags
-rpc 
    Request the rpc node of the blockchain, which can be customized.
    If not set, the default rpc node will be used.
-run 
    mining: By default, it runs as mining. This flag will start your NNG mining journey.
    pledge: You can stake your tokens via -run pledge
    claim: You can withdraw your staked tokens by -run claim
-thread
    Defalt is 8. You can use this flag to define the number of coroutines enabled for mining
-recipient
    You can customize the address where you want to receive mining rewards.
    If not set, the default is the mining address.      
```


