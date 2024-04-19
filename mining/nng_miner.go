package mining

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"nng_ming/pkg/nng_token"
	"strings"
	"sync"
	"time"
)

func NewNngMiner(nng *nng_token.Nng, ethCli *ethclient.Client, tranOpts *bind.TransactOpts, workers int) NngMiner {
	return NngMiner{
		NNG:        nng,
		EthCli:     ethCli,
		TransOpts:  tranOpts,
		MaxWorkers: workers,
		Nonces:     make(map[string]NngReward),
		Rewards:    make(map[string]common.Hash),
	}
}

// Run mining mining mining
func (d *NngMiner) Run() {
	// nonces
	var (
		wg      sync.WaitGroup
		rewards = make(chan NngReward, d.MaxWorkers)
	)
	// wait
	defer wg.Wait()
	// start nng worker
	for i := 0; i < d.MaxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// mining
			mining(d.Bytecode, rewards)
		}()
	}
	// create a submit ticker
	go d.submitRewardTicker()
	// receive mining rewards
	for reward := range rewards {
		d.submitReward(reward)
	}
}

// startMining
func mining(bytecode []byte, reward chan<- NngReward) {
	for {
		var salt [32]byte

		rand.Read(salt[:])

		address := crypto.CreateAddress2(NNG_Factory, salt, bytecode)
		addrHex := address.Hex()

		if strings.HasPrefix(addrHex, "0x000000") {

			reward <- NngReward{
				Address: address,
				Nonce:   salt,
			}

			continue
		}
	}
}

func (d *NngMiner) submitReward(reward NngReward) {
	// rewardAddress
	rewardAddress := reward.Address.Hex()
	// set to cache list
	d.Nonces[rewardAddress] = reward
	// After some time, try submitting again
	if d.LastSubmitTime+3000 > time.Now().Unix() {
		return
	}
	// submit mining reward
	tx, err := d.NNG.Mine(d.TransOpts, reward.Nonce, d.Recipient)
	if err != nil {
		fmt.Printf("Mining error: %v \n", err)
		return
	}
	// cache rewards
	d.Rewards[rewardAddress] = tx.Hash()
	// To cache list
	fmt.Println("Found Reward.address:", rewardAddress)
	fmt.Println("Found Reward.nonce:", "0x"+hex.EncodeToString(reward.Nonce[:]))
	// TxHash
	fmt.Println("Submitting reward. TxHash: ", tx.Hash())
}

func (d *NngMiner) submitRewardTicker() {
	// create a ticker for submit reward
	ticker := time.NewTicker(1 * time.Minute)
	// check nonce
	for range ticker.C {
		// check Tx Hash
		if len(d.Rewards) > 0 {
			// foreach all rewards
			for addr, hash := range d.Rewards {
				// Check txHash
				receipt, err := d.EthCli.TransactionReceipt(context.Background(), hash)
				if err != nil {
					continue
				}
				// remove rewards cache
				delete(d.Rewards, addr)
				// Mining success
				if receipt.Status == 1 {
					delete(d.Nonces, addr)
					d.LastSubmitTime = time.Now().Unix()
				}
			}
		}
		// has Nonce
		if len(d.Nonces) > 0 {
			// Get rand nonce to submit
			for _, v := range d.Nonces {
				d.submitReward(v)
				break
			}
		}
	}
}
