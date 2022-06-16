package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/registry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
	"time"
)

var (
	privateKey, _ = crypto.HexToECDSA("01F903CE0C960FF3A9E68E80FF5FFC344358D80CE1C221C3F9711AF07F83A3BD")
	testRpcUrl    = "https://api.s0.ps.hmny.io"
	testContract  = common.HexToAddress("0x0b9269e8947e46Bb60FFc54C137e7093907fD273")
)

func TestCheckSignature(t *testing.T) {
	register := &models.RegisterDID{
		Did:     "did:metablox:506690700",
		Account: "0x611b915b936Fde54Dfc309c0B31430aD345c4596",
		SigV:    27,
		SigR:    "0xf008542cc23c47972fdce69cb743d372d88917986f181c7c495005b75c374a86",
		SigS:    "0x61371be1c5ccb45bc5490453ec70ac1112a7ce45413fae880fc62ff608d4ca83",
	}
	userAddress := common.HexToAddress(register.Account)
	// Use fixed nonce =  1 here for testing
	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(register.Did), userAddress.Bytes(), []byte(big.NewInt(0).String()), []byte("register")}, nil)
	msgHash := crypto.Keccak256Hash(messageBytes)
	comboHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), msgHash.Bytes())
	userPub, err := crypto.SigToPub(comboHash.Bytes(), register.ToSigBytes())
	assert.NoError(t, err)
	assert.Equal(t, userAddress, crypto.PubkeyToAddress(*userPub))

	// Test error sigV
	register.SigV = 28
	userPub1, err := crypto.SigToPub(comboHash.Bytes(), register.ToSigBytes())
	assert.NoError(t, err)
	assert.NotEqual(t, userAddress, crypto.PubkeyToAddress(*userPub1))
}

func TestEstimateGas(t *testing.T) {
	// private
	key, err := crypto.HexToECDSA("877a728b5f40a375ea97914bd44bf31419ae6a1bb39eace7fe49cdf915c1183b")
	assert.NoError(t, err)
	userAddress := crypto.PubkeyToAddress(key.PublicKey)
	did := "did:metablox:123456789ABC"
	testClient, err := ethclient.Dial(testRpcUrl)
	assert.NoError(t, err)
	abi, err := registry.RegistryMetaData.GetAbi()
	assert.NoError(t, err, "load abi data failed")
	testInstance, err := registry.NewRegistry(testContract, testClient)
	assert.NoError(t, err, "initial test instance failed")
	nonce, err := testInstance.Nonce(nil, userAddress)
	fmt.Printf("print nonce : %s\n", nonce)
	assert.NoError(t, err, "get nonce failed")
	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(did), userAddress.Bytes(), []byte(nonce.String()), []byte("register")}, nil)
	r, s, v, err := createSignatureFromMessage(messageBytes, key)

	tempRegister := &models.RegisterDID{
		Did:     did,
		Account: userAddress.Hex(),
		SigV:    v,
		SigR:    hexutil.Encode(r[:]),
		SigS:    hexutil.Encode(s[:]),
	}
	jsonRegister, _ := json.MarshalIndent(tempRegister, "", "\t")
	fmt.Printf("\nprint request json:\n %s\n", jsonRegister)

	input, err := abi.Pack("registerDid", did, userAddress, v, r, s)
	assert.NoError(t, err)

	msg := ethereum.CallMsg{From: userAddress, To: &testContract, Data: input}
	gas, err := testClient.EstimateGas(context.Background(), msg)
	assert.Equal(t, uint64(0xf407), gas)
}

func TestRegisterDid(t *testing.T) {
	randomkey, err := crypto.GenerateKey()
	fmt.Println("randomkey generated: ", hexutil.Encode(crypto.FromECDSA(randomkey)))
	randomdid := strconv.Itoa(time.Now().Nanosecond())
	address := crypto.PubkeyToAddress(randomkey.PublicKey)

	testClient, err := ethclient.Dial(testRpcUrl)
	assert.NoError(t, err)
	testInstance, err := registry.NewRegistry(testContract, testClient)
	assert.NoError(t, err, "initial test instance failed")
	nonce, err := testInstance.Nonce(nil, address)

	fmt.Println("nonce=", nonce)
	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(randomdid), address.Bytes(), []byte(nonce.String()), []byte("register")}, nil)
	r, s, v, err := createSignatureFromMessage(messageBytes, randomkey)
	assert.NoError(t, err, "gen signature failed")

	auth, err := tempGenerateAuth(privateKey, testClient)
	assert.NoError(t, err)

	tempRegister := &models.RegisterDID{
		Did:     "did:metablox:" + randomdid,
		Account: address.Hex(),
		SigV:    v,
		SigR:    hexutil.Encode(r[:]),
		SigS:    hexutil.Encode(s[:]),
	}
	jsonRegister, _ := json.MarshalIndent(tempRegister, "", "\t")
	fmt.Printf("\nprint request json:\n %s\n", jsonRegister)
	fmt.Printf("\nprint contract input:\n \tr=%v\n\ts=%v\n\tv=%d\n\taccount=%s\n\tdid=%s\n", r, s, v, address.Hex(), randomdid)
	tx, err := testInstance.RegisterDid(auth, randomdid, address, v, r, s)
	assert.NoError(t, err)
	json, _ := tx.MarshalJSON()
	fmt.Printf("\nprint tx:\n \t%v\n", string(json))

	var receipt *types.Receipt
	for receipt == nil {
		receipt, _ = testClient.TransactionReceipt(context.Background(), tx.Hash())
		time.Sleep(time.Second)
	}
	receiptJson, _ := receipt.MarshalJSON()
	fmt.Printf("\nprint receipt:\n \t%v\n", string(receiptJson))
	assert.Equal(t, uint64(1), receipt.Status)

}

func tempGenerateAuth(privateKey *ecdsa.PrivateKey, backend *ethclient.Client) (*bind.TransactOpts, error) {
	chainid, err := backend.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainid)
	if err != nil {
		return nil, err
	}
	authNonce, err := backend.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}

	gasPrice, err := backend.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(authNonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth, nil
}
