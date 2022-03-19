package did

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/models"
)

func CreateDID() (*models.DIDDocument, error) {

	document := new(models.DIDDocument)

	privKey, err := secp256k1.NewECDSAPrivateKey()
	if err != nil {
		return nil, err
	}

	privData, err := secp256k1.FromECDSAPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	hash.Write(privData)
	hashData := hash.Sum(nil)
	didString := base58.Encode(hashData)
	document.ID = "did:metablox:" + didString
	document.Context = "https://w3id.org/did/v1"
	document.Created = time.Now().Format("2006-01-02 15:04:05")
	document.Updated = document.Created
	document.Version = 1

	pubData, err := secp256k1.FromECDSAPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}

	VM := models.VerificationMethod{}
	VM.ID = document.ID + "#verification"
	VM.Key = hex.EncodeToString(pubData)
	VM.Controller = document.ID
	VM.MethodType = "Secp256k1"

	document.VerificationMethod = append(document.VerificationMethod, VM)
	document.Authentication = VM.ID

	return document, nil
}
