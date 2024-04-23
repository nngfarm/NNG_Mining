package mining

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"math/big"
	"nng_ming/pkg/nng_token"
	"nng_ming/utils"
	"regexp"
)

func Run() {
	var (
		err         error
		privateKey  *ecdsa.PrivateKey
		nngContract = common.HexToAddress(NNG_Contract)
	)
	// Get Current account
	for {
		privateKey, err = AccountFromKeystore()
		// nil err
		if err == nil {
			break
		}
		fmt.Println("Failed to read private key")
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	// print address
	fmt.Println("Your are start with address:", address)
	// rpc
	rpc := flag.String("rpc", Default_RPC, "Request the endpoint of the mining chain")
	// new NNGMiner
	ethCli, err := ethclient.Dial(*rpc)
	if err != nil {
		fmt.Printf("failed to connect to ethclient: %v", err)
		return
	}
	// Check if pledged
	nng, err := nng_token.NewNng(nngContract, ethCli)
	if err != nil {
		fmt.Printf("failed to create nng: %v", err)
		return
	}
	netWorkId, err := ethCli.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Failed to get network ID: %v\n", err)
		return
	}
	// new keyed transactor
	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, netWorkId)
	if err != nil {
		fmt.Printf("Failed to get transact Opts: %s", err.Error())
		return
	}
	// bigZero
	big100 := utils.AmountToWei(100, 18)
	// read the flag from console
	run := flag.String("run", "mining", "Mining NNG. Required: [mining|pledge|claim]")
	// Get Recipient
	recipient := flag.String("recipient", "", "The address to receive NNG")
	// thread
	thread := flag.Int("thread", 8, "Number of threads to mine nng")
	// parse flag params
	flag.Parse()
	//  switch script to run
	switch *run {
	case "mining":
		// start with thread
		fmt.Println("Thread", *thread, ", Mining ...")
		// Check Pledge
		pledge, err := nng.UserPledges(nil, address)
		if err != nil {
			fmt.Println("Failed to get pledges, Please try again later.")
			return
		}
		// No pledge
		if pledge.IsPledged && pledge.Amount.Cmp(big100) >= 0 {
			// start Mining
			var inputData []byte
			// get code
			methodID, _ := hex.DecodeString("0c6008af")
			// encodeParams
			paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
			// inputData
			inputData = append(inputData, methodID...)
			inputData = append(inputData, paddedAddress...)
			// Get address bytecode
			resData, err := ethCli.CallContract(context.Background(), ethereum.CallMsg{To: &NNG_Factory, Data: inputData}, nil)
			if err != nil || len(resData) < 64 {
				fmt.Println("Failed to get bytecode")
				return
			}
			// codeLength
			codeLength := new(big.Int).SetBytes(resData[32:64]).Uint64()
			// required > 64+codeLength
			if len(resData) < 64 + int(codeLength) {
				fmt.Println("Failed to get bytecode")
				return
			}
			// run
			miner := NewNngMiner(nng, ethCli, transOpts, *thread)
			// keccak256(bytecode)
			miner.Bytecode = crypto.Keccak256Hash(resData[64:64+codeLength]).Bytes()
			// Recipient
			miner.Recipient = address
			// has recipient address
			if regexp.MustCompile("^0x[0-9a-fA-F]{40}$").MatchString(*recipient) {
				miner.Recipient = common.HexToAddress(*recipient)
			}
			// run miner
			miner.Run()
			return
		}
		isToPledge := "n"
		// pledge amount < 100
		fmt.Println("Your pledged NNG: ", utils.WeiToAmount(pledge.Amount, 18))
		fmt.Println("Please pledge at lease [100 NNG] to start mining.")
		fmt.Printf("Whether to go to pledge: [y/n] ")
		fmt.Scan(&isToPledge)
		// to pledge
		if isToPledge != "y" && isToPledge != "Y" {
			return
		}
		fallthrough
	case "pledge":
		var (
			period          uint8
			amount          string
			minPledgeAmount = decimal.NewFromInt(100)
			maxPledgeAmount = decimal.NewFromInt(1000)
		)
		// Check Balance
		balance, err := ethCli.BalanceAt(context.Background(), address, nil)
		if err != nil {
			fmt.Println("Failed to get Ether balance of address, Please try again later.")
			return
		}
		// No ether found
		if balance.Cmp(new(big.Int)) == 0 {
			fmt.Println("Insufficient balance of this address")
			return
		}
		// Check if pledged
		pledge, err := nng.UserPledges(nil, address)
		if err != nil {
			fmt.Println("Failed to get pledges, Please try again later.")
			return
		}
		// pledged
		if pledge.IsPledged && pledge.Amount.Cmp(big100) >= 0 {
			fmt.Println("You have already pledged")
			fmt.Println("Pledge Amount: ", utils.WeiToAmount(pledge.Amount, 18), "NNG")
			fmt.Println("Pledge Unlock Time:", pledge.UnlockTime)
			return
		}
		// pledge
		for {
			fmt.Println("NNG Token Address: ", NNG_Contract)
			fmt.Printf("Pledge Input Your Pledge Amount (Range[100, 1000]): ")
			fmt.Scan(&amount)
			//
			pledgeAmount, err := decimal.NewFromString(amount)
			// has err
			if err != nil {
				fmt.Printf("Failed to parse pledge amount: %v \n", err)
				continue
			}
			// only [100, 1000]
			if pledgeAmount.GreaterThan(maxPledgeAmount) || pledgeAmount.LessThan(minPledgeAmount) {
				fmt.Println("Pledge amount must be range [100, 1000] NNG")
				continue
			}
			// Pledge Period
			fmt.Println("0 ==> 1 Month")
			fmt.Println("1 ==> 3 Month")
			fmt.Println("2 ==> 6 Month")
			fmt.Println("3 ==> 1 Year")
			fmt.Printf("Pledge Choose Your Pledge Period: ")
			fmt.Scan(&period)
			// require period < 3
			if period > 3 {
				fmt.Printf("Pledge period should be [0,1,2,3] \n")
				continue
			}
			break
		}
		// pledge amount
		fmt.Println("Pledge Amount:", amount, "NNG")
		// Pledge
		tx, err := nng.Pledge(transOpts, utils.AmountToWei(amount, 18), period)
		// has err
		if err != nil {
			fmt.Printf("Failed to pledge transaction: %v \n", err)
			return
		}
		// Pledge pending to chain
		fmt.Printf("Pledge transaction successfully submitted: %s \n", tx.Hash())

	case "claim":
		// if user want claim their pledged nng.
		// Check if pledged
		pledge, err := nng.UserPledges(nil, address)
		if err != nil {
			fmt.Println("Failed to get pledges, Please try again later.")
			return
		}
		// pledged
		if !pledge.IsPledged && pledge.Amount.Cmp(new(big.Int)) == 0 {
			fmt.Println("You haven’t pledged yet")
			return
		}

		tx, err := nng.ClaimPledge(transOpts)
		// err
		if err != nil {
			fmt.Println("Failed to claim pledge: ", err.Error())
			return
		}
		fmt.Println("Claim transaction successfully submitted: ", tx.Hash())

	default:
		panic("Only [mining|stake|unStake] are supported")
	}
}
