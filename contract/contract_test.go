package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
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
	privateKey, _ = crypto.HexToECDSA("9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968")
	testRpcUrl    = "https://api.s0.b.hmny.io"
	testContract  = common.HexToAddress("0xd3E90701C814aA7C5eBD1B62395311d5C7B71f5e")
)

func TestCheckSignature(t *testing.T) {
	register := &models.RegisterDID{
		Did:     "did:metablox:123456",
		Account: "0x56BdBb8eCB54570b5a3971Aaacf85040E7AC3B4F",
		SigV:    28,
		SigR:    "0xee2ec79d91108a250de1a88bc2158ad07784a49edea283f1b4d34adf7e623a87",
		SigS:    "0x6a92eb22aa0189fb7009f55f2fa8170bfbbd5e7ff882051f199005b79f7439d2",
	}
	userAddress := common.HexToAddress(register.Account)
	// Use fixed nonce =  1 here for testing
	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(register.Did), userAddress.Bytes(), common.LeftPadBytes(big.NewInt(1).Bytes(), 32), []byte("register")}, nil)
	msgHash := crypto.Keccak256Hash(messageBytes)
	comboHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), msgHash.Bytes())
	userPub, err := crypto.SigToPub(comboHash.Bytes(), register.ToSigBytes())
	assert.NoError(t, err)
	assert.Equal(t, userAddress, crypto.PubkeyToAddress(*userPub))

	// Test error sigV
	register.SigV = 27
	userPub1, err := crypto.SigToPub(comboHash.Bytes(), register.ToSigBytes())
	assert.NoError(t, err)
	assert.NotEqual(t, userAddress, crypto.PubkeyToAddress(*userPub1))
}

func TestEstimateGas(t *testing.T) {

	//9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968
	userAddress := common.HexToAddress("0x56BdBb8eCB54570b5a3971Aaacf85040E7AC3B4F")
	did := "did:metablox:123456"
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
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(did), userAddress.Bytes(), common.LeftPadBytes(nonce.Bytes(), 32), []byte("register")}, nil)
	r, s, v, err := createSignatureFromMessage(messageBytes, privateKey)

	fmt.Println(hexutil.Encode(r[:]))
	fmt.Println(hexutil.Encode(s[:]))
	messageHash := crypto.Keccak256Hash(messageBytes)
	comboHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), messageHash.Bytes())

	var signbytes [65]byte
	copy(signbytes[:32], r[:])
	copy(signbytes[32:64], s[:])
	signbytes[64] = v - 27

	pubKey, _ := crypto.SigToPub(comboHash[:], signbytes[:])
	resultAddress := crypto.PubkeyToAddress(*pubKey)
	assert.Equal(t, resultAddress, userAddress)

	input, err := abi.Pack("registerDid", did, userAddress, v, r, s)

	assert.NoError(t, err)
	msg := ethereum.CallMsg{From: userAddress, To: &testContract, Data: input}
	_, err = testClient.EstimateGas(context.Background(), msg)
	assert.Equal(t, "execution reverted: did_exist", err.Error())
}

func TestRegisterDid(t *testing.T) {
	randomkey, err := crypto.GenerateKey()
	fmt.Println("randomkey generated: ", hexutil.Encode(crypto.FromECDSA(randomkey)))
	randomdid := "did:metablox:" + strconv.Itoa(time.Now().Nanosecond())
	address := crypto.PubkeyToAddress(randomkey.PublicKey)

	testClient, err := ethclient.Dial(testRpcUrl)
	assert.NoError(t, err)
	testInstance, err := registry.NewRegistry(testContract, testClient)
	assert.NoError(t, err, "initial test instance failed")
	nonce, err := testInstance.Nonce(nil, address)

	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(randomdid), address.Bytes(), common.LeftPadBytes(nonce.Bytes(), 32), []byte("register")}, nil)
	r, s, v, err := createSignatureFromMessage(messageBytes, randomkey)
	assert.NoError(t, err, "gen signature failed")

	auth, err := tempGenerateAuth(privateKey, testClient)
	assert.NoError(t, err)

	fmt.Printf("\nprint request json:\n \tr=%s\n\ts=%s\n\tv=%d\n\taccount=%s\n\tdid=%s\n", hexutil.Encode(r[:]), hexutil.Encode(s[:]), v, address.Hex(), randomdid)
	fmt.Printf("\nprint contract input:\n \tr=%v\n\ts=%v\n\tv=%d\n\taccount=%s\n\tdid=%s\n", r, s, v, address.Hex(), randomdid)
	tx, err := testInstance.RegisterDid(auth, randomdid, address, v, r, s)
	assert.NoError(t, err)
	json, _ := tx.MarshalJSON()
	fmt.Printf("\nprint tx:\n \t%v\n", string(json))

	var receipt *types.Receipt
	for receipt == nil {
		receipt, _ = testClient.TransactionReceipt(context.Background(), tx.Hash())
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
