// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

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

// RarityAttributesMetaData contains all meta data concerning the RarityAttributes contract.
var RarityAttributesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"leveler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"summoner\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"name\":\"Leveled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current_level\",\"type\":\"uint256\"}],\"name\":\"abilities_by_level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ability_scores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"strength\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"dexterity\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"constitution\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"intelligence\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"wisdom\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"charisma\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"calc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_str\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_const\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_int\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_wis\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cha\",\"type\":\"uint256\"}],\"name\":\"calculate_point_buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"character_created\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_charisma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_constitution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_dexterity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_intelligence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_strength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"increase_wisdom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"level_points_spent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_str\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_dex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_const\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_int\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_wis\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_cha\",\"type\":\"uint32\"}],\"name\":\"point_buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summoner\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RarityAttributesABI is the input ABI used to generate the binding from.
// Deprecated: Use RarityAttributesMetaData.ABI instead.
var RarityAttributesABI = RarityAttributesMetaData.ABI

// RarityAttributes is an auto generated Go binding around an Ethereum contract.
type RarityAttributes struct {
	RarityAttributesCaller     // Read-only binding to the contract
	RarityAttributesTransactor // Write-only binding to the contract
	RarityAttributesFilterer   // Log filterer for contract events
}

// RarityAttributesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RarityAttributesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarityAttributesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RarityAttributesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarityAttributesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RarityAttributesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarityAttributesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RarityAttributesSession struct {
	Contract     *RarityAttributes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RarityAttributesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RarityAttributesCallerSession struct {
	Contract *RarityAttributesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RarityAttributesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RarityAttributesTransactorSession struct {
	Contract     *RarityAttributesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RarityAttributesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RarityAttributesRaw struct {
	Contract *RarityAttributes // Generic contract binding to access the raw methods on
}

// RarityAttributesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RarityAttributesCallerRaw struct {
	Contract *RarityAttributesCaller // Generic read-only contract binding to access the raw methods on
}

// RarityAttributesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RarityAttributesTransactorRaw struct {
	Contract *RarityAttributesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRarityAttributes creates a new instance of RarityAttributes, bound to a specific deployed contract.
func NewRarityAttributes(address common.Address, backend bind.ContractBackend) (*RarityAttributes, error) {
	contract, err := bindRarityAttributes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RarityAttributes{RarityAttributesCaller: RarityAttributesCaller{contract: contract}, RarityAttributesTransactor: RarityAttributesTransactor{contract: contract}, RarityAttributesFilterer: RarityAttributesFilterer{contract: contract}}, nil
}

// NewRarityAttributesCaller creates a new read-only instance of RarityAttributes, bound to a specific deployed contract.
func NewRarityAttributesCaller(address common.Address, caller bind.ContractCaller) (*RarityAttributesCaller, error) {
	contract, err := bindRarityAttributes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RarityAttributesCaller{contract: contract}, nil
}

// NewRarityAttributesTransactor creates a new write-only instance of RarityAttributes, bound to a specific deployed contract.
func NewRarityAttributesTransactor(address common.Address, transactor bind.ContractTransactor) (*RarityAttributesTransactor, error) {
	contract, err := bindRarityAttributes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RarityAttributesTransactor{contract: contract}, nil
}

// NewRarityAttributesFilterer creates a new log filterer instance of RarityAttributes, bound to a specific deployed contract.
func NewRarityAttributesFilterer(address common.Address, filterer bind.ContractFilterer) (*RarityAttributesFilterer, error) {
	contract, err := bindRarityAttributes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RarityAttributesFilterer{contract: contract}, nil
}

// bindRarityAttributes binds a generic wrapper to an already deployed contract.
func bindRarityAttributes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RarityAttributesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RarityAttributes *RarityAttributesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RarityAttributes.Contract.RarityAttributesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RarityAttributes *RarityAttributesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RarityAttributes.Contract.RarityAttributesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RarityAttributes *RarityAttributesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RarityAttributes.Contract.RarityAttributesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RarityAttributes *RarityAttributesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RarityAttributes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RarityAttributes *RarityAttributesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RarityAttributes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RarityAttributes *RarityAttributesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RarityAttributes.Contract.contract.Transact(opts, method, params...)
}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCaller) AbilitiesByLevel(opts *bind.CallOpts, current_level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "abilities_by_level", current_level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_RarityAttributes *RarityAttributesSession) AbilitiesByLevel(current_level *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.AbilitiesByLevel(&_RarityAttributes.CallOpts, current_level)
}

