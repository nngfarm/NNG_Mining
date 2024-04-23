// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nng_token

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// NebulaNuggetPledgeInfo is an auto generated low-level Go binding around an user-defined struct.
type NebulaNuggetPledgeInfo struct {
	Amount     *big.Int
	UnlockTime *big.Int
	Period     uint8
	IsPledged  bool
}

// NngMetaData contains all meta data concerning the Nng contract.
var NngMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"AmountRewardFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"HalvingInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_HALVINGS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MaxPledgeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinPledgeAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PeriodRewardFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UserMiningLimitOfBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimPledge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"difficulty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPledgeMiningReward\",\"inputs\":[{\"name\":\"s\",\"type\":\"tuple\",\"internalType\":\"structNebulaNugget.PledgeInfo\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"UnlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumNebulaNugget.PledgePeriod\"},{\"name\":\"isPledged\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"halvingCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isContract\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mine\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"miningReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextHalvingAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonceBelongsTo\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pledge\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumNebulaNugget.PledgePeriod\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"poolList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"presaleApplyRefund\",\"inputs\":[{\"name\":\"_burnNNGAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"presaleReceivedBNB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"presaleReceivedNNG\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardPerMinute\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFactory\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolAddress\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_switch\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTradeOpen\",\"inputs\":[{\"name\":\"_switch\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMined\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tradeIsOpen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userLastMintTime\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userPledges\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"UnlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint8\",\"internalType\":\"enumNebulaNugget.PledgePeriod\"},{\"name\":\"isPledged\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEther\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MiningRewardAdjusted\",\"inputs\":[{\"name\":\"oldReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newReward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"halvingCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nextHalvingAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MiningSuccessful\",\"inputs\":[{\"name\":\"miner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_factory\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenClaimed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenPledged\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"unlockTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BlockMiningLimit\",\"inputs\":[{\"name\":\"blockNum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastMiningBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MiningHasCompleted\",\"inputs\":[{\"name\":\"curMining\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxMining\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MinuteRewardLimit\",\"inputs\":[{\"name\":\"blockMinute\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustPledgeToken\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NonceIsAlreadyUsed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PledgeAmountLimit\",\"inputs\":[{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"VerifyPowFailed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// NngABI is the input ABI used to generate the binding from.
// Deprecated: Use NngMetaData.ABI instead.
var NngABI = NngMetaData.ABI

// Nng is an auto generated Go binding around an Ethereum contract.
type Nng struct {
	NngCaller     // Read-only binding to the contract
	NngTransactor // Write-only binding to the contract
	NngFilterer   // Log filterer for contract events
}

