// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
// abigen version: 1.10.8-stable-26675454
// solidity v0.8.7+commit.e28d00a7

package rarity

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

// RarityMetaData contains all meta data concerning the Rarity contract.
var RarityMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"}],\"name\":\"leveled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"class\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"}],\"name\":\"summoned\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"adventure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"adventurers_log\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"classes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"level_up\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"next_summoner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_xp\",\"type\":\"uint256\"}],\"name\":\"spend_xp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"}],\"name\":\"summon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"summoner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_xp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_log\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"xp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curent_level\",\"type\":\"uint256\"}],\"name\":\"xp_required\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"xp_to_next_level\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// RarityABI is the input ABI used to generate the binding from.
// Deprecated: Use RarityMetaData.ABI instead.
var RarityABI = RarityMetaData.ABI

// Rarity is an auto generated Go binding around an Ethereum contract.
type Rarity struct {
	RarityCaller     // Read-only binding to the contract
	RarityTransactor // Write-only binding to the contract
	RarityFilterer   // Log filterer for contract events
}

// RarityCaller is an auto generated read-only Go binding around an Ethereum contract.
type RarityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RarityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RarityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaritySession struct {
	Contract     *Rarity           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RarityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RarityCallerSession struct {
	Contract *RarityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RarityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RarityTransactorSession struct {
	Contract     *RarityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RarityRaw is an auto generated low-level Go binding around an Ethereum contract.
type RarityRaw struct {
	Contract *Rarity // Generic contract binding to access the raw methods on
}

// RarityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RarityCallerRaw struct {
	Contract *RarityCaller // Generic read-only contract binding to access the raw methods on
}

// RarityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RarityTransactorRaw struct {
	Contract *RarityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRarity creates a new instance of Rarity, bound to a specific deployed contract.
func NewRarity(address common.Address, backend bind.ContractBackend) (*Rarity, error) {
	contract, err := bindRarity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rarity{RarityCaller: RarityCaller{contract: contract}, RarityTransactor: RarityTransactor{contract: contract}, RarityFilterer: RarityFilterer{contract: contract}}, nil
}

// NewRarityCaller creates a new read-only instance of Rarity, bound to a specific deployed contract.
func NewRarityCaller(address common.Address, caller bind.ContractCaller) (*RarityCaller, error) {
	contract, err := bindRarity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RarityCaller{contract: contract}, nil
}

// NewRarityTransactor creates a new write-only instance of Rarity, bound to a specific deployed contract.
func NewRarityTransactor(address common.Address, transactor bind.ContractTransactor) (*RarityTransactor, error) {
	contract, err := bindRarity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RarityTransactor{contract: contract}, nil
}

// NewRarityFilterer creates a new log filterer instance of Rarity, bound to a specific deployed contract.
func NewRarityFilterer(address common.Address, filterer bind.ContractFilterer) (*RarityFilterer, error) {
	contract, err := bindRarity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RarityFilterer{contract: contract}, nil
}

// bindRarity binds a generic wrapper to an already deployed contract.
func bindRarity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RarityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rarity *RarityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rarity.Contract.RarityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rarity *RarityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rarity.Contract.RarityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rarity *RarityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rarity.Contract.RarityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rarity *RarityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rarity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rarity *RarityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rarity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rarity *RarityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rarity.Contract.contract.Transact(opts, method, params...)
}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_Rarity *RarityCaller) AdventurersLog(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "adventurers_log", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_Rarity *RaritySession) AdventurersLog(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.AdventurersLog(&_Rarity.CallOpts, arg0)
}

