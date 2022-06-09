// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package biddingpk

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
)

// BiddingpkMetaData contains all meta data concerning the Biddingpk contract.
var BiddingpkMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AuctionAlreadyEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionNotYetEnded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"}],\"name\":\"BidNotHighEnough\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"HighestBidIncreased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHighestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gettotalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e104261001f9190610063565b6001819055506100b9565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061006e8261002a565b91506100798361002a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156100ae576100ad610034565b5b828201905092915050565b610752806100c86000396000f3fe6080604052600436106100865760003560e01c80634b449cba116100595780634b449cba1461010257806391f901571461012d578063d57bde7914610158578063e148919114610183578063e65d6b49146101ae57610086565b80631998aeef1461008b5780632a24f46c146100955780632cba772e146100ac5780634979440a146100d7575b600080fd5b6100936101d9565b005b3480156100a157600080fd5b506100aa610377565b005b3480156100b857600080fd5b506100c1610490565b6040516100ce9190610566565b60405180910390f35b3480156100e357600080fd5b506100ec61049a565b6040516100f99190610566565b60405180910390f35b34801561010e57600080fd5b506101176104a4565b6040516101249190610566565b60405180910390f35b34801561013957600080fd5b506101426104aa565b60405161014f91906105c2565b60405180910390f35b34801561016457600080fd5b5061016d6104d0565b60405161017a9190610566565b60405180910390f35b34801561018f57600080fd5b506101986104d6565b6040516101a59190610566565b60405180910390f35b3480156101ba57600080fd5b506101c36104dc565b6040516101d09190610566565b60405180910390f35b600034116101e657600080fd5b600154421115610222576040517fd02e774d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354341161026a576003546040517f4e12c1bb0000000000000000000000000000000000000000000000000000000081526004016102619190610566565b60405180910390fd5b6064605f34610279919061060c565b6102839190610695565b6004600082825461029491906106c6565b9250508190555060646005346102aa919061060c565b6102b49190610695565b6000808282546102c491906106c6565b9250508190555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550346003819055506102586001600082825461032791906106c6565b92505081905550600054343373ffffffffffffffffffffffffffffffffffffffff167fdafc4a123c6bb3b49dd38a0cba299808581a0126a37248a5f1102d5e5fa0633760405160405180910390a4565b6001544210156103b3576040517f44cee29000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda60405160405180910390a3600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc6004549081150290604051600060405180830381858888f19350505050158015610485573d6000803e3d6000fd5b5061048e6104e5565b565b6000600454905090565b6000600354905090565b60015481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60005481565b60008054905090565b610e10426104f391906106c6565b6001819055506000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006003819055506000600481905550565b6000819050919050565b6105608161054d565b82525050565b600060208201905061057b6000830184610557565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105ac82610581565b9050919050565b6105bc816105a1565b82525050565b60006020820190506105d760008301846105b3565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106178261054d565b91506106228361054d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561065b5761065a6105dd565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006106a08261054d565b91506106ab8361054d565b9250826106bb576106ba610666565b5b828204905092915050565b60006106d18261054d565b91506106dc8361054d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610711576107106105dd565b5b82820190509291505056fea2646970667358221220819fd82fafde91acbb496e0db366cb076b42b364ad9936cd570a124be7bcd0b564736f6c634300080e0033",
}

// BiddingpkABI is the input ABI used to generate the binding from.
// Deprecated: Use BiddingpkMetaData.ABI instead.
var BiddingpkABI = BiddingpkMetaData.ABI

// BiddingpkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BiddingpkMetaData.Bin instead.
var BiddingpkBin = BiddingpkMetaData.Bin

// DeployBiddingpk deploys a new Ethereum contract, binding an instance of Biddingpk to it.
func DeployBiddingpk(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Biddingpk, error) {
	parsed, err := BiddingpkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BiddingpkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Biddingpk{BiddingpkCaller: BiddingpkCaller{contract: contract}, BiddingpkTransactor: BiddingpkTransactor{contract: contract}, BiddingpkFilterer: BiddingpkFilterer{contract: contract}}, nil
}

