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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"vcName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"VCSchemaChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"changed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"}],\"name\":\"createVcDef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"}],\"name\":\"registerDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"vcName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"setVcAttributeSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"nSigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"nSigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nSigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"oSigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"oSigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oSigS\",\"type\":\"bytes32\"}],\"name\":\"updateDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"vcIssuers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Registry *RegistryCaller) Changed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "changed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Registry *RegistrySession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.Changed(&_Registry.CallOpts, arg0)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed(address ) view returns(uint256)
func (_Registry *RegistryCallerSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.Changed(&_Registry.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(address)
func (_Registry *RegistryCaller) Dids(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "dids", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(address)
func (_Registry *RegistrySession) Dids(arg0 string) (common.Address, error) {
	return _Registry.Contract.Dids(&_Registry.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(address)
func (_Registry *RegistryCallerSession) Dids(arg0 string) (common.Address, error) {
	return _Registry.Contract.Dids(&_Registry.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Registry *RegistryCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Registry *RegistrySession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.Nonce(&_Registry.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_Registry *RegistryCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.Nonce(&_Registry.CallOpts, arg0)
}

// VcIssuers is a free data retrieval call binding the contract method 0x52b1cf64.
//
// Solidity: function vcIssuers(string ) view returns(string)
func (_Registry *RegistryCaller) VcIssuers(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "vcIssuers", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VcIssuers is a free data retrieval call binding the contract method 0x52b1cf64.
//
// Solidity: function vcIssuers(string ) view returns(string)
func (_Registry *RegistrySession) VcIssuers(arg0 string) (string, error) {
	return _Registry.Contract.VcIssuers(&_Registry.CallOpts, arg0)
}

// VcIssuers is a free data retrieval call binding the contract method 0x52b1cf64.
//
// Solidity: function vcIssuers(string ) view returns(string)
func (_Registry *RegistryCallerSession) VcIssuers(arg0 string) (string, error) {
	return _Registry.Contract.VcIssuers(&_Registry.CallOpts, arg0)
}

// CreateVcDef is a paid mutator transaction binding the contract method 0x4d618c5e.
//
// Solidity: function createVcDef(string name, string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistryTransactor) CreateVcDef(opts *bind.TransactOpts, name string, did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createVcDef", name, did, sigV, sigR, sigS)
}

// CreateVcDef is a paid mutator transaction binding the contract method 0x4d618c5e.
//
// Solidity: function createVcDef(string name, string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistrySession) CreateVcDef(name string, did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.CreateVcDef(&_Registry.TransactOpts, name, did, sigV, sigR, sigS)
}

// CreateVcDef is a paid mutator transaction binding the contract method 0x4d618c5e.
//
// Solidity: function createVcDef(string name, string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistryTransactorSession) CreateVcDef(name string, did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.CreateVcDef(&_Registry.TransactOpts, name, did, sigV, sigR, sigS)
}

// RegisterDid is a paid mutator transaction binding the contract method 0xd0923626.
//
// Solidity: function registerDid(string did, address account, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistryTransactor) RegisterDid(opts *bind.TransactOpts, did string, account common.Address, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerDid", did, account, sigV, sigR, sigS)
}

// RegisterDid is a paid mutator transaction binding the contract method 0xd0923626.
//
// Solidity: function registerDid(string did, address account, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistrySession) RegisterDid(did string, account common.Address, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterDid(&_Registry.TransactOpts, did, account, sigV, sigR, sigS)
}

// RegisterDid is a paid mutator transaction binding the contract method 0xd0923626.
//
// Solidity: function registerDid(string did, address account, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_Registry *RegistryTransactorSession) RegisterDid(did string, account common.Address, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterDid(&_Registry.TransactOpts, did, account, sigV, sigR, sigS)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xbb55343f.
//
// Solidity: function setAttributeSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistryTransactor) SetAttributeSigned(opts *bind.TransactOpts, did string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAttributeSigned", did, sigV, sigR, sigS, name, value)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xbb55343f.
//
// Solidity: function setAttributeSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistrySession) SetAttributeSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAttributeSigned(&_Registry.TransactOpts, did, sigV, sigR, sigS, name, value)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0xbb55343f.
//
// Solidity: function setAttributeSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistryTransactorSession) SetAttributeSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAttributeSigned(&_Registry.TransactOpts, did, sigV, sigR, sigS, name, value)
}

// SetVcAttributeSigned is a paid mutator transaction binding the contract method 0x46f4303a.
//
// Solidity: function setVcAttributeSigned(string vcName, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistryTransactor) SetVcAttributeSigned(opts *bind.TransactOpts, vcName string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setVcAttributeSigned", vcName, sigV, sigR, sigS, name, value)
}

// SetVcAttributeSigned is a paid mutator transaction binding the contract method 0x46f4303a.
//
// Solidity: function setVcAttributeSigned(string vcName, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistrySession) SetVcAttributeSigned(vcName string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetVcAttributeSigned(&_Registry.TransactOpts, vcName, sigV, sigR, sigS, name, value)
}