// AdventurersLog is a free data retrieval call binding the contract method 0xeed25028.
//
// Solidity: function adventurers_log(uint256 ) view returns(uint256)
func (_Rarity *RarityCallerSession) AdventurersLog(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.AdventurersLog(&_Rarity.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Rarity *RarityCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Rarity *RaritySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Rarity.Contract.BalanceOf(&_Rarity.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Rarity *RarityCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Rarity.Contract.BalanceOf(&_Rarity.CallOpts, owner)
}

// Class is a free data retrieval call binding the contract method 0x3613a9f4.
//
// Solidity: function class(uint256 ) view returns(uint256)
func (_Rarity *RarityCaller) Class(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "class", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Class is a free data retrieval call binding the contract method 0x3613a9f4.
//
// Solidity: function class(uint256 ) view returns(uint256)
func (_Rarity *RaritySession) Class(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Class(&_Rarity.CallOpts, arg0)
}

// Class is a free data retrieval call binding the contract method 0x3613a9f4.
//
// Solidity: function class(uint256 ) view returns(uint256)
func (_Rarity *RarityCallerSession) Class(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Class(&_Rarity.CallOpts, arg0)
}

// Classes is a free data retrieval call binding the contract method 0x817dbe9f.
//
// Solidity: function classes(uint256 id) pure returns(string description)
func (_Rarity *RarityCaller) Classes(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "classes", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Classes is a free data retrieval call binding the contract method 0x817dbe9f.
//
// Solidity: function classes(uint256 id) pure returns(string description)
func (_Rarity *RaritySession) Classes(id *big.Int) (string, error) {
	return _Rarity.Contract.Classes(&_Rarity.CallOpts, id)
}

// Classes is a free data retrieval call binding the contract method 0x817dbe9f.
//
// Solidity: function classes(uint256 id) pure returns(string description)
func (_Rarity *RarityCallerSession) Classes(id *big.Int) (string, error) {
	return _Rarity.Contract.Classes(&_Rarity.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Rarity *RarityCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Rarity *RaritySession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Rarity.Contract.GetApproved(&_Rarity.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Rarity *RarityCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Rarity.Contract.GetApproved(&_Rarity.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Rarity *RarityCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Rarity *RaritySession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Rarity.Contract.IsApprovedForAll(&_Rarity.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Rarity *RarityCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Rarity.Contract.IsApprovedForAll(&_Rarity.CallOpts, owner, operator)
}

// Level is a free data retrieval call binding the contract method 0x05c58df2.
//
// Solidity: function level(uint256 ) view returns(uint256)
func (_Rarity *RarityCaller) Level(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "level", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Level is a free data retrieval call binding the contract method 0x05c58df2.
//
// Solidity: function level(uint256 ) view returns(uint256)
func (_Rarity *RaritySession) Level(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Level(&_Rarity.CallOpts, arg0)
}

// Level is a free data retrieval call binding the contract method 0x05c58df2.
//
// Solidity: function level(uint256 ) view returns(uint256)
func (_Rarity *RarityCallerSession) Level(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Level(&_Rarity.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Rarity *RarityCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Rarity *RaritySession) Name() (string, error) {
	return _Rarity.Contract.Name(&_Rarity.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Rarity *RarityCallerSession) Name() (string, error) {
	return _Rarity.Contract.Name(&_Rarity.CallOpts)
}

// NextSummoner is a free data retrieval call binding the contract method 0x07b018ef.
//
// Solidity: function next_summoner() view returns(uint256)
func (_Rarity *RarityCaller) NextSummoner(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "next_summoner")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSummoner is a free data retrieval call binding the contract method 0x07b018ef.
//
// Solidity: function next_summoner() view returns(uint256)
func (_Rarity *RaritySession) NextSummoner() (*big.Int, error) {
	return _Rarity.Contract.NextSummoner(&_Rarity.CallOpts)
}

// NextSummoner is a free data retrieval call binding the contract method 0x07b018ef.
//
// Solidity: function next_summoner() view returns(uint256)
func (_Rarity *RarityCallerSession) NextSummoner() (*big.Int, error) {
	return _Rarity.Contract.NextSummoner(&_Rarity.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Rarity *RarityCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Rarity *RaritySession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Rarity.Contract.OwnerOf(&_Rarity.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Rarity *RarityCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Rarity.Contract.OwnerOf(&_Rarity.CallOpts, tokenId)
}

// Summoner is a free data retrieval call binding the contract method 0xc363b9eb.
//
// Solidity: function summoner(uint256 _summoner) view returns(uint256 _xp, uint256 _log, uint256 _class, uint256 _level)
func (_Rarity *RarityCaller) Summoner(opts *bind.CallOpts, _summoner *big.Int) (struct {
	Xp    *big.Int
	Log   *big.Int
	Class *big.Int
	Level *big.Int
}, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "summoner", _summoner)

	outstruct := new(struct {
		Xp    *big.Int
		Log   *big.Int
		Class *big.Int
		Level *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Xp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Log = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Class = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Level = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Summoner is a free data retrieval call binding the contract method 0xc363b9eb.
//
// Solidity: function summoner(uint256 _summoner) view returns(uint256 _xp, uint256 _log, uint256 _class, uint256 _level)
func (_Rarity *RaritySession) Summoner(_summoner *big.Int) (struct {
	Xp    *big.Int
	Log   *big.Int
	Class *big.Int
	Level *big.Int
}, error) {
	return _Rarity.Contract.Summoner(&_Rarity.CallOpts, _summoner)
}

// Summoner is a free data retrieval call binding the contract method 0xc363b9eb.
//
// Solidity: function summoner(uint256 _summoner) view returns(uint256 _xp, uint256 _log, uint256 _class, uint256 _level)
func (_Rarity *RarityCallerSession) Summoner(_summoner *big.Int) (struct {
	Xp    *big.Int
	Log   *big.Int
	Class *big.Int
	Level *big.Int
}, error) {
	return _Rarity.Contract.Summoner(&_Rarity.CallOpts, _summoner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Rarity *RarityCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Rarity *RaritySession) Symbol() (string, error) {
	return _Rarity.Contract.Symbol(&_Rarity.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Rarity *RarityCallerSession) Symbol() (string, error) {
	return _Rarity.Contract.Symbol(&_Rarity.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Rarity *RarityCaller) TokenURI(opts *bind.CallOpts, _summoner *big.Int) (string, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "tokenURI", _summoner)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Rarity *RaritySession) TokenURI(_summoner *big.Int) (string, error) {
	return _Rarity.Contract.TokenURI(&_Rarity.CallOpts, _summoner)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_Rarity *RarityCallerSession) TokenURI(_summoner *big.Int) (string, error) {
	return _Rarity.Contract.TokenURI(&_Rarity.CallOpts, _summoner)
}

// Xp is a free data retrieval call binding the contract method 0x94b69ffa.
//
// Solidity: function xp(uint256 ) view returns(uint256)
func (_Rarity *RarityCaller) Xp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "xp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Xp is a free data retrieval call binding the contract method 0x94b69ffa.
//
// Solidity: function xp(uint256 ) view returns(uint256)
func (_Rarity *RaritySession) Xp(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Xp(&_Rarity.CallOpts, arg0)
}

// Xp is a free data retrieval call binding the contract method 0x94b69ffa.
//
// Solidity: function xp(uint256 ) view returns(uint256)
func (_Rarity *RarityCallerSession) Xp(arg0 *big.Int) (*big.Int, error) {
	return _Rarity.Contract.Xp(&_Rarity.CallOpts, arg0)
}

// XpRequired is a free data retrieval call binding the contract method 0xaca10be3.
//
// Solidity: function xp_required(uint256 curent_level) pure returns(uint256 xp_to_next_level)
func (_Rarity *RarityCaller) XpRequired(opts *bind.CallOpts, curent_level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rarity.contract.Call(opts, &out, "xp_required", curent_level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XpRequired is a free data retrieval call binding the contract method 0xaca10be3.
//
// Solidity: function xp_required(uint256 curent_level) pure returns(uint256 xp_to_next_level)
func (_Rarity *RaritySession) XpRequired(curent_level *big.Int) (*big.Int, error) {
	return _Rarity.Contract.XpRequired(&_Rarity.CallOpts, curent_level)
}

// XpRequired is a free data retrieval call binding the contract method 0xaca10be3.
//
// Solidity: function xp_required(uint256 curent_level) pure returns(uint256 xp_to_next_level)
func (_Rarity *RarityCallerSession) XpRequired(curent_level *big.Int) (*big.Int, error) {
	return _Rarity.Contract.XpRequired(&_Rarity.CallOpts, curent_level)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns()
func (_Rarity *RarityTransactor) Adventure(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "adventure", _summoner)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns()
func (_Rarity *RaritySession) Adventure(_summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Adventure(&_Rarity.TransactOpts, _summoner)
}

// Adventure is a paid mutator transaction binding the contract method 0xb00b52f1.
//
// Solidity: function adventure(uint256 _summoner) returns()
func (_Rarity *RarityTransactorSession) Adventure(_summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Adventure(&_Rarity.TransactOpts, _summoner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Rarity *RaritySession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Approve(&_Rarity.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Approve(&_Rarity.TransactOpts, to, tokenId)
}

// LevelUp is a paid mutator transaction binding the contract method 0x90249448.
//
// Solidity: function level_up(uint256 _summoner) returns()
func (_Rarity *RarityTransactor) LevelUp(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "level_up", _summoner)
}

// LevelUp is a paid mutator transaction binding the contract method 0x90249448.
//
// Solidity: function level_up(uint256 _summoner) returns()
func (_Rarity *RaritySession) LevelUp(_summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.LevelUp(&_Rarity.TransactOpts, _summoner)
}

// LevelUp is a paid mutator transaction binding the contract method 0x90249448.
//
// Solidity: function level_up(uint256 _summoner) returns()
func (_Rarity *RarityTransactorSession) LevelUp(_summoner *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.LevelUp(&_Rarity.TransactOpts, _summoner)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RaritySession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.SafeTransferFrom(&_Rarity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.SafeTransferFrom(&_Rarity.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Rarity *RarityTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Rarity *RaritySession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Rarity.Contract.SafeTransferFrom0(&_Rarity.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Rarity *RarityTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Rarity.Contract.SafeTransferFrom0(&_Rarity.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Rarity *RarityTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Rarity *RaritySession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Rarity.Contract.SetApprovalForAll(&_Rarity.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Rarity *RarityTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Rarity.Contract.SetApprovalForAll(&_Rarity.TransactOpts, operator, approved)
}

// SpendXp is a paid mutator transaction binding the contract method 0xe58410bb.
//
// Solidity: function spend_xp(uint256 _summoner, uint256 _xp) returns()
func (_Rarity *RarityTransactor) SpendXp(opts *bind.TransactOpts, _summoner *big.Int, _xp *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "spend_xp", _summoner, _xp)
}

// SpendXp is a paid mutator transaction binding the contract method 0xe58410bb.
//
// Solidity: function spend_xp(uint256 _summoner, uint256 _xp) returns()
func (_Rarity *RaritySession) SpendXp(_summoner *big.Int, _xp *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.SpendXp(&_Rarity.TransactOpts, _summoner, _xp)
}

// SpendXp is a paid mutator transaction binding the contract method 0xe58410bb.
//
// Solidity: function spend_xp(uint256 _summoner, uint256 _xp) returns()
func (_Rarity *RarityTransactorSession) SpendXp(_summoner *big.Int, _xp *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.SpendXp(&_Rarity.TransactOpts, _summoner, _xp)
}

// Summon is a paid mutator transaction binding the contract method 0x035d9f2a.
//
// Solidity: function summon(uint256 _class) returns()
func (_Rarity *RarityTransactor) Summon(opts *bind.TransactOpts, _class *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "summon", _class)
}

// Summon is a paid mutator transaction binding the contract method 0x035d9f2a.
//
// Solidity: function summon(uint256 _class) returns()
func (_Rarity *RaritySession) Summon(_class *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Summon(&_Rarity.TransactOpts, _class)
}

// Summon is a paid mutator transaction binding the contract method 0x035d9f2a.
//
// Solidity: function summon(uint256 _class) returns()
func (_Rarity *RarityTransactorSession) Summon(_class *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.Summon(&_Rarity.TransactOpts, _class)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RaritySession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.TransferFrom(&_Rarity.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Rarity *RarityTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Rarity.Contract.TransferFrom(&_Rarity.TransactOpts, from, to, tokenId)
}

// RarityApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Rarity contract.
type RarityApprovalIterator struct {
	Event *RarityApproval // Event containing the contract specifics and raw log

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
func (it *RarityApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityApproval)
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
		it.Event = new(RarityApproval)
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
func (it *RarityApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityApproval represents a Approval event raised by the Rarity contract.
type RarityApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RarityApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Rarity.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RarityApprovalIterator{contract: _Rarity.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RarityApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Rarity.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityApproval)
				if err := _Rarity.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) ParseApproval(log types.Log) (*RarityApproval, error) {
	event := new(RarityApproval)
	if err := _Rarity.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarityApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Rarity contract.
type RarityApprovalForAllIterator struct {
	Event *RarityApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RarityApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityApprovalForAll)
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
		it.Event = new(RarityApprovalForAll)
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
func (it *RarityApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityApprovalForAll represents a ApprovalForAll event raised by the Rarity contract.
type RarityApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Rarity *RarityFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RarityApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Rarity.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RarityApprovalForAllIterator{contract: _Rarity.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Rarity *RarityFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RarityApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Rarity.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityApprovalForAll)
				if err := _Rarity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Rarity *RarityFilterer) ParseApprovalForAll(log types.Log) (*RarityApprovalForAll, error) {
	event := new(RarityApprovalForAll)
	if err := _Rarity.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarityTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Rarity contract.
type RarityTransferIterator struct {
	Event *RarityTransfer // Event containing the contract specifics and raw log

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
func (it *RarityTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityTransfer)
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
		it.Event = new(RarityTransfer)
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
func (it *RarityTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityTransfer represents a Transfer event raised by the Rarity contract.
type RarityTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RarityTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Rarity.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RarityTransferIterator{contract: _Rarity.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RarityTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Rarity.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityTransfer)
				if err := _Rarity.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Rarity *RarityFilterer) ParseTransfer(log types.Log) (*RarityTransfer, error) {
	event := new(RarityTransfer)
	if err := _Rarity.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarityLeveledIterator is returned from FilterLeveled and is used to iterate over the raw logs and unpacked data for Leveled events raised by the Rarity contract.
type RarityLeveledIterator struct {
	Event *RarityLeveled // Event containing the contract specifics and raw log

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
func (it *RarityLeveledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityLeveled)
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
		it.Event = new(RarityLeveled)
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
func (it *RarityLeveledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityLeveledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityLeveled represents a Leveled event raised by the Rarity contract.
type RarityLeveled struct {
	Owner    common.Address
	Level    *big.Int
	Summoner *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLeveled is a free log retrieval operation binding the contract event 0x943b00a67689342d89989c0bf7b5f4840b8c5a29e1383dc391a903626b3327c0.
//
// Solidity: event leveled(address indexed owner, uint256 level, uint256 summoner)
func (_Rarity *RarityFilterer) FilterLeveled(opts *bind.FilterOpts, owner []common.Address) (*RarityLeveledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rarity.contract.FilterLogs(opts, "leveled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &RarityLeveledIterator{contract: _Rarity.contract, event: "leveled", logs: logs, sub: sub}, nil
}

// WatchLeveled is a free log subscription operation binding the contract event 0x943b00a67689342d89989c0bf7b5f4840b8c5a29e1383dc391a903626b3327c0.
//
// Solidity: event leveled(address indexed owner, uint256 level, uint256 summoner)
func (_Rarity *RarityFilterer) WatchLeveled(opts *bind.WatchOpts, sink chan<- *RarityLeveled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rarity.contract.WatchLogs(opts, "leveled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityLeveled)
				if err := _Rarity.contract.UnpackLog(event, "leveled", log); err != nil {
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

// ParseLeveled is a log parse operation binding the contract event 0x943b00a67689342d89989c0bf7b5f4840b8c5a29e1383dc391a903626b3327c0.
//
// Solidity: event leveled(address indexed owner, uint256 level, uint256 summoner)
func (_Rarity *RarityFilterer) ParseLeveled(log types.Log) (*RarityLeveled, error) {
	event := new(RarityLeveled)
	if err := _Rarity.contract.UnpackLog(event, "leveled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaritySummonedIterator is returned from FilterSummoned and is used to iterate over the raw logs and unpacked data for Summoned events raised by the Rarity contract.
type RaritySummonedIterator struct {
	Event *RaritySummoned // Event containing the contract specifics and raw log

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
func (it *RaritySummonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaritySummoned)
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
		it.Event = new(RaritySummoned)
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
func (it *RaritySummonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaritySummonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaritySummoned represents a Summoned event raised by the Rarity contract.
type RaritySummoned struct {
	Owner    common.Address
	Class    *big.Int
	Summoner *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSummoned is a free log retrieval operation binding the contract event 0x13635f4880051027279fad8cb34c7f6b430bee464ad05f777d9abff668fc1be8.
//
// Solidity: event summoned(address indexed owner, uint256 class, uint256 summoner)
func (_Rarity *RarityFilterer) FilterSummoned(opts *bind.FilterOpts, owner []common.Address) (*RaritySummonedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rarity.contract.FilterLogs(opts, "summoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return &RaritySummonedIterator{contract: _Rarity.contract, event: "summoned", logs: logs, sub: sub}, nil
}

// WatchSummoned is a free log subscription operation binding the contract event 0x13635f4880051027279fad8cb34c7f6b430bee464ad05f777d9abff668fc1be8.
//
// Solidity: event summoned(address indexed owner, uint256 class, uint256 summoner)
func (_Rarity *RarityFilterer) WatchSummoned(opts *bind.WatchOpts, sink chan<- *RaritySummoned, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Rarity.contract.WatchLogs(opts, "summoned", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaritySummoned)
				if err := _Rarity.contract.UnpackLog(event, "summoned", log); err != nil {
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

// ParseSummoned is a log parse operation binding the contract event 0x13635f4880051027279fad8cb34c7f6b430bee464ad05f777d9abff668fc1be8.
//
// Solidity: event summoned(address indexed owner, uint256 class, uint256 summoner)
func (_Rarity *RarityFilterer) ParseSummoned(log types.Log) (*RaritySummoned, error) {
	event := new(RaritySummoned)
	if err := _Rarity.contract.UnpackLog(event, "summoned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
