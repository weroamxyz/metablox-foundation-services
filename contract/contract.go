package contract

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/testCon"
)

var client *ethclient.Client
var instance *testCon.TestCon
var foundationPrivateKey *ecdsa.PrivateKey

func Init() error {
	client, err := ethclient.Dial("https://api.s0.b.hmny.io")
	if err != nil {
		return err
	}
	foundationPrivateKey, err = crypto.HexToECDSA("fdebd2c79a17bbea3f69b6ec146bc49b968a63bd24ec342e1bd22830d13f2687")
	if err != nil {
		return err
	}
	address := common.HexToAddress("0xD5ef7723BDA781212BB0b3609f712D7317618B44")
	instance, err = testCon.NewTestCon(address, client)
	if err != nil {
		return err
	}

	return nil
}

func CreateVC(vcBytes [32]byte) error {
	/*publicKey := foundationPrivateKey.Public()	//todo: uncomment once smart contract is ready
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(foundationPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	_, err = instance.UploadVC(auth, vcBytes)
	if err != nil {
		return err
	}*/

	return nil
}

func RenewVC(vcBytes [32]byte) error {
	/*publicKey := foundationPrivateKey.Public()	//todo: uncomment once smart contract is ready
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(foundationPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	_, err = instance.RenewVC(auth, vcBytes)
	if err != nil {
		return err
	}*/

	return nil
}

func RevokeVC(vcBytes [32]byte) error {
	/*publicKey := foundationPrivateKey.Public()	//todo: uncomment once smart contract is ready
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(foundationPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	_, err = instance.RevokeVC(auth, vcBytes)
	if err != nil {
		return err
	}*/

	return nil
}

func UploadDocument(docBytes [32]byte) error { //todo: actual implementation
	return nil
}

func GetDocument(targetDID string) (*models.DIDDocument, [32]byte, error) { //todo: actual implementation
	placeholderDoc := models.GenerateTestDIDDocument()
	placeholderHash := [32]byte{159, 210, 117, 26, 68, 195, 94, 82, 100, 225, 26, 113, 147, 246, 48, 225, 11, 103, 151, 249, 84, 104, 245, 122, 25, 36, 253, 166, 177, 201, 51, 0} //sha256.Sum256(ConvertDocToBytes(*placeholderDoc))
	return placeholderDoc, placeholderHash, nil
}

func TestSignatures(privateKey *ecdsa.PrivateKey, message string) error {
	signer := common.HexToAddress("0xB1453Ab8a8BBeB66098023138e070fBEa1624184")
	messageHash := crypto.Keccak256Hash([]byte(message))
	var comboBytes []byte
	comboBytes = bytes.Join([][]byte{comboBytes, []byte("\x19Ethereum Signed Message:\n32"), messageHash.Bytes()}, []byte{})
	fullHash := crypto.Keccak256Hash(comboBytes)
	signature, err := crypto.Sign(fullHash.Bytes(), privateKey)
	if err != nil {
		return err
	}
	signature[64] += 27 //have to increment this manually as the smart contract expects v to be 27 or 28, while the crypto package generates it as 0 or 1

	result, err := instance.Verify(nil, signer, message, signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return nil
}