// AbilitiesByLevel is a free data retrieval call binding the contract method 0x51c23226.
//
// Solidity: function abilities_by_level(uint256 current_level) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCallerSession) AbilitiesByLevel(current_level *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.AbilitiesByLevel(&_RarityAttributes.CallOpts, current_level)
}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesCaller) AbilityScores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "ability_scores", arg0)

	outstruct := new(struct {
		Strength     uint32
		Dexterity    uint32
		Constitution uint32
		Intelligence uint32
		Wisdom       uint32
		Charisma     uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strength = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Dexterity = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Constitution = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Intelligence = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Wisdom = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Charisma = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesSession) AbilityScores(arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	return _RarityAttributes.Contract.AbilityScores(&_RarityAttributes.CallOpts, arg0)
}

// AbilityScores is a free data retrieval call binding the contract method 0x77d9e11c.
//
// Solidity: function ability_scores(uint256 ) view returns(uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesCallerSession) AbilityScores(arg0 *big.Int) (struct {
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
}, error) {
	return _RarityAttributes.Contract.AbilityScores(&_RarityAttributes.CallOpts, arg0)
}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCaller) Calc(opts *bind.CallOpts, score *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "calc", score)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_RarityAttributes *RarityAttributesSession) Calc(score *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.Calc(&_RarityAttributes.CallOpts, score)
}

// Calc is a free data retrieval call binding the contract method 0x38c9027a.
//
// Solidity: function calc(uint256 score) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCallerSession) Calc(score *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.Calc(&_RarityAttributes.CallOpts, score)
}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCaller) CalculatePointBuy(opts *bind.CallOpts, _str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "calculate_point_buy", _str, _dex, _const, _int, _wis, _cha)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_RarityAttributes *RarityAttributesSession) CalculatePointBuy(_str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.CalculatePointBuy(&_RarityAttributes.CallOpts, _str, _dex, _const, _int, _wis, _cha)
}

// CalculatePointBuy is a free data retrieval call binding the contract method 0x9ed24c8e.
//
// Solidity: function calculate_point_buy(uint256 _str, uint256 _dex, uint256 _const, uint256 _int, uint256 _wis, uint256 _cha) pure returns(uint256)
func (_RarityAttributes *RarityAttributesCallerSession) CalculatePointBuy(_str *big.Int, _dex *big.Int, _const *big.Int, _int *big.Int, _wis *big.Int, _cha *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.CalculatePointBuy(&_RarityAttributes.CallOpts, _str, _dex, _const, _int, _wis, _cha)
}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_RarityAttributes *RarityAttributesCaller) CharacterCreated(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "character_created", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_RarityAttributes *RarityAttributesSession) CharacterCreated(arg0 *big.Int) (bool, error) {
	return _RarityAttributes.Contract.CharacterCreated(&_RarityAttributes.CallOpts, arg0)
}

// CharacterCreated is a free data retrieval call binding the contract method 0xf8955432.
//
// Solidity: function character_created(uint256 ) view returns(bool)
func (_RarityAttributes *RarityAttributesCallerSession) CharacterCreated(arg0 *big.Int) (bool, error) {
	return _RarityAttributes.Contract.CharacterCreated(&_RarityAttributes.CallOpts, arg0)
}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_RarityAttributes *RarityAttributesCaller) LevelPointsSpent(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "level_points_spent", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_RarityAttributes *RarityAttributesSession) LevelPointsSpent(arg0 *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.LevelPointsSpent(&_RarityAttributes.CallOpts, arg0)
}