// Biddingpk is an auto generated Go binding around an Ethereum contract.
type Biddingpk struct {
	BiddingpkCaller     // Read-only binding to the contract
	BiddingpkTransactor // Write-only binding to the contract
	BiddingpkFilterer   // Log filterer for contract events
}

// BiddingpkCaller is an auto generated read-only Go binding around an Ethereum contract.
type BiddingpkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingpkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BiddingpkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingpkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BiddingpkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BiddingpkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BiddingpkSession struct {
	Contract     *Biddingpk        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BiddingpkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BiddingpkCallerSession struct {
	Contract *BiddingpkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BiddingpkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BiddingpkTransactorSession struct {
	Contract     *BiddingpkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BiddingpkRaw is an auto generated low-level Go binding around an Ethereum contract.
type BiddingpkRaw struct {
	Contract *Biddingpk // Generic contract binding to access the raw methods on
}

// BiddingpkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BiddingpkCallerRaw struct {
	Contract *BiddingpkCaller // Generic read-only contract binding to access the raw methods on
}

// BiddingpkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BiddingpkTransactorRaw struct {
	Contract *BiddingpkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBiddingpk creates a new instance of Biddingpk, bound to a specific deployed contract.
func NewBiddingpk(address common.Address, backend bind.ContractBackend) (*Biddingpk, error) {
	contract, err := bindBiddingpk(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Biddingpk{BiddingpkCaller: BiddingpkCaller{contract: contract}, BiddingpkTransactor: BiddingpkTransactor{contract: contract}, BiddingpkFilterer: BiddingpkFilterer{contract: contract}}, nil
}

// NewBiddingpkCaller creates a new read-only instance of Biddingpk, bound to a specific deployed contract.
func NewBiddingpkCaller(address common.Address, caller bind.ContractCaller) (*BiddingpkCaller, error) {
	contract, err := bindBiddingpk(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingpkCaller{contract: contract}, nil
}

// NewBiddingpkTransactor creates a new write-only instance of Biddingpk, bound to a specific deployed contract.
func NewBiddingpkTransactor(address common.Address, transactor bind.ContractTransactor) (*BiddingpkTransactor, error) {
	contract, err := bindBiddingpk(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BiddingpkTransactor{contract: contract}, nil
}

// NewBiddingpkFilterer creates a new log filterer instance of Biddingpk, bound to a specific deployed contract.
func NewBiddingpkFilterer(address common.Address, filterer bind.ContractFilterer) (*BiddingpkFilterer, error) {
	contract, err := bindBiddingpk(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BiddingpkFilterer{contract: contract}, nil
}

// bindBiddingpk binds a generic wrapper to an already deployed contract.
func bindBiddingpk(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BiddingpkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingpk *BiddingpkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingpk.Contract.BiddingpkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingpk *BiddingpkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingpk.Contract.BiddingpkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingpk *BiddingpkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingpk.Contract.BiddingpkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Biddingpk *BiddingpkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Biddingpk.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Biddingpk *BiddingpkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingpk.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Biddingpk *BiddingpkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Biddingpk.Contract.contract.Transact(opts, method, params...)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) AuctionEndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "auctionEndTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_Biddingpk *BiddingpkSession) AuctionEndTime() (*big.Int, error) {
	return _Biddingpk.Contract.AuctionEndTime(&_Biddingpk.CallOpts)
}

// AuctionEndTime is a free data retrieval call binding the contract method 0x4b449cba.
//
// Solidity: function auctionEndTime() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) AuctionEndTime() (*big.Int, error) {
	return _Biddingpk.Contract.AuctionEndTime(&_Biddingpk.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) Commission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingpk *BiddingpkSession) Commission() (*big.Int, error) {
	return _Biddingpk.Contract.Commission(&_Biddingpk.CallOpts)
}

// Commission is a free data retrieval call binding the contract method 0xe1489191.
//
// Solidity: function commission() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) Commission() (*big.Int, error) {
	return _Biddingpk.Contract.Commission(&_Biddingpk.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0xe65d6b49.
//
// Solidity: function getCommission() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) GetCommission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "getCommission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommission is a free data retrieval call binding the contract method 0xe65d6b49.
//
// Solidity: function getCommission() view returns(uint256)
func (_Biddingpk *BiddingpkSession) GetCommission() (*big.Int, error) {
	return _Biddingpk.Contract.GetCommission(&_Biddingpk.CallOpts)
}

// GetCommission is a free data retrieval call binding the contract method 0xe65d6b49.
//
// Solidity: function getCommission() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) GetCommission() (*big.Int, error) {
	return _Biddingpk.Contract.GetCommission(&_Biddingpk.CallOpts)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) GetHighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "getHighestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(uint256)
func (_Biddingpk *BiddingpkSession) GetHighestBid() (*big.Int, error) {
	return _Biddingpk.Contract.GetHighestBid(&_Biddingpk.CallOpts)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) GetHighestBid() (*big.Int, error) {
	return _Biddingpk.Contract.GetHighestBid(&_Biddingpk.CallOpts)
}

// GettotalRewards is a free data retrieval call binding the contract method 0x2cba772e.
//
// Solidity: function gettotalRewards() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) GettotalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "gettotalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GettotalRewards is a free data retrieval call binding the contract method 0x2cba772e.
//
// Solidity: function gettotalRewards() view returns(uint256)
func (_Biddingpk *BiddingpkSession) GettotalRewards() (*big.Int, error) {
	return _Biddingpk.Contract.GettotalRewards(&_Biddingpk.CallOpts)
}

// GettotalRewards is a free data retrieval call binding the contract method 0x2cba772e.
//
// Solidity: function gettotalRewards() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) GettotalRewards() (*big.Int, error) {
	return _Biddingpk.Contract.GettotalRewards(&_Biddingpk.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingpk *BiddingpkCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingpk *BiddingpkSession) HighestBid() (*big.Int, error) {
	return _Biddingpk.Contract.HighestBid(&_Biddingpk.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Biddingpk *BiddingpkCallerSession) HighestBid() (*big.Int, error) {
	return _Biddingpk.Contract.HighestBid(&_Biddingpk.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingpk *BiddingpkCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Biddingpk.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingpk *BiddingpkSession) HighestBidder() (common.Address, error) {
	return _Biddingpk.Contract.HighestBidder(&_Biddingpk.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Biddingpk *BiddingpkCallerSession) HighestBidder() (common.Address, error) {
	return _Biddingpk.Contract.HighestBidder(&_Biddingpk.CallOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_Biddingpk *BiddingpkTransactor) AuctionEnd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingpk.contract.Transact(opts, "auctionEnd")
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_Biddingpk *BiddingpkSession) AuctionEnd() (*types.Transaction, error) {
	return _Biddingpk.Contract.AuctionEnd(&_Biddingpk.TransactOpts)
}

// AuctionEnd is a paid mutator transaction binding the contract method 0x2a24f46c.
//
// Solidity: function auctionEnd() returns()
func (_Biddingpk *BiddingpkTransactorSession) AuctionEnd() (*types.Transaction, error) {
	return _Biddingpk.Contract.AuctionEnd(&_Biddingpk.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingpk *BiddingpkTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Biddingpk.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingpk *BiddingpkSession) Bid() (*types.Transaction, error) {
	return _Biddingpk.Contract.Bid(&_Biddingpk.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Biddingpk *BiddingpkTransactorSession) Bid() (*types.Transaction, error) {
	return _Biddingpk.Contract.Bid(&_Biddingpk.TransactOpts)
}

// BiddingpkAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Biddingpk contract.
type BiddingpkAuctionEndedIterator struct {
	Event *BiddingpkAuctionEnded // Event containing the contract specifics and raw log

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
func (it *BiddingpkAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingpkAuctionEnded)
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
		it.Event = new(BiddingpkAuctionEnded)
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
func (it *BiddingpkAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingpkAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingpkAuctionEnded represents a AuctionEnded event raised by the Biddingpk contract.
type BiddingpkAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 indexed amount)
func (_Biddingpk *BiddingpkFilterer) FilterAuctionEnded(opts *bind.FilterOpts, winner []common.Address, amount []*big.Int) (*BiddingpkAuctionEndedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Biddingpk.contract.FilterLogs(opts, "AuctionEnded", winnerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &BiddingpkAuctionEndedIterator{contract: _Biddingpk.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 indexed amount)
func (_Biddingpk *BiddingpkFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *BiddingpkAuctionEnded, winner []common.Address, amount []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Biddingpk.contract.WatchLogs(opts, "AuctionEnded", winnerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingpkAuctionEnded)
				if err := _Biddingpk.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address indexed winner, uint256 indexed amount)
func (_Biddingpk *BiddingpkFilterer) ParseAuctionEnded(log types.Log) (*BiddingpkAuctionEnded, error) {
	event := new(BiddingpkAuctionEnded)
	if err := _Biddingpk.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BiddingpkHighestBidIncreasedIterator is returned from FilterHighestBidIncreased and is used to iterate over the raw logs and unpacked data for HighestBidIncreased events raised by the Biddingpk contract.
type BiddingpkHighestBidIncreasedIterator struct {
	Event *BiddingpkHighestBidIncreased // Event containing the contract specifics and raw log

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
func (it *BiddingpkHighestBidIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BiddingpkHighestBidIncreased)
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
		it.Event = new(BiddingpkHighestBidIncreased)
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
func (it *BiddingpkHighestBidIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BiddingpkHighestBidIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BiddingpkHighestBidIncreased represents a HighestBidIncreased event raised by the Biddingpk contract.
type BiddingpkHighestBidIncreased struct {
	Bidder     common.Address
	Amount     *big.Int
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHighestBidIncreased is a free log retrieval operation binding the contract event 0xdafc4a123c6bb3b49dd38a0cba299808581a0126a37248a5f1102d5e5fa06337.
//
// Solidity: event HighestBidIncreased(address indexed bidder, uint256 indexed amount, uint256 indexed commission)
func (_Biddingpk *BiddingpkFilterer) FilterHighestBidIncreased(opts *bind.FilterOpts, bidder []common.Address, amount []*big.Int, commission []*big.Int) (*BiddingpkHighestBidIncreasedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var commissionRule []interface{}
	for _, commissionItem := range commission {
		commissionRule = append(commissionRule, commissionItem)
	}

	logs, sub, err := _Biddingpk.contract.FilterLogs(opts, "HighestBidIncreased", bidderRule, amountRule, commissionRule)
	if err != nil {
		return nil, err
	}
	return &BiddingpkHighestBidIncreasedIterator{contract: _Biddingpk.contract, event: "HighestBidIncreased", logs: logs, sub: sub}, nil
}

// WatchHighestBidIncreased is a free log subscription operation binding the contract event 0xdafc4a123c6bb3b49dd38a0cba299808581a0126a37248a5f1102d5e5fa06337.
//
// Solidity: event HighestBidIncreased(address indexed bidder, uint256 indexed amount, uint256 indexed commission)
func (_Biddingpk *BiddingpkFilterer) WatchHighestBidIncreased(opts *bind.WatchOpts, sink chan<- *BiddingpkHighestBidIncreased, bidder []common.Address, amount []*big.Int, commission []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var commissionRule []interface{}
	for _, commissionItem := range commission {
		commissionRule = append(commissionRule, commissionItem)
	}

	logs, sub, err := _Biddingpk.contract.WatchLogs(opts, "HighestBidIncreased", bidderRule, amountRule, commissionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BiddingpkHighestBidIncreased)
				if err := _Biddingpk.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
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

// ParseHighestBidIncreased is a log parse operation binding the contract event 0xdafc4a123c6bb3b49dd38a0cba299808581a0126a37248a5f1102d5e5fa06337.
//
// Solidity: event HighestBidIncreased(address indexed bidder, uint256 indexed amount, uint256 indexed commission)
func (_Biddingpk *BiddingpkFilterer) ParseHighestBidIncreased(log types.Log) (*BiddingpkHighestBidIncreased, error) {
	event := new(BiddingpkHighestBidIncreased)
	if err := _Biddingpk.contract.UnpackLog(event, "HighestBidIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