// SetVcAttributeSigned is a paid mutator transaction binding the contract method 0x46f4303a.
//
// Solidity: function setVcAttributeSigned(string vcName, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 name, bytes value) returns()
func (_Registry *RegistryTransactorSession) SetVcAttributeSigned(vcName string, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetVcAttributeSigned(&_Registry.TransactOpts, vcName, sigV, sigR, sigS, name, value)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xb6f8a115.
//
// Solidity: function updateDid(string did, address account, uint8 nSigV, bytes32 nSigR, bytes32 nSigS, uint8 oSigV, bytes32 oSigR, bytes32 oSigS) returns()
func (_Registry *RegistryTransactor) UpdateDid(opts *bind.TransactOpts, did string, account common.Address, nSigV uint8, nSigR [32]byte, nSigS [32]byte, oSigV uint8, oSigR [32]byte, oSigS [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateDid", did, account, nSigV, nSigR, nSigS, oSigV, oSigR, oSigS)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xb6f8a115.
//
// Solidity: function updateDid(string did, address account, uint8 nSigV, bytes32 nSigR, bytes32 nSigS, uint8 oSigV, bytes32 oSigR, bytes32 oSigS) returns()
func (_Registry *RegistrySession) UpdateDid(did string, account common.Address, nSigV uint8, nSigR [32]byte, nSigS [32]byte, oSigV uint8, oSigR [32]byte, oSigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateDid(&_Registry.TransactOpts, did, account, nSigV, nSigR, nSigS, oSigV, oSigR, oSigS)
}

// UpdateDid is a paid mutator transaction binding the contract method 0xb6f8a115.
//
// Solidity: function updateDid(string did, address account, uint8 nSigV, bytes32 nSigR, bytes32 nSigS, uint8 oSigV, bytes32 oSigR, bytes32 oSigS) returns()
func (_Registry *RegistryTransactorSession) UpdateDid(did string, account common.Address, nSigV uint8, nSigR [32]byte, nSigS [32]byte, oSigV uint8, oSigR [32]byte, oSigS [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateDid(&_Registry.TransactOpts, did, account, nSigV, nSigR, nSigS, oSigV, oSigR, oSigS)
}

// RegistryDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the Registry contract.
type RegistryDIDAttributeChangedIterator struct {
	Event *RegistryDIDAttributeChanged // Event containing the contract specifics and raw log

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
func (it *RegistryDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDIDAttributeChanged)
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
		it.Event = new(RegistryDIDAttributeChanged)
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
func (it *RegistryDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDIDAttributeChanged represents a DIDAttributeChanged event raised by the Registry contract.
type RegistryDIDAttributeChanged struct {
	Did   common.Hash
	Name  [32]byte
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x45731992f6f6bd99a592d32cd338d6875a784dbb138bc7a051f541d0f6d56e37.
//
// Solidity: event DIDAttributeChanged(string indexed did, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, did []string) (*RegistryDIDAttributeChangedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DIDAttributeChanged", didRule)
	if err != nil {
		return nil, err
	}
	return &RegistryDIDAttributeChangedIterator{contract: _Registry.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x45731992f6f6bd99a592d32cd338d6875a784dbb138bc7a051f541d0f6d56e37.
//
// Solidity: event DIDAttributeChanged(string indexed did, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *RegistryDIDAttributeChanged, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DIDAttributeChanged", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDIDAttributeChanged)
				if err := _Registry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
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

// ParseDIDAttributeChanged is a log parse operation binding the contract event 0x45731992f6f6bd99a592d32cd338d6875a784dbb138bc7a051f541d0f6d56e37.
//
// Solidity: event DIDAttributeChanged(string indexed did, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) ParseDIDAttributeChanged(log types.Log) (*RegistryDIDAttributeChanged, error) {
	event := new(RegistryDIDAttributeChanged)
	if err := _Registry.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryVCSchemaChangedIterator is returned from FilterVCSchemaChanged and is used to iterate over the raw logs and unpacked data for VCSchemaChanged events raised by the Registry contract.
type RegistryVCSchemaChangedIterator struct {
	Event *RegistryVCSchemaChanged // Event containing the contract specifics and raw log

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
func (it *RegistryVCSchemaChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVCSchemaChanged)
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
		it.Event = new(RegistryVCSchemaChanged)
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
func (it *RegistryVCSchemaChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVCSchemaChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVCSchemaChanged represents a VCSchemaChanged event raised by the Registry contract.
type RegistryVCSchemaChanged struct {
	VcName common.Hash
	Name   [32]byte
	Value  []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVCSchemaChanged is a free log retrieval operation binding the contract event 0x7650c16466dfc67d3aa1b3b380b7425844e301ce12fc320f5ce442a39d222af4.
//
// Solidity: event VCSchemaChanged(string indexed vcName, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) FilterVCSchemaChanged(opts *bind.FilterOpts, vcName []string) (*RegistryVCSchemaChangedIterator, error) {

	var vcNameRule []interface{}
	for _, vcNameItem := range vcName {
		vcNameRule = append(vcNameRule, vcNameItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VCSchemaChanged", vcNameRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVCSchemaChangedIterator{contract: _Registry.contract, event: "VCSchemaChanged", logs: logs, sub: sub}, nil
}

// WatchVCSchemaChanged is a free log subscription operation binding the contract event 0x7650c16466dfc67d3aa1b3b380b7425844e301ce12fc320f5ce442a39d222af4.
//
// Solidity: event VCSchemaChanged(string indexed vcName, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) WatchVCSchemaChanged(opts *bind.WatchOpts, sink chan<- *RegistryVCSchemaChanged, vcName []string) (event.Subscription, error) {

	var vcNameRule []interface{}
	for _, vcNameItem := range vcName {
		vcNameRule = append(vcNameRule, vcNameItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VCSchemaChanged", vcNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVCSchemaChanged)
				if err := _Registry.contract.UnpackLog(event, "VCSchemaChanged", log); err != nil {
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

// ParseVCSchemaChanged is a log parse operation binding the contract event 0x7650c16466dfc67d3aa1b3b380b7425844e301ce12fc320f5ce442a39d222af4.
//
// Solidity: event VCSchemaChanged(string indexed vcName, bytes32 name, bytes value)
func (_Registry *RegistryFilterer) ParseVCSchemaChanged(log types.Log) (*RegistryVCSchemaChanged, error) {
	event := new(RegistryVCSchemaChanged)
	if err := _Registry.contract.UnpackLog(event, "VCSchemaChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
