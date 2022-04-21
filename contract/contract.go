package contract

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
