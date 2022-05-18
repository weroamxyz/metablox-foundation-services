package key

import (
	"bytes"
	"crypto/ecdsa"
	"os"

	"github.com/MetaBloxIO/metablox-foundation-services/errval"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	gojose "gopkg.in/square/go-jose.v2"
)

func GenerateNewPrivateKey(fileName string) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privData := crypto.FromECDSA(privKey)

	saveLocation := viper.GetString("storage.key_saving")

	err = os.WriteFile(saveLocation+fileName, privData, 0644)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func LoadPrivateKey(keyFileName string) (*ecdsa.PrivateKey, error) {
	loadLocation := viper.GetString("storage.key_loading")
	privData, err := os.ReadFile(loadLocation + keyFileName)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.ToECDSA(privData)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

//load private key from file, or create it if it doesn't exist yet
func GetIssuerPrivateKey() (*ecdsa.PrivateKey, error) {
	fileName := viper.GetString("storage.issuer_key_file")
	key, err := LoadPrivateKey(fileName)
	if err != nil {
		if err.Error() == "open ./wallet/issuer: no such file or directory" {
			return GenerateNewPrivateKey(fileName)
		}
		return nil, err
	}
	return key, nil
}

func CreateJWSSignature(privKey *ecdsa.PrivateKey, message []byte) (string, error) {
	signer, err := gojose.NewSigner(gojose.SigningKey{Algorithm: gojose.ES256, Key: privKey}, nil)
	if err != nil {
		return "", err
	}

	signature, err := signer.Sign(message)
	if err != nil {
		return "", err
	}

	compactserialized, err := signature.DetachedCompactSerialize()
	if err != nil {
		return "", err
	}
	return compactserialized, nil
}

func VerifyJWSSignature(signature string, pubKey *ecdsa.PublicKey, message []byte) (bool, error) {
	sigObject, err := gojose.ParseDetached(signature, message)
	if err != nil {
		return false, err
	}

	result, err := sigObject.Verify(pubKey)
	if err != nil {
		return false, err
	}

	if !bytes.Equal(message, result) {
		return false, nil
	} else {
		return true, nil
	}
}

func CompareAddresses(vm models.VerificationMethod, pubKey *ecdsa.PublicKey) (bool, error) {
	givenAddress := crypto.PubkeyToAddress(*pubKey)
	givenAccountID := "eip155:1:" + givenAddress.Hex()
	if vm.BlockchainAccountId != givenAccountID {
		return false, errval.ErrWrongAddress
	}

	return true, nil
}