// LevelPointsSpent is a free data retrieval call binding the contract method 0x849ada74.
//
// Solidity: function level_points_spent(uint256 ) view returns(uint256)
func (_RarityAttributes *RarityAttributesCallerSession) LevelPointsSpent(arg0 *big.Int) (*big.Int, error) {
	return _RarityAttributes.Contract.LevelPointsSpent(&_RarityAttributes.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_RarityAttributes *RarityAttributesCaller) TokenURI(opts *bind.CallOpts, _summoner *big.Int) (string, error) {
	var out []interface{}
	err := _RarityAttributes.contract.Call(opts, &out, "tokenURI", _summoner)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_RarityAttributes *RarityAttributesSession) TokenURI(_summoner *big.Int) (string, error) {
	return _RarityAttributes.Contract.TokenURI(&_RarityAttributes.CallOpts, _summoner)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _summoner) view returns(string)
func (_RarityAttributes *RarityAttributesCallerSession) TokenURI(_summoner *big.Int) (string, error) {
	return _RarityAttributes.Contract.TokenURI(&_RarityAttributes.CallOpts, _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseCharisma(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_charisma", _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseCharisma(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseCharisma(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseCharisma is a paid mutator transaction binding the contract method 0xbf2bf895.
//
// Solidity: function increase_charisma(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseCharisma(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseCharisma(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseConstitution(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_constitution", _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseConstitution(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseConstitution(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseConstitution is a paid mutator transaction binding the contract method 0xe0d92c4a.
//
// Solidity: function increase_constitution(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseConstitution(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseConstitution(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseDexterity(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_dexterity", _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseDexterity(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseDexterity(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseDexterity is a paid mutator transaction binding the contract method 0x05934d9c.
//
// Solidity: function increase_dexterity(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseDexterity(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseDexterity(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseIntelligence(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_intelligence", _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseIntelligence(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseIntelligence(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseIntelligence is a paid mutator transaction binding the contract method 0x96cf4c4b.
//
// Solidity: function increase_intelligence(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseIntelligence(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseIntelligence(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseStrength(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_strength", _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseStrength(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseStrength(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseStrength is a paid mutator transaction binding the contract method 0xde999039.
//
// Solidity: function increase_strength(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseStrength(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseStrength(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactor) IncreaseWisdom(opts *bind.TransactOpts, _summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "increase_wisdom", _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesSession) IncreaseWisdom(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseWisdom(&_RarityAttributes.TransactOpts, _summoner)
}

// IncreaseWisdom is a paid mutator transaction binding the contract method 0xfe6676b3.
//
// Solidity: function increase_wisdom(uint256 _summoner) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) IncreaseWisdom(_summoner *big.Int) (*types.Transaction, error) {
	return _RarityAttributes.Contract.IncreaseWisdom(&_RarityAttributes.TransactOpts, _summoner)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_RarityAttributes *RarityAttributesTransactor) PointBuy(opts *bind.TransactOpts, _summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _RarityAttributes.contract.Transact(opts, "point_buy", _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_RarityAttributes *RarityAttributesSession) PointBuy(_summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _RarityAttributes.Contract.PointBuy(&_RarityAttributes.TransactOpts, _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// PointBuy is a paid mutator transaction binding the contract method 0xc3c2407c.
//
// Solidity: function point_buy(uint256 _summoner, uint32 _str, uint32 _dex, uint32 _const, uint32 _int, uint32 _wis, uint32 _cha) returns()
func (_RarityAttributes *RarityAttributesTransactorSession) PointBuy(_summoner *big.Int, _str uint32, _dex uint32, _const uint32, _int uint32, _wis uint32, _cha uint32) (*types.Transaction, error) {
	return _RarityAttributes.Contract.PointBuy(&_RarityAttributes.TransactOpts, _summoner, _str, _dex, _const, _int, _wis, _cha)
}

// RarityAttributesCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the RarityAttributes contract.
type RarityAttributesCreatedIterator struct {
	Event *RarityAttributesCreated // Event containing the contract specifics and raw log

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
func (it *RarityAttributesCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityAttributesCreated)
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
		it.Event = new(RarityAttributesCreated)
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
func (it *RarityAttributesCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityAttributesCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityAttributesCreated represents a Created event raised by the RarityAttributes contract.
type RarityAttributesCreated struct {
	Creator      common.Address
	Summoner     *big.Int
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) FilterCreated(opts *bind.FilterOpts, creator []common.Address) (*RarityAttributesCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _RarityAttributes.contract.FilterLogs(opts, "Created", creatorRule)
	if err != nil {
		return nil, err
	}
	return &RarityAttributesCreatedIterator{contract: _RarityAttributes.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *RarityAttributesCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _RarityAttributes.contract.WatchLogs(opts, "Created", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityAttributesCreated)
				if err := _RarityAttributes.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x471b798bfd4e4756e197917dd3ba2cbf782f09c2deef48cd5fa694cdba976328.
//
// Solidity: event Created(address indexed creator, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) ParseCreated(log types.Log) (*RarityAttributesCreated, error) {
	event := new(RarityAttributesCreated)
	if err := _RarityAttributes.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarityAttributesLeveledIterator is returned from FilterLeveled and is used to iterate over the raw logs and unpacked data for Leveled events raised by the RarityAttributes contract.
type RarityAttributesLeveledIterator struct {
	Event *RarityAttributesLeveled // Event containing the contract specifics and raw log

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
func (it *RarityAttributesLeveledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarityAttributesLeveled)
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
		it.Event = new(RarityAttributesLeveled)
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
func (it *RarityAttributesLeveledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarityAttributesLeveledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarityAttributesLeveled represents a Leveled event raised by the RarityAttributes contract.
type RarityAttributesLeveled struct {
	Leveler      common.Address
	Summoner     *big.Int
	Strength     uint32
	Dexterity    uint32
	Constitution uint32
	Intelligence uint32
	Wisdom       uint32
	Charisma     uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLeveled is a free log retrieval operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) FilterLeveled(opts *bind.FilterOpts, leveler []common.Address) (*RarityAttributesLeveledIterator, error) {

	var levelerRule []interface{}
	for _, levelerItem := range leveler {
		levelerRule = append(levelerRule, levelerItem)
	}

	logs, sub, err := _RarityAttributes.contract.FilterLogs(opts, "Leveled", levelerRule)
	if err != nil {
		return nil, err
	}
	return &RarityAttributesLeveledIterator{contract: _RarityAttributes.contract, event: "Leveled", logs: logs, sub: sub}, nil
}

// WatchLeveled is a free log subscription operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) WatchLeveled(opts *bind.WatchOpts, sink chan<- *RarityAttributesLeveled, leveler []common.Address) (event.Subscription, error) {

	var levelerRule []interface{}
	for _, levelerItem := range leveler {
		levelerRule = append(levelerRule, levelerItem)
	}

	logs, sub, err := _RarityAttributes.contract.WatchLogs(opts, "Leveled", levelerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarityAttributesLeveled)
				if err := _RarityAttributes.contract.UnpackLog(event, "Leveled", log); err != nil {
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

// ParseLeveled is a log parse operation binding the contract event 0x207b63f0a8632134b76c8dd10b38a52179fc29b3fa01eb51e4662e4c720ec8d4.
//
// Solidity: event Leveled(address indexed leveler, uint256 summoner, uint32 strength, uint32 dexterity, uint32 constitution, uint32 intelligence, uint32 wisdom, uint32 charisma)
func (_RarityAttributes *RarityAttributesFilterer) ParseLeveled(log types.Log) (*RarityAttributesLeveled, error) {
	event := new(RarityAttributesLeveled)
	if err := _RarityAttributes.contract.UnpackLog(event, "Leveled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
