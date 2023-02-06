package did

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"os"

	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	gojose "gopkg.in/square/go-jose.v2"
)

// create new private key and save it to a target file
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

// load an existing private key from a target file
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

// load private key from file, or create it if it doesn't exist yet
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

// use a private key and a message to create a JWS format signature
func CreateJWSSignature(privKey *ecdsa.PrivateKey, message []byte) (string, error) {
	signer, err := gojose.NewSigner(gojose.SigningKey{Algorithm: gojose.ES256, Key: privKey}, nil)
	if err != nil {
		return "", err
	}
	c := privKey.PublicKey.Curve
	N := c.Params().N

	signature, err := signer.Sign(message)
	if err != nil {
		return "", err
	}

	sBytes := make([]byte, 32)
	copy(sBytes, signature.Signatures[0].Signature[32:])
	var s = new(big.Int).SetBytes(sBytes)

	m := new(big.Int).Div(N, big.NewInt(2))
	q := s.Cmp(m)
	if q > 0 || s.Cmp(big.NewInt(1)) < 0 {
		sub := new(big.Int).Sub(N, s)
		s = new(big.Int).Mod(sub, N)
		newBytes := s.Bytes()
		sByte := make([]byte, 32)
		copy(sByte[32-len(newBytes):], newBytes)
		copy(signature.Signatures[0].Signature[32:], sByte)
	}

	compactserialized, err := signature.DetachedCompactSerialize()
	if err != nil {
		return "", err
	}
	return compactserialized, nil
}

// verify a JWS format signature using the matching public key and the original message
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

// make sure that the address created from pubKey matches the address stored in vm's BlockChainAccountId field
func CompareAddresses(vm models.VerificationMethod, pubKey *ecdsa.PublicKey) bool {
	givenAddress := crypto.PubkeyToAddress(*pubKey)
	givenAccountID := "eip155:1666600000:" + givenAddress.Hex()
	if vm.BlockchainAccountId != givenAccountID {
		return false
	}

	return true
}
