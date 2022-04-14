package key

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	gojose "gopkg.in/square/go-jose.v2"
)

func GenerateNewPrivateKey() (*ecdsa.PrivateKey, string, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, "", err
	}

	privData := crypto.FromECDSA(privKey)

	saveLocation := viper.GetString("storage.key_saving")
	fileNumber := 1

	for {
		if _, err := os.Stat(saveLocation + "privateKey" + strconv.Itoa(fileNumber)); errors.Is(err, os.ErrNotExist) {
			break
		}
		fileNumber++
	}

	fileName := "privateKey" + strconv.Itoa(fileNumber)

	err = os.WriteFile(saveLocation+fileName, privData, 0644)
	if err != nil {
		return nil, "", err
	}

	return privKey, fileName, nil
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