// NngCaller is an auto generated read-only Go binding around an Ethereum contract.
type NngCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NngTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NngTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NngFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NngFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NngSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NngSession struct {
	Contract     *Nng              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NngCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NngCallerSession struct {
	Contract *NngCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NngTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NngTransactorSession struct {
	Contract     *NngTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NngRaw is an auto generated low-level Go binding around an Ethereum contract.
type NngRaw struct {
	Contract *Nng // Generic contract binding to access the raw methods on
}

// NngCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NngCallerRaw struct {
	Contract *NngCaller // Generic read-only contract binding to access the raw methods on
}

// NngTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NngTransactorRaw struct {
	Contract *NngTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNng creates a new instance of Nng, bound to a specific deployed contract.
func NewNng(address common.Address, backend bind.ContractBackend) (*Nng, error) {
	contract, err := bindNng(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nng{NngCaller: NngCaller{contract: contract}, NngTransactor: NngTransactor{contract: contract}, NngFilterer: NngFilterer{contract: contract}}, nil
}

// NewNngCaller creates a new read-only instance of Nng, bound to a specific deployed contract.
func NewNngCaller(address common.Address, caller bind.ContractCaller) (*NngCaller, error) {
	contract, err := bindNng(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NngCaller{contract: contract}, nil
}

// NewNngTransactor creates a new write-only instance of Nng, bound to a specific deployed contract.
func NewNngTransactor(address common.Address, transactor bind.ContractTransactor) (*NngTransactor, error) {
	contract, err := bindNng(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NngTransactor{contract: contract}, nil
}

// NewNngFilterer creates a new log filterer instance of Nng, bound to a specific deployed contract.
func NewNngFilterer(address common.Address, filterer bind.ContractFilterer) (*NngFilterer, error) {
	contract, err := bindNng(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NngFilterer{contract: contract}, nil
}

// bindNng binds a generic wrapper to an already deployed contract.
func bindNng(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NngMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nng *NngRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nng.Contract.NngCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nng *NngRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nng.Contract.NngTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nng *NngRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nng.Contract.NngTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nng *NngCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nng.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nng *NngTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nng.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nng *NngTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nng.Contract.contract.Transact(opts, method, params...)
}

// AmountRewardFactor is a free data retrieval call binding the contract method 0x48a1d37d.
//
// Solidity: function AmountRewardFactor() view returns(uint256)
func (_Nng *NngCaller) AmountRewardFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "AmountRewardFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountRewardFactor is a free data retrieval call binding the contract method 0x48a1d37d.
//
// Solidity: function AmountRewardFactor() view returns(uint256)
func (_Nng *NngSession) AmountRewardFactor() (*big.Int, error) {
	return _Nng.Contract.AmountRewardFactor(&_Nng.CallOpts)
}

// AmountRewardFactor is a free data retrieval call binding the contract method 0x48a1d37d.
//
// Solidity: function AmountRewardFactor() view returns(uint256)
func (_Nng *NngCallerSession) AmountRewardFactor() (*big.Int, error) {
	return _Nng.Contract.AmountRewardFactor(&_Nng.CallOpts)
}

// HalvingInterval is a free data retrieval call binding the contract method 0x3dcdba39.
//
// Solidity: function HalvingInterval() view returns(uint256)
func (_Nng *NngCaller) HalvingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "HalvingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalvingInterval is a free data retrieval call binding the contract method 0x3dcdba39.
//
// Solidity: function HalvingInterval() view returns(uint256)
func (_Nng *NngSession) HalvingInterval() (*big.Int, error) {
	return _Nng.Contract.HalvingInterval(&_Nng.CallOpts)
}

// HalvingInterval is a free data retrieval call binding the contract method 0x3dcdba39.
//
// Solidity: function HalvingInterval() view returns(uint256)
func (_Nng *NngCallerSession) HalvingInterval() (*big.Int, error) {
	return _Nng.Contract.HalvingInterval(&_Nng.CallOpts)
}

// MAXHALVINGS is a free data retrieval call binding the contract method 0x36f1e6c0.
//
// Solidity: function MAX_HALVINGS() view returns(uint256)
func (_Nng *NngCaller) MAXHALVINGS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "MAX_HALVINGS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXHALVINGS is a free data retrieval call binding the contract method 0x36f1e6c0.
//
// Solidity: function MAX_HALVINGS() view returns(uint256)
func (_Nng *NngSession) MAXHALVINGS() (*big.Int, error) {
	return _Nng.Contract.MAXHALVINGS(&_Nng.CallOpts)
}

// MAXHALVINGS is a free data retrieval call binding the contract method 0x36f1e6c0.
//
// Solidity: function MAX_HALVINGS() view returns(uint256)
func (_Nng *NngCallerSession) MAXHALVINGS() (*big.Int, error) {
	return _Nng.Contract.MAXHALVINGS(&_Nng.CallOpts)
}

// MaxPledgeAmount is a free data retrieval call binding the contract method 0x606a96fa.
//
// Solidity: function MaxPledgeAmount() view returns(uint256)
func (_Nng *NngCaller) MaxPledgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "MaxPledgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPledgeAmount is a free data retrieval call binding the contract method 0x606a96fa.
//
// Solidity: function MaxPledgeAmount() view returns(uint256)
func (_Nng *NngSession) MaxPledgeAmount() (*big.Int, error) {
	return _Nng.Contract.MaxPledgeAmount(&_Nng.CallOpts)
}

// MaxPledgeAmount is a free data retrieval call binding the contract method 0x606a96fa.
//
// Solidity: function MaxPledgeAmount() view returns(uint256)
func (_Nng *NngCallerSession) MaxPledgeAmount() (*big.Int, error) {
	return _Nng.Contract.MaxPledgeAmount(&_Nng.CallOpts)
}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xdcbc36aa.
//
// Solidity: function MinPledgeAmount() view returns(uint256)
func (_Nng *NngCaller) MinPledgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "MinPledgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xdcbc36aa.
//
// Solidity: function MinPledgeAmount() view returns(uint256)
func (_Nng *NngSession) MinPledgeAmount() (*big.Int, error) {
	return _Nng.Contract.MinPledgeAmount(&_Nng.CallOpts)
}

// MinPledgeAmount is a free data retrieval call binding the contract method 0xdcbc36aa.
//
// Solidity: function MinPledgeAmount() view returns(uint256)
func (_Nng *NngCallerSession) MinPledgeAmount() (*big.Int, error) {
	return _Nng.Contract.MinPledgeAmount(&_Nng.CallOpts)
}

// PeriodRewardFactor is a free data retrieval call binding the contract method 0xdf5d1e46.
//
// Solidity: function PeriodRewardFactor() view returns(uint256)
func (_Nng *NngCaller) PeriodRewardFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "PeriodRewardFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodRewardFactor is a free data retrieval call binding the contract method 0xdf5d1e46.
//
// Solidity: function PeriodRewardFactor() view returns(uint256)
func (_Nng *NngSession) PeriodRewardFactor() (*big.Int, error) {
	return _Nng.Contract.PeriodRewardFactor(&_Nng.CallOpts)
}

// PeriodRewardFactor is a free data retrieval call binding the contract method 0xdf5d1e46.
//
// Solidity: function PeriodRewardFactor() view returns(uint256)
func (_Nng *NngCallerSession) PeriodRewardFactor() (*big.Int, error) {
	return _Nng.Contract.PeriodRewardFactor(&_Nng.CallOpts)
}

// UserMiningLimitOfBlock is a free data retrieval call binding the contract method 0x2a1b4de7.
//
// Solidity: function UserMiningLimitOfBlock() view returns(uint256)
func (_Nng *NngCaller) UserMiningLimitOfBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "UserMiningLimitOfBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMiningLimitOfBlock is a free data retrieval call binding the contract method 0x2a1b4de7.
//
// Solidity: function UserMiningLimitOfBlock() view returns(uint256)
func (_Nng *NngSession) UserMiningLimitOfBlock() (*big.Int, error) {
	return _Nng.Contract.UserMiningLimitOfBlock(&_Nng.CallOpts)
}

// UserMiningLimitOfBlock is a free data retrieval call binding the contract method 0x2a1b4de7.
//
// Solidity: function UserMiningLimitOfBlock() view returns(uint256)
func (_Nng *NngCallerSession) UserMiningLimitOfBlock() (*big.Int, error) {
	return _Nng.Contract.UserMiningLimitOfBlock(&_Nng.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nng *NngCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nng *NngSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nng.Contract.Allowance(&_Nng.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nng *NngCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nng.Contract.Allowance(&_Nng.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nng *NngCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nng *NngSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nng.Contract.BalanceOf(&_Nng.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nng *NngCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nng.Contract.BalanceOf(&_Nng.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nng *NngCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nng *NngSession) Decimals() (uint8, error) {
	return _Nng.Contract.Decimals(&_Nng.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nng *NngCallerSession) Decimals() (uint8, error) {
	return _Nng.Contract.Decimals(&_Nng.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Nng *NngCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Nng *NngSession) Difficulty() (*big.Int, error) {
	return _Nng.Contract.Difficulty(&_Nng.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_Nng *NngCallerSession) Difficulty() (*big.Int, error) {
	return _Nng.Contract.Difficulty(&_Nng.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Nng *NngCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Nng *NngSession) Factory() (common.Address, error) {
	return _Nng.Contract.Factory(&_Nng.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Nng *NngCallerSession) Factory() (common.Address, error) {
	return _Nng.Contract.Factory(&_Nng.CallOpts)
}

// GetPledgeMiningReward is a free data retrieval call binding the contract method 0x9b2727ab.
//
// Solidity: function getPledgeMiningReward((uint256,uint256,uint8,bool) s) view returns(uint256)
func (_Nng *NngCaller) GetPledgeMiningReward(opts *bind.CallOpts, s NebulaNuggetPledgeInfo) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "getPledgeMiningReward", s)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledgeMiningReward is a free data retrieval call binding the contract method 0x9b2727ab.
//
// Solidity: function getPledgeMiningReward((uint256,uint256,uint8,bool) s) view returns(uint256)
func (_Nng *NngSession) GetPledgeMiningReward(s NebulaNuggetPledgeInfo) (*big.Int, error) {
	return _Nng.Contract.GetPledgeMiningReward(&_Nng.CallOpts, s)
}

// GetPledgeMiningReward is a free data retrieval call binding the contract method 0x9b2727ab.
//
// Solidity: function getPledgeMiningReward((uint256,uint256,uint8,bool) s) view returns(uint256)
func (_Nng *NngCallerSession) GetPledgeMiningReward(s NebulaNuggetPledgeInfo) (*big.Int, error) {
	return _Nng.Contract.GetPledgeMiningReward(&_Nng.CallOpts, s)
}

// HalvingCount is a free data retrieval call binding the contract method 0xfc84d913.
//
// Solidity: function halvingCount() view returns(uint256)
func (_Nng *NngCaller) HalvingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "halvingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HalvingCount is a free data retrieval call binding the contract method 0xfc84d913.
//
// Solidity: function halvingCount() view returns(uint256)
func (_Nng *NngSession) HalvingCount() (*big.Int, error) {
	return _Nng.Contract.HalvingCount(&_Nng.CallOpts)
}

// HalvingCount is a free data retrieval call binding the contract method 0xfc84d913.
//
// Solidity: function halvingCount() view returns(uint256)
func (_Nng *NngCallerSession) HalvingCount() (*big.Int, error) {
	return _Nng.Contract.HalvingCount(&_Nng.CallOpts)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_Nng *NngCaller) IsContract(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "isContract", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_Nng *NngSession) IsContract(_addr common.Address) (bool, error) {
	return _Nng.Contract.IsContract(&_Nng.CallOpts, _addr)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address _addr) view returns(bool)
func (_Nng *NngCallerSession) IsContract(_addr common.Address) (bool, error) {
	return _Nng.Contract.IsContract(&_Nng.CallOpts, _addr)
}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_Nng *NngCaller) MiningReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "miningReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_Nng *NngSession) MiningReward() (*big.Int, error) {
	return _Nng.Contract.MiningReward(&_Nng.CallOpts)
}

// MiningReward is a free data retrieval call binding the contract method 0x4ac2d103.
//
// Solidity: function miningReward() view returns(uint256)
func (_Nng *NngCallerSession) MiningReward() (*big.Int, error) {
	return _Nng.Contract.MiningReward(&_Nng.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nng *NngCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nng *NngSession) Name() (string, error) {
	return _Nng.Contract.Name(&_Nng.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nng *NngCallerSession) Name() (string, error) {
	return _Nng.Contract.Name(&_Nng.CallOpts)
}

// NextHalvingAt is a free data retrieval call binding the contract method 0x016a6241.
//
// Solidity: function nextHalvingAt() view returns(uint256)
func (_Nng *NngCaller) NextHalvingAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "nextHalvingAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextHalvingAt is a free data retrieval call binding the contract method 0x016a6241.
//
// Solidity: function nextHalvingAt() view returns(uint256)
func (_Nng *NngSession) NextHalvingAt() (*big.Int, error) {
	return _Nng.Contract.NextHalvingAt(&_Nng.CallOpts)
}

// NextHalvingAt is a free data retrieval call binding the contract method 0x016a6241.
//
// Solidity: function nextHalvingAt() view returns(uint256)
func (_Nng *NngCallerSession) NextHalvingAt() (*big.Int, error) {
	return _Nng.Contract.NextHalvingAt(&_Nng.CallOpts)
}

// NonceBelongsTo is a free data retrieval call binding the contract method 0xd84ab1d9.
//
// Solidity: function nonceBelongsTo(bytes32 ) view returns(address)
func (_Nng *NngCaller) NonceBelongsTo(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "nonceBelongsTo", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonceBelongsTo is a free data retrieval call binding the contract method 0xd84ab1d9.
//
// Solidity: function nonceBelongsTo(bytes32 ) view returns(address)
func (_Nng *NngSession) NonceBelongsTo(arg0 [32]byte) (common.Address, error) {
	return _Nng.Contract.NonceBelongsTo(&_Nng.CallOpts, arg0)
}

// NonceBelongsTo is a free data retrieval call binding the contract method 0xd84ab1d9.
//
// Solidity: function nonceBelongsTo(bytes32 ) view returns(address)
func (_Nng *NngCallerSession) NonceBelongsTo(arg0 [32]byte) (common.Address, error) {
	return _Nng.Contract.NonceBelongsTo(&_Nng.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nng *NngCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nng *NngSession) Owner() (common.Address, error) {
	return _Nng.Contract.Owner(&_Nng.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nng *NngCallerSession) Owner() (common.Address, error) {
	return _Nng.Contract.Owner(&_Nng.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0xa36cc981.
//
// Solidity: function poolList(address ) view returns(bool)
func (_Nng *NngCaller) PoolList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "poolList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0xa36cc981.
//
// Solidity: function poolList(address ) view returns(bool)
func (_Nng *NngSession) PoolList(arg0 common.Address) (bool, error) {
	return _Nng.Contract.PoolList(&_Nng.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0xa36cc981.
//
// Solidity: function poolList(address ) view returns(bool)
func (_Nng *NngCallerSession) PoolList(arg0 common.Address) (bool, error) {
	return _Nng.Contract.PoolList(&_Nng.CallOpts, arg0)
}

// PresaleReceivedBNB is a free data retrieval call binding the contract method 0xc330dc23.
//
// Solidity: function presaleReceivedBNB() view returns(uint256)
func (_Nng *NngCaller) PresaleReceivedBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "presaleReceivedBNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleReceivedBNB is a free data retrieval call binding the contract method 0xc330dc23.
//
// Solidity: function presaleReceivedBNB() view returns(uint256)
func (_Nng *NngSession) PresaleReceivedBNB() (*big.Int, error) {
	return _Nng.Contract.PresaleReceivedBNB(&_Nng.CallOpts)
}

// PresaleReceivedBNB is a free data retrieval call binding the contract method 0xc330dc23.
//
// Solidity: function presaleReceivedBNB() view returns(uint256)
func (_Nng *NngCallerSession) PresaleReceivedBNB() (*big.Int, error) {
	return _Nng.Contract.PresaleReceivedBNB(&_Nng.CallOpts)
}

// PresaleReceivedNNG is a free data retrieval call binding the contract method 0x697e63c1.
//
// Solidity: function presaleReceivedNNG(address ) view returns(uint256)
func (_Nng *NngCaller) PresaleReceivedNNG(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "presaleReceivedNNG", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleReceivedNNG is a free data retrieval call binding the contract method 0x697e63c1.
//
// Solidity: function presaleReceivedNNG(address ) view returns(uint256)
func (_Nng *NngSession) PresaleReceivedNNG(arg0 common.Address) (*big.Int, error) {
	return _Nng.Contract.PresaleReceivedNNG(&_Nng.CallOpts, arg0)
}

// PresaleReceivedNNG is a free data retrieval call binding the contract method 0x697e63c1.
//
// Solidity: function presaleReceivedNNG(address ) view returns(uint256)
func (_Nng *NngCallerSession) PresaleReceivedNNG(arg0 common.Address) (*big.Int, error) {
	return _Nng.Contract.PresaleReceivedNNG(&_Nng.CallOpts, arg0)
}

// RewardPerMinute is a free data retrieval call binding the contract method 0x791bb6c0.
//
// Solidity: function rewardPerMinute(uint256 ) view returns(uint256)
func (_Nng *NngCaller) RewardPerMinute(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "rewardPerMinute", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerMinute is a free data retrieval call binding the contract method 0x791bb6c0.
//
// Solidity: function rewardPerMinute(uint256 ) view returns(uint256)
func (_Nng *NngSession) RewardPerMinute(arg0 *big.Int) (*big.Int, error) {
	return _Nng.Contract.RewardPerMinute(&_Nng.CallOpts, arg0)
}

// RewardPerMinute is a free data retrieval call binding the contract method 0x791bb6c0.
//
// Solidity: function rewardPerMinute(uint256 ) view returns(uint256)
func (_Nng *NngCallerSession) RewardPerMinute(arg0 *big.Int) (*big.Int, error) {
	return _Nng.Contract.RewardPerMinute(&_Nng.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nng *NngCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nng *NngSession) Symbol() (string, error) {
	return _Nng.Contract.Symbol(&_Nng.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nng *NngCallerSession) Symbol() (string, error) {
	return _Nng.Contract.Symbol(&_Nng.CallOpts)
}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_Nng *NngCaller) TotalMined(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "totalMined")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_Nng *NngSession) TotalMined() (*big.Int, error) {
	return _Nng.Contract.TotalMined(&_Nng.CallOpts)
}

// TotalMined is a free data retrieval call binding the contract method 0x5556db65.
//
// Solidity: function totalMined() view returns(uint256)
func (_Nng *NngCallerSession) TotalMined() (*big.Int, error) {
	return _Nng.Contract.TotalMined(&_Nng.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nng *NngCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nng *NngSession) TotalSupply() (*big.Int, error) {
	return _Nng.Contract.TotalSupply(&_Nng.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nng *NngCallerSession) TotalSupply() (*big.Int, error) {
	return _Nng.Contract.TotalSupply(&_Nng.CallOpts)
}

// TradeIsOpen is a free data retrieval call binding the contract method 0x8f327f70.
//
// Solidity: function tradeIsOpen() view returns(bool)
func (_Nng *NngCaller) TradeIsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "tradeIsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TradeIsOpen is a free data retrieval call binding the contract method 0x8f327f70.
//
// Solidity: function tradeIsOpen() view returns(bool)
func (_Nng *NngSession) TradeIsOpen() (bool, error) {
	return _Nng.Contract.TradeIsOpen(&_Nng.CallOpts)
}

// TradeIsOpen is a free data retrieval call binding the contract method 0x8f327f70.
//
// Solidity: function tradeIsOpen() view returns(bool)
func (_Nng *NngCallerSession) TradeIsOpen() (bool, error) {
	return _Nng.Contract.TradeIsOpen(&_Nng.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nng *NngCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nng *NngSession) Treasury() (common.Address, error) {
	return _Nng.Contract.Treasury(&_Nng.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Nng *NngCallerSession) Treasury() (common.Address, error) {
	return _Nng.Contract.Treasury(&_Nng.CallOpts)
}

// UserLastMintTime is a free data retrieval call binding the contract method 0x3ebb032a.
//
// Solidity: function userLastMintTime(address ) view returns(uint256)
func (_Nng *NngCaller) UserLastMintTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "userLastMintTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLastMintTime is a free data retrieval call binding the contract method 0x3ebb032a.
//
// Solidity: function userLastMintTime(address ) view returns(uint256)
func (_Nng *NngSession) UserLastMintTime(arg0 common.Address) (*big.Int, error) {
	return _Nng.Contract.UserLastMintTime(&_Nng.CallOpts, arg0)
}

// UserLastMintTime is a free data retrieval call binding the contract method 0x3ebb032a.
//
// Solidity: function userLastMintTime(address ) view returns(uint256)
func (_Nng *NngCallerSession) UserLastMintTime(arg0 common.Address) (*big.Int, error) {
	return _Nng.Contract.UserLastMintTime(&_Nng.CallOpts, arg0)
}

// UserPledges is a free data retrieval call binding the contract method 0x8272b1cf.
//
// Solidity: function userPledges(address ) view returns(uint256 amount, uint256 UnlockTime, uint8 period, bool isPledged)
func (_Nng *NngCaller) UserPledges(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
	Period     uint8
	IsPledged  bool
}, error) {
	var out []interface{}
	err := _Nng.contract.Call(opts, &out, "userPledges", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		UnlockTime *big.Int
		Period     uint8
		IsPledged  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Period = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.IsPledged = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// UserPledges is a free data retrieval call binding the contract method 0x8272b1cf.
//
// Solidity: function userPledges(address ) view returns(uint256 amount, uint256 UnlockTime, uint8 period, bool isPledged)
func (_Nng *NngSession) UserPledges(arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
	Period     uint8
	IsPledged  bool
}, error) {
	return _Nng.Contract.UserPledges(&_Nng.CallOpts, arg0)
}

// UserPledges is a free data retrieval call binding the contract method 0x8272b1cf.
//
// Solidity: function userPledges(address ) view returns(uint256 amount, uint256 UnlockTime, uint8 period, bool isPledged)
func (_Nng *NngCallerSession) UserPledges(arg0 common.Address) (struct {
	Amount     *big.Int
	UnlockTime *big.Int
	Period     uint8
	IsPledged  bool
}, error) {
	return _Nng.Contract.UserPledges(&_Nng.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nng *NngTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nng *NngSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.Approve(&_Nng.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Nng *NngTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.Approve(&_Nng.TransactOpts, spender, amount)
}

// ClaimPledge is a paid mutator transaction binding the contract method 0xc632c9a0.
//
// Solidity: function claimPledge() returns()
func (_Nng *NngTransactor) ClaimPledge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "claimPledge")
}

// ClaimPledge is a paid mutator transaction binding the contract method 0xc632c9a0.
//
// Solidity: function claimPledge() returns()
func (_Nng *NngSession) ClaimPledge() (*types.Transaction, error) {
	return _Nng.Contract.ClaimPledge(&_Nng.TransactOpts)
}

// ClaimPledge is a paid mutator transaction binding the contract method 0xc632c9a0.
//
// Solidity: function claimPledge() returns()
func (_Nng *NngTransactorSession) ClaimPledge() (*types.Transaction, error) {
	return _Nng.Contract.ClaimPledge(&_Nng.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Nng *NngTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Nng *NngSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.DecreaseAllowance(&_Nng.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Nng *NngTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.DecreaseAllowance(&_Nng.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Nng *NngTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Nng *NngSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.IncreaseAllowance(&_Nng.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Nng *NngTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.IncreaseAllowance(&_Nng.TransactOpts, spender, addedValue)
}

// Mine is a paid mutator transaction binding the contract method 0x373d418a.
//
// Solidity: function mine(bytes32 nonce, address recipient) returns()
func (_Nng *NngTransactor) Mine(opts *bind.TransactOpts, nonce [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "mine", nonce, recipient)
}

// Mine is a paid mutator transaction binding the contract method 0x373d418a.
//
// Solidity: function mine(bytes32 nonce, address recipient) returns()
func (_Nng *NngSession) Mine(nonce [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Nng.Contract.Mine(&_Nng.TransactOpts, nonce, recipient)
}

// Mine is a paid mutator transaction binding the contract method 0x373d418a.
//
// Solidity: function mine(bytes32 nonce, address recipient) returns()
func (_Nng *NngTransactorSession) Mine(nonce [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Nng.Contract.Mine(&_Nng.TransactOpts, nonce, recipient)
}

// Pledge is a paid mutator transaction binding the contract method 0xf9fb3e0e.
//
// Solidity: function pledge(uint256 amount, uint8 period) returns()
func (_Nng *NngTransactor) Pledge(opts *bind.TransactOpts, amount *big.Int, period uint8) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "pledge", amount, period)
}

// Pledge is a paid mutator transaction binding the contract method 0xf9fb3e0e.
//
// Solidity: function pledge(uint256 amount, uint8 period) returns()
func (_Nng *NngSession) Pledge(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _Nng.Contract.Pledge(&_Nng.TransactOpts, amount, period)
}

// Pledge is a paid mutator transaction binding the contract method 0xf9fb3e0e.
//
// Solidity: function pledge(uint256 amount, uint8 period) returns()
func (_Nng *NngTransactorSession) Pledge(amount *big.Int, period uint8) (*types.Transaction, error) {
	return _Nng.Contract.Pledge(&_Nng.TransactOpts, amount, period)
}

// PresaleApplyRefund is a paid mutator transaction binding the contract method 0x3fd3afe2.
//
// Solidity: function presaleApplyRefund(uint256 _burnNNGAmount) returns()
func (_Nng *NngTransactor) PresaleApplyRefund(opts *bind.TransactOpts, _burnNNGAmount *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "presaleApplyRefund", _burnNNGAmount)
}

// PresaleApplyRefund is a paid mutator transaction binding the contract method 0x3fd3afe2.
//
// Solidity: function presaleApplyRefund(uint256 _burnNNGAmount) returns()
func (_Nng *NngSession) PresaleApplyRefund(_burnNNGAmount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.PresaleApplyRefund(&_Nng.TransactOpts, _burnNNGAmount)
}

// PresaleApplyRefund is a paid mutator transaction binding the contract method 0x3fd3afe2.
//
// Solidity: function presaleApplyRefund(uint256 _burnNNGAmount) returns()
func (_Nng *NngTransactorSession) PresaleApplyRefund(_burnNNGAmount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.PresaleApplyRefund(&_Nng.TransactOpts, _burnNNGAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nng *NngTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nng *NngSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nng.Contract.RenounceOwnership(&_Nng.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nng *NngTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nng.Contract.RenounceOwnership(&_Nng.TransactOpts)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address _factory) returns()
func (_Nng *NngTransactor) SetFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "setFactory", _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address _factory) returns()
func (_Nng *NngSession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _Nng.Contract.SetFactory(&_Nng.TransactOpts, _factory)
}

// SetFactory is a paid mutator transaction binding the contract method 0x5bb47808.
//
// Solidity: function setFactory(address _factory) returns()
func (_Nng *NngTransactorSession) SetFactory(_factory common.Address) (*types.Transaction, error) {
	return _Nng.Contract.SetFactory(&_Nng.TransactOpts, _factory)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address _pool, bool _switch) returns()
func (_Nng *NngTransactor) SetPoolAddress(opts *bind.TransactOpts, _pool common.Address, _switch bool) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "setPoolAddress", _pool, _switch)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address _pool, bool _switch) returns()
func (_Nng *NngSession) SetPoolAddress(_pool common.Address, _switch bool) (*types.Transaction, error) {
	return _Nng.Contract.SetPoolAddress(&_Nng.TransactOpts, _pool, _switch)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address _pool, bool _switch) returns()
func (_Nng *NngTransactorSession) SetPoolAddress(_pool common.Address, _switch bool) (*types.Transaction, error) {
	return _Nng.Contract.SetPoolAddress(&_Nng.TransactOpts, _pool, _switch)
}

// SetTradeOpen is a paid mutator transaction binding the contract method 0xcca8fd4a.
//
// Solidity: function setTradeOpen(bool _switch) returns()
func (_Nng *NngTransactor) SetTradeOpen(opts *bind.TransactOpts, _switch bool) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "setTradeOpen", _switch)
}

// SetTradeOpen is a paid mutator transaction binding the contract method 0xcca8fd4a.
//
// Solidity: function setTradeOpen(bool _switch) returns()
func (_Nng *NngSession) SetTradeOpen(_switch bool) (*types.Transaction, error) {
	return _Nng.Contract.SetTradeOpen(&_Nng.TransactOpts, _switch)
}

// SetTradeOpen is a paid mutator transaction binding the contract method 0xcca8fd4a.
//
// Solidity: function setTradeOpen(bool _switch) returns()
func (_Nng *NngTransactorSession) SetTradeOpen(_switch bool) (*types.Transaction, error) {
	return _Nng.Contract.SetTradeOpen(&_Nng.TransactOpts, _switch)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nng *NngTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nng *NngSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.Transfer(&_Nng.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Nng *NngTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.Transfer(&_Nng.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nng *NngTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nng *NngSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.TransferFrom(&_Nng.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Nng *NngTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.TransferFrom(&_Nng.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nng *NngTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nng *NngSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nng.Contract.TransferOwnership(&_Nng.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nng *NngTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nng.Contract.TransferOwnership(&_Nng.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address tokenAddress, uint256 amount) returns()
func (_Nng *NngTransactor) WithdrawERC20(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "withdrawERC20", tokenAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address tokenAddress, uint256 amount) returns()
func (_Nng *NngSession) WithdrawERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.WithdrawERC20(&_Nng.TransactOpts, tokenAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address tokenAddress, uint256 amount) returns()
func (_Nng *NngTransactorSession) WithdrawERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Nng.Contract.WithdrawERC20(&_Nng.TransactOpts, tokenAddress, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _to) returns()
func (_Nng *NngTransactor) WithdrawEther(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Nng.contract.Transact(opts, "withdrawEther", _to)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _to) returns()
func (_Nng *NngSession) WithdrawEther(_to common.Address) (*types.Transaction, error) {
	return _Nng.Contract.WithdrawEther(&_Nng.TransactOpts, _to)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xaf933b57.
//
// Solidity: function withdrawEther(address _to) returns()
func (_Nng *NngTransactorSession) WithdrawEther(_to common.Address) (*types.Transaction, error) {
	return _Nng.Contract.WithdrawEther(&_Nng.TransactOpts, _to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nng *NngTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Nng.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nng *NngSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nng.Contract.Fallback(&_Nng.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Nng *NngTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Nng.Contract.Fallback(&_Nng.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nng *NngTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nng.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nng *NngSession) Receive() (*types.Transaction, error) {
	return _Nng.Contract.Receive(&_Nng.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nng *NngTransactorSession) Receive() (*types.Transaction, error) {
	return _Nng.Contract.Receive(&_Nng.TransactOpts)
}

// NngApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nng contract.
type NngApprovalIterator struct {
	Event *NngApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngApproval represents a Approval event raised by the Nng contract.
type NngApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nng *NngFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NngApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NngApprovalIterator{contract: _Nng.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nng *NngFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NngApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngApproval)
				if err := _Nng.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nng *NngFilterer) ParseApproval(log types.Log) (*NngApproval, error) {
	event := new(NngApproval)
	if err := _Nng.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngMiningRewardAdjustedIterator is returned from FilterMiningRewardAdjusted and is used to iterate over the raw logs and unpacked data for MiningRewardAdjusted events raised by the Nng contract.
type NngMiningRewardAdjustedIterator struct {
	Event *NngMiningRewardAdjusted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngMiningRewardAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngMiningRewardAdjusted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngMiningRewardAdjusted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngMiningRewardAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngMiningRewardAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngMiningRewardAdjusted represents a MiningRewardAdjusted event raised by the Nng contract.
type NngMiningRewardAdjusted struct {
	OldReward     *big.Int
	NewReward     *big.Int
	HalvingCount  *big.Int
	NextHalvingAt *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardAdjusted is a free log retrieval operation binding the contract event 0x8ee9dc90a5216f4cd80013a8125361f97419228a3b2f91b8829e8614851c8143.
//
// Solidity: event MiningRewardAdjusted(uint256 oldReward, uint256 newReward, uint256 halvingCount, uint256 nextHalvingAt)
func (_Nng *NngFilterer) FilterMiningRewardAdjusted(opts *bind.FilterOpts) (*NngMiningRewardAdjustedIterator, error) {

	logs, sub, err := _Nng.contract.FilterLogs(opts, "MiningRewardAdjusted")
	if err != nil {
		return nil, err
	}
	return &NngMiningRewardAdjustedIterator{contract: _Nng.contract, event: "MiningRewardAdjusted", logs: logs, sub: sub}, nil
}

// WatchMiningRewardAdjusted is a free log subscription operation binding the contract event 0x8ee9dc90a5216f4cd80013a8125361f97419228a3b2f91b8829e8614851c8143.
//
// Solidity: event MiningRewardAdjusted(uint256 oldReward, uint256 newReward, uint256 halvingCount, uint256 nextHalvingAt)
func (_Nng *NngFilterer) WatchMiningRewardAdjusted(opts *bind.WatchOpts, sink chan<- *NngMiningRewardAdjusted) (event.Subscription, error) {

	logs, sub, err := _Nng.contract.WatchLogs(opts, "MiningRewardAdjusted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngMiningRewardAdjusted)
				if err := _Nng.contract.UnpackLog(event, "MiningRewardAdjusted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningRewardAdjusted is a log parse operation binding the contract event 0x8ee9dc90a5216f4cd80013a8125361f97419228a3b2f91b8829e8614851c8143.
//
// Solidity: event MiningRewardAdjusted(uint256 oldReward, uint256 newReward, uint256 halvingCount, uint256 nextHalvingAt)
func (_Nng *NngFilterer) ParseMiningRewardAdjusted(log types.Log) (*NngMiningRewardAdjusted, error) {
	event := new(NngMiningRewardAdjusted)
	if err := _Nng.contract.UnpackLog(event, "MiningRewardAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngMiningSuccessfulIterator is returned from FilterMiningSuccessful and is used to iterate over the raw logs and unpacked data for MiningSuccessful events raised by the Nng contract.
type NngMiningSuccessfulIterator struct {
	Event *NngMiningSuccessful // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngMiningSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngMiningSuccessful)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngMiningSuccessful)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngMiningSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngMiningSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngMiningSuccessful represents a MiningSuccessful event raised by the Nng contract.
type NngMiningSuccessful struct {
	Miner   common.Address
	Reward  *big.Int
	Factory common.Address
	Nonce   [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMiningSuccessful is a free log retrieval operation binding the contract event 0xf1d4f3ff248de94d623d4070f9d7e4f40aad7e87cdeef0df4598c5555dabe261.
//
// Solidity: event MiningSuccessful(address indexed miner, uint256 reward, address _factory, bytes32 nonce)
func (_Nng *NngFilterer) FilterMiningSuccessful(opts *bind.FilterOpts, miner []common.Address) (*NngMiningSuccessfulIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "MiningSuccessful", minerRule)
	if err != nil {
		return nil, err
	}
	return &NngMiningSuccessfulIterator{contract: _Nng.contract, event: "MiningSuccessful", logs: logs, sub: sub}, nil
}

// WatchMiningSuccessful is a free log subscription operation binding the contract event 0xf1d4f3ff248de94d623d4070f9d7e4f40aad7e87cdeef0df4598c5555dabe261.
//
// Solidity: event MiningSuccessful(address indexed miner, uint256 reward, address _factory, bytes32 nonce)
func (_Nng *NngFilterer) WatchMiningSuccessful(opts *bind.WatchOpts, sink chan<- *NngMiningSuccessful, miner []common.Address) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "MiningSuccessful", minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngMiningSuccessful)
				if err := _Nng.contract.UnpackLog(event, "MiningSuccessful", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMiningSuccessful is a log parse operation binding the contract event 0xf1d4f3ff248de94d623d4070f9d7e4f40aad7e87cdeef0df4598c5555dabe261.
//
// Solidity: event MiningSuccessful(address indexed miner, uint256 reward, address _factory, bytes32 nonce)
func (_Nng *NngFilterer) ParseMiningSuccessful(log types.Log) (*NngMiningSuccessful, error) {
	event := new(NngMiningSuccessful)
	if err := _Nng.contract.UnpackLog(event, "MiningSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nng contract.
type NngOwnershipTransferredIterator struct {
	Event *NngOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngOwnershipTransferred represents a OwnershipTransferred event raised by the Nng contract.
type NngOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nng *NngFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NngOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NngOwnershipTransferredIterator{contract: _Nng.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nng *NngFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NngOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngOwnershipTransferred)
				if err := _Nng.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nng *NngFilterer) ParseOwnershipTransferred(log types.Log) (*NngOwnershipTransferred, error) {
	event := new(NngOwnershipTransferred)
	if err := _Nng.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngTokenClaimedIterator is returned from FilterTokenClaimed and is used to iterate over the raw logs and unpacked data for TokenClaimed events raised by the Nng contract.
type NngTokenClaimedIterator struct {
	Event *NngTokenClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngTokenClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngTokenClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngTokenClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngTokenClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngTokenClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngTokenClaimed represents a TokenClaimed event raised by the Nng contract.
type NngTokenClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenClaimed is a free log retrieval operation binding the contract event 0xe42df0d9493dfd0d7f69902c895b94c190a53e8c27876a86f45e7c997d9d8f7c.
//
// Solidity: event TokenClaimed(address indexed user, uint256 amount)
func (_Nng *NngFilterer) FilterTokenClaimed(opts *bind.FilterOpts, user []common.Address) (*NngTokenClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "TokenClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &NngTokenClaimedIterator{contract: _Nng.contract, event: "TokenClaimed", logs: logs, sub: sub}, nil
}

// WatchTokenClaimed is a free log subscription operation binding the contract event 0xe42df0d9493dfd0d7f69902c895b94c190a53e8c27876a86f45e7c997d9d8f7c.
//
// Solidity: event TokenClaimed(address indexed user, uint256 amount)
func (_Nng *NngFilterer) WatchTokenClaimed(opts *bind.WatchOpts, sink chan<- *NngTokenClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "TokenClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngTokenClaimed)
				if err := _Nng.contract.UnpackLog(event, "TokenClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenClaimed is a log parse operation binding the contract event 0xe42df0d9493dfd0d7f69902c895b94c190a53e8c27876a86f45e7c997d9d8f7c.
//
// Solidity: event TokenClaimed(address indexed user, uint256 amount)
func (_Nng *NngFilterer) ParseTokenClaimed(log types.Log) (*NngTokenClaimed, error) {
	event := new(NngTokenClaimed)
	if err := _Nng.contract.UnpackLog(event, "TokenClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngTokenPledgedIterator is returned from FilterTokenPledged and is used to iterate over the raw logs and unpacked data for TokenPledged events raised by the Nng contract.
type NngTokenPledgedIterator struct {
	Event *NngTokenPledged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngTokenPledgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngTokenPledged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngTokenPledged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngTokenPledgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngTokenPledgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngTokenPledged represents a TokenPledged event raised by the Nng contract.
type NngTokenPledged struct {
	User       common.Address
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenPledged is a free log retrieval operation binding the contract event 0x357769b0fe70c49f09b7a708db2d48a10325d26b2e0ea0d5e437d0cd78dc1e14.
//
// Solidity: event TokenPledged(address indexed user, uint256 amount, uint256 unlockTime)
func (_Nng *NngFilterer) FilterTokenPledged(opts *bind.FilterOpts, user []common.Address) (*NngTokenPledgedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "TokenPledged", userRule)
	if err != nil {
		return nil, err
	}
	return &NngTokenPledgedIterator{contract: _Nng.contract, event: "TokenPledged", logs: logs, sub: sub}, nil
}

// WatchTokenPledged is a free log subscription operation binding the contract event 0x357769b0fe70c49f09b7a708db2d48a10325d26b2e0ea0d5e437d0cd78dc1e14.
//
// Solidity: event TokenPledged(address indexed user, uint256 amount, uint256 unlockTime)
func (_Nng *NngFilterer) WatchTokenPledged(opts *bind.WatchOpts, sink chan<- *NngTokenPledged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "TokenPledged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngTokenPledged)
				if err := _Nng.contract.UnpackLog(event, "TokenPledged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenPledged is a log parse operation binding the contract event 0x357769b0fe70c49f09b7a708db2d48a10325d26b2e0ea0d5e437d0cd78dc1e14.
//
// Solidity: event TokenPledged(address indexed user, uint256 amount, uint256 unlockTime)
func (_Nng *NngFilterer) ParseTokenPledged(log types.Log) (*NngTokenPledged, error) {
	event := new(NngTokenPledged)
	if err := _Nng.contract.UnpackLog(event, "TokenPledged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NngTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nng contract.
type NngTransferIterator struct {
	Event *NngTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NngTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NngTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NngTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NngTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NngTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NngTransfer represents a Transfer event raised by the Nng contract.
type NngTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nng *NngFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NngTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nng.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NngTransferIterator{contract: _Nng.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nng *NngFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NngTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nng.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NngTransfer)
				if err := _Nng.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nng *NngFilterer) ParseTransfer(log types.Log) (*NngTransfer, error) {
	event := new(NngTransfer)
	if err := _Nng.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
