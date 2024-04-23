package mining

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"nng_ming/pkg/nng_token"
)

const (
	NNG_Contract = "0xCa3951797567AdeFfb25410C2b6c17e84b71Af0b"
	Default_RPC  = "https://bsc-dataseed1.ninicoin.io"
	// Default_RPC  = "https://bsc-testnet.public.blastapi.io"
)

var (
	NNG_Factory = common.HexToAddress("0x05440e395abf9bAC34259e6616999a2529C431bb")
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
