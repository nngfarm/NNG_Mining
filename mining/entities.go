package mining

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"nng_ming/pkg/nng_token"
)

const (
	NNG_Contract = "0x4Dd448987835a5d6067aa0D3e3D0629B7bCDA174"
	// Default_RPC  = "https://bsc-dataseed1.ninicoin.io"
	Default_RPC  = "https://bsc-testnet.public.blastapi.io"
)

var (
	NNG_Factory = common.HexToAddress("0xB52797449cce9b39a068345c524e3716d81Ed630")
)

// NngMiner dispatch job to worker
type NngMiner struct {
	EthCli         *ethclient.Client
	NNG            *nng_token.Nng
	TransOpts      *bind.TransactOpts
	Recipient      common.Address
	Bytecode       []byte
	Nonces         map[string]NngReward
	Rewards        map[string]common.Hash
	LastSubmitTime int64
	// A pool of workers channels that are registered with the dispatcher
	MaxWorkers int
}

type NngReward struct {
	Nonce   [32]byte
	Address common.Address
}
