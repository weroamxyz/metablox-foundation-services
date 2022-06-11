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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	privateKey, _ = crypto.HexToECDSA("9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968")
	testRpcUrl    = "https://api.s0.b.hmny.io"
	testContract  = common.HexToAddress("0xd3E90701C814aA7C5eBD1B62395311d5C7B71f5e")
	//testContract = common.HexToAddress("0xf880b97Be7c402Cc441895bF397c3f865BfE1Cb2")
)

func NewRegister() *models.RegisterDID {
	return &models.RegisterDID{
		Did:     "did:metablox:123456",
		Account: "0x56BdBb8eCB54570b5a3971Aaacf85040E7AC3B4F",
		SigV:    28,
		SigR:    "0x806a150ac6fc425c4948e9adba9257f7ab42d76f30c2ee32150a49951b7e86d2",
		SigS:    "0x5e071842982074a6baba33dd3a763af04c304a74f769adaebe5696b13bdfea5a",
	}
}

func TestCheckSignature(t *testing.T) {
	register := NewRegister()
	userAddress := common.HexToAddress(register.Account)
	// Use fixed nonce =  1 here for testing
	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(register.Did), userAddress.Bytes(), common.LeftPadBytes(big.NewInt(0).Bytes(), 32), []byte("register")}, nil)
	msgHash := crypto.Keccak256Hash(messageBytes)
	comboHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), msgHash.Bytes())
	userPub, err := crypto.SigToPub(comboHash.Bytes(), register.ToSigBytes())
	assert.NoError(t, err)
	assert.Equal(t, userAddress, crypto.PubkeyToAddress(*userPub))

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
	//var nonceByte [32]byte
	//copy(nonceByte[:], nonce.Bytes())

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
	register := NewRegister()
	testClient, err := ethclient.Dial(testRpcUrl)
	assert.NoError(t, err)
	testInstance, err := registry.NewRegistry(testContract, testClient)
	assert.NoError(t, err)
	auth, err := tempGenerateAuth(privateKey, testClient)
	assert.NoError(t, err)
	_, err = testInstance.RegisterDid(auth, register.Did, register.Address(), register.SigV, register.SigRBytes32(), register.SigSBytes32())
	assert.NoError(t, err)

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
