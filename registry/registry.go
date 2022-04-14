// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"documents\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vc\",\"type\":\"bytes32\"}],\"name\":\"exchangeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vc\",\"type\":\"bytes32\"}],\"name\":\"renewVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vc\",\"type\":\"bytes32\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vc\",\"type\":\"bytes32\"}],\"name\":\"uploadVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Documents is a free data retrieval call binding the contract method 0xb3dd697b.
//
// Solidity: function documents(string ) view returns(bytes32)
func (_Registry *RegistryCaller) Documents(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "documents", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Documents is a free data retrieval call binding the contract method 0xb3dd697b.
//
// Solidity: function documents(string ) view returns(bytes32)
func (_Registry *RegistrySession) Documents(arg0 string) ([32]byte, error) {
	return _Registry.Contract.Documents(&_Registry.CallOpts, arg0)
}

// Documents is a free data retrieval call binding the contract method 0xb3dd697b.
//
// Solidity: function documents(string ) view returns(bytes32)
func (_Registry *RegistryCallerSession) Documents(arg0 string) ([32]byte, error) {
	return _Registry.Contract.Documents(&_Registry.CallOpts, arg0)
}

// ExchangeToken is a paid mutator transaction binding the contract method 0x8ed88dbf.
//
// Solidity: function exchangeToken(bytes32 vc) returns()
func (_Registry *RegistryTransactor) ExchangeToken(opts *bind.TransactOpts, vc [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "exchangeToken", vc)
}

// ExchangeToken is a paid mutator transaction binding the contract method 0x8ed88dbf.
//
// Solidity: function exchangeToken(bytes32 vc) returns()
func (_Registry *RegistrySession) ExchangeToken(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.ExchangeToken(&_Registry.TransactOpts, vc)
}

// ExchangeToken is a paid mutator transaction binding the contract method 0x8ed88dbf.
//
// Solidity: function exchangeToken(bytes32 vc) returns()
func (_Registry *RegistryTransactorSession) ExchangeToken(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.ExchangeToken(&_Registry.TransactOpts, vc)
}

// RenewVC is a paid mutator transaction binding the contract method 0x69c09d54.
//
// Solidity: function renewVC(bytes32 vc) returns()
func (_Registry *RegistryTransactor) RenewVC(opts *bind.TransactOpts, vc [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renewVC", vc)
}

// RenewVC is a paid mutator transaction binding the contract method 0x69c09d54.
//
// Solidity: function renewVC(bytes32 vc) returns()
func (_Registry *RegistrySession) RenewVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RenewVC(&_Registry.TransactOpts, vc)
}

// RenewVC is a paid mutator transaction binding the contract method 0x69c09d54.
//
// Solidity: function renewVC(bytes32 vc) returns()
func (_Registry *RegistryTransactorSession) RenewVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RenewVC(&_Registry.TransactOpts, vc)
}

// RevokeVC is a paid mutator transaction binding the contract method 0xd5c77944.
//
// Solidity: function revokeVC(bytes32 vc) returns()
func (_Registry *RegistryTransactor) RevokeVC(opts *bind.TransactOpts, vc [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeVC", vc)
}

// RevokeVC is a paid mutator transaction binding the contract method 0xd5c77944.
//
// Solidity: function revokeVC(bytes32 vc) returns()
func (_Registry *RegistrySession) RevokeVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RevokeVC(&_Registry.TransactOpts, vc)
}

// RevokeVC is a paid mutator transaction binding the contract method 0xd5c77944.
//
// Solidity: function revokeVC(bytes32 vc) returns()
func (_Registry *RegistryTransactorSession) RevokeVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RevokeVC(&_Registry.TransactOpts, vc)
}

// UploadVC is a paid mutator transaction binding the contract method 0x8b685a0a.
//
// Solidity: function uploadVC(bytes32 vc) returns()
func (_Registry *RegistryTransactor) UploadVC(opts *bind.TransactOpts, vc [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "uploadVC", vc)
}

// UploadVC is a paid mutator transaction binding the contract method 0x8b685a0a.
//
// Solidity: function uploadVC(bytes32 vc) returns()
func (_Registry *RegistrySession) UploadVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UploadVC(&_Registry.TransactOpts, vc)
}

// UploadVC is a paid mutator transaction binding the contract method 0x8b685a0a.
//
// Solidity: function uploadVC(bytes32 vc) returns()
func (_Registry *RegistryTransactorSession) UploadVC(vc [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UploadVC(&_Registry.TransactOpts, vc)
}
