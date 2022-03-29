package key

import (
	"crypto/ecdsa"
	"errors"
	"os"
	"strconv"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/spf13/viper"
)

func GenerateNewPrivateKey() (*ecdsa.PrivateKey, string, error) {
	privKey, err := secp256k1.NewECDSAPrivateKey()
	if err != nil {
		return nil, "", err
	}

	privData, err := secp256k1.FromECDSAPrivateKey(privKey)
	if err != nil {
		return nil, "", err
	}

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

	privKey, err := secp256k1.ToECDSAPrivateKey(privData)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}
