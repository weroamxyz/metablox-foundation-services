package contract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxDID/credentials"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/registry"
)

var client *ethclient.Client
var instance *registry.Registry
var foundationPrivateKey *ecdsa.PrivateKey

func Init() error {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		return err
	}
	foundationPrivateKey, err = crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		return err
	}
	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err = registry.NewRegistry(address, client)
	if err != nil {
		return err
	}
	return nil
}

func CreateVC(vc *models.VerifiableCredential) error {

	publicKey := foundationPrivateKey.Public()
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
	vcByte := [32]byte{}
	copy(vcByte[:], credentials.ConvertVCToBytes(*vc))
	_, err = instance.UploadVC(auth, vcByte)
	if err != nil {
		return err
	}

	return nil
}

func RenewVC(vc *models.VerifiableCredential) error {

	publicKey := foundationPrivateKey.Public()
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
	vcByte := [32]byte{}
	copy(vcByte[:], credentials.ConvertVCToBytes(*vc))
	_, err = instance.RenewVC(auth, vcByte)
	if err != nil {
		return err
	}

	return nil
}

func RevokeVC(vc *models.VerifiableCredential) error {

	publicKey := foundationPrivateKey.Public()
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
	vcByte := [32]byte{}
	copy(vcByte[:], credentials.ConvertVCToBytes(*vc))
	_, err = instance.RevokeVC(auth, vcByte)
	if err != nil {
		return err
	}

	return nil
}

func GetDocument(targetDID string) ([32]byte, error) { //not sure what format we're using to store documents, so currently leaving it as bytes
	documentData, err := instance.Documents(nil, targetDID)
	if err != nil {
		return [32]byte{}, err
	}
	return documentData, nil
}
