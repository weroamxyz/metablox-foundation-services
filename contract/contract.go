package contract

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metabloxDID/models"
	"github.com/metabloxDID/registry"
)

var client *ethclient.Client
var instance *registry.Registry
var foundationPrivateKey *ecdsa.PrivateKey
var Address common.Address

func Init() error {
	var err error
	client, err = ethclient.Dial("https://api.s0.b.hmny.io")
	if err != nil {
		return err
	}
	foundationPrivateKey, err = crypto.HexToECDSA("fdebd2c79a17bbea3f69b6ec146bc49b968a63bd24ec342e1bd22830d13f2687")
	if err != nil {
		return err
	}
	Address = common.HexToAddress("0x8CeDd60c472164ab3aae55E69D9B7E514AB972d8")
	instance, err = registry.NewRegistry(Address, client)
	if err != nil {
		return err
	}

	return nil
}

func RegisterVC(vcBytes [32]byte) error {
	/*	fromAddress := crypto.PubkeyToAddress(foundationPrivateKey.PublicKey)	//todo: uncomment once smart contract is ready
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
	/*	fromAddress := crypto.PubkeyToAddress(foundationPrivateKey.PublicKey)	//todo: uncomment once smart contract is ready
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
	/*	fromAddress := crypto.PubkeyToAddress(foundationPrivateKey.PublicKey)	//todo: uncomment once smart contract is ready
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

func UploadDocument(document *models.DIDDocument, privateKey *ecdsa.PrivateKey) error { //todo: actual implementation
	//signer := common.HexToAddress("0xB1453Ab8a8BBeB66098023138e070fBEa1624184")
	//pubBytes := crypto.CompressPubkey(&privateKey.PublicKey)
	pubAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	//pubAddress := common.BytesToAddress(pubBytes)
	nonce, err := instance.Nonce(nil, pubAddress)
	if err != nil {
		return err
	}
	var messageBytes []byte
	var nonceBytes [32]byte

	copy(nonceBytes[:], nonce.Bytes())
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(document.ID), pubAddress.Bytes(), nonceBytes[:], []byte("registerOrUpdate")}, nil)
	messageHash := sha256.Sum256(messageBytes)
	signature, err := crypto.Sign(messageHash[:], privateKey)
	if err != nil {
		return err
	}
	var r [32]byte
	var s [32]byte
	var v uint8

	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v = signature[64] + 27 //have to increment this manually as the smart contract expects v to be 27 or 28, while the crypto package generates it as 0 or 1

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1666700000))
	if err != nil {
		return err
	}
	authNonce, err := client.PendingNonceAt(context.Background(), pubAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(authNonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	tx, err := instance.RegisterDid(auth, document.ID, pubAddress, v, r, s)
	if err != nil {
		return err
	}

	fmt.Println("transaction address: ", tx.Hash().Hex())
	return nil
}

func GetDocument(targetDID string) (*models.DIDDocument, [32]byte, error) { //todo: actual implementation
	placeholderDoc := models.GenerateTestDIDDocument()
	placeholderHash := [32]byte{159, 210, 117, 26, 68, 195, 94, 82, 100, 225, 26, 113, 147, 246, 48, 225, 11, 103, 151, 249, 84, 104, 245, 122, 25, 36, 253, 166, 177, 201, 51, 0} //sha256.Sum256(ConvertDocToBytes(*placeholderDoc))
	return placeholderDoc, placeholderHash, nil
}
