// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testCon

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

// TestConMetaData contains all meta data concerning the TestCon contract.
var TestConMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"getEthSignedMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethSignedMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// TestConABI is the input ABI used to generate the binding from.
// Deprecated: Use TestConMetaData.ABI instead.
var TestConABI = TestConMetaData.ABI

// TestCon is an auto generated Go binding around an Ethereum contract.
type TestCon struct {
	TestConCaller     // Read-only binding to the contract
	TestConTransactor // Write-only binding to the contract
	TestConFilterer   // Log filterer for contract events
}

// TestConCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestConCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestConTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestConTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestConFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestConFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestConSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestConSession struct {
	Contract     *TestCon          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestConCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestConCallerSession struct {
	Contract *TestConCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestConTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestConTransactorSession struct {
	Contract     *TestConTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestConRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestConRaw struct {
	Contract *TestCon // Generic contract binding to access the raw methods on
}

// TestConCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestConCallerRaw struct {
	Contract *TestConCaller // Generic read-only contract binding to access the raw methods on
}

// TestConTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestConTransactorRaw struct {
	Contract *TestConTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestCon creates a new instance of TestCon, bound to a specific deployed contract.
func NewTestCon(address common.Address, backend bind.ContractBackend) (*TestCon, error) {
	contract, err := bindTestCon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestCon{TestConCaller: TestConCaller{contract: contract}, TestConTransactor: TestConTransactor{contract: contract}, TestConFilterer: TestConFilterer{contract: contract}}, nil
}

// NewTestConCaller creates a new read-only instance of TestCon, bound to a specific deployed contract.
func NewTestConCaller(address common.Address, caller bind.ContractCaller) (*TestConCaller, error) {
	contract, err := bindTestCon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestConCaller{contract: contract}, nil
}

// NewTestConTransactor creates a new write-only instance of TestCon, bound to a specific deployed contract.
func NewTestConTransactor(address common.Address, transactor bind.ContractTransactor) (*TestConTransactor, error) {
	contract, err := bindTestCon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestConTransactor{contract: contract}, nil
}

// NewTestConFilterer creates a new log filterer instance of TestCon, bound to a specific deployed contract.
func NewTestConFilterer(address common.Address, filterer bind.ContractFilterer) (*TestConFilterer, error) {
	contract, err := bindTestCon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestConFilterer{contract: contract}, nil
}

// bindTestCon binds a generic wrapper to an already deployed contract.
func bindTestCon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestConABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCon *TestConRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCon.Contract.TestConCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCon *TestConRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCon.Contract.TestConTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCon *TestConRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCon.Contract.TestConTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCon *TestConCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCon *TestConTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCon *TestConTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCon.Contract.contract.Transact(opts, method, params...)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_TestCon *TestConCaller) GetEthSignedMessageHash(opts *bind.CallOpts, _messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TestCon.contract.Call(opts, &out, "getEthSignedMessageHash", _messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_TestCon *TestConSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _TestCon.Contract.GetEthSignedMessageHash(&_TestCon.CallOpts, _messageHash)
}

// GetEthSignedMessageHash is a free data retrieval call binding the contract method 0xfa540801.
//
// Solidity: function getEthSignedMessageHash(bytes32 _messageHash) pure returns(bytes32)
func (_TestCon *TestConCallerSession) GetEthSignedMessageHash(_messageHash [32]byte) ([32]byte, error) {
	return _TestCon.Contract.GetEthSignedMessageHash(&_TestCon.CallOpts, _messageHash)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_TestCon *TestConCaller) GetMessageHash(opts *bind.CallOpts, _message string) ([32]byte, error) {
	var out []interface{}
	err := _TestCon.contract.Call(opts, &out, "getMessageHash", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_TestCon *TestConSession) GetMessageHash(_message string) ([32]byte, error) {
	return _TestCon.Contract.GetMessageHash(&_TestCon.CallOpts, _message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0xb446f3b2.
//
// Solidity: function getMessageHash(string _message) pure returns(bytes32)
func (_TestCon *TestConCallerSession) GetMessageHash(_message string) ([32]byte, error) {
	return _TestCon.Contract.GetMessageHash(&_TestCon.CallOpts, _message)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address, bytes32 r, bytes32 s, uint8 v)
func (_TestCon *TestConCaller) Recover(opts *bind.CallOpts, _ethSignedMessageHash [32]byte, _sig []byte) (common.Address, [32]byte, [32]byte, uint8, error) {
	var out []interface{}
	err := _TestCon.contract.Call(opts, &out, "recover", _ethSignedMessageHash, _sig)

	if err != nil {
		return *new(common.Address), *new([32]byte), *new([32]byte), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return out0, out1, out2, out3, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address, bytes32 r, bytes32 s, uint8 v)
func (_TestCon *TestConSession) Recover(_ethSignedMessageHash [32]byte, _sig []byte) (common.Address, [32]byte, [32]byte, uint8, error) {
	return _TestCon.Contract.Recover(&_TestCon.CallOpts, _ethSignedMessageHash, _sig)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 _ethSignedMessageHash, bytes _sig) pure returns(address, bytes32 r, bytes32 s, uint8 v)
func (_TestCon *TestConCallerSession) Recover(_ethSignedMessageHash [32]byte, _sig []byte) (common.Address, [32]byte, [32]byte, uint8, error) {
	return _TestCon.Contract.Recover(&_TestCon.CallOpts, _ethSignedMessageHash, _sig)
}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes _sig) pure returns(bool)
func (_TestCon *TestConCaller) Verify(opts *bind.CallOpts, _signer common.Address, _message string, _sig []byte) (bool, error) {
	var out []interface{}
	err := _TestCon.contract.Call(opts, &out, "verify", _signer, _message, _sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes _sig) pure returns(bool)
func (_TestCon *TestConSession) Verify(_signer common.Address, _message string, _sig []byte) (bool, error) {
	return _TestCon.Contract.Verify(&_TestCon.CallOpts, _signer, _message, _sig)
}

// Verify is a free data retrieval call binding the contract method 0x2dd34f0f.
//
// Solidity: function verify(address _signer, string _message, bytes _sig) pure returns(bool)
func (_TestCon *TestConCallerSession) Verify(_signer common.Address, _message string, _sig []byte) (bool, error) {
	return _TestCon.Contract.Verify(&_TestCon.CallOpts, _signer, _message, _sig)
}
