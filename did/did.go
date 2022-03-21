package did

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/models"
)

func CreateDID() (*models.DIDDocument, *ecdsa.PrivateKey, error) {

	document := new(models.DIDDocument)

	//document subject needs to keep this private key to use the authentication method and prove ownership
	privKey, err := secp256k1.NewECDSAPrivateKey()
	if err != nil {
		return nil, nil, err
	}

	privData, err := secp256k1.FromECDSAPrivateKey(privKey)
	if err != nil {
		return nil, nil, err
	}

	hash := sha256.New()
	hash.Write(privData)
	hashData := hash.Sum(nil)
	didString := base58.Encode(hashData)
	document.ID = "did:metablox:" + didString
	document.Context = "https://w3id.org/did/v1"
	document.Created = time.Now().Format(time.RFC3339)
	document.Updated = document.Created
	document.Version = 1

	pubData, err := secp256k1.FromECDSAPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}

	VM := models.VerificationMethod{}
	VM.ID = document.ID + "#verification"
	VM.Key = hex.EncodeToString(pubData)
	VM.Controller = document.ID
	VM.MethodType = "Secp256k1"

	document.VerificationMethod = append(document.VerificationMethod, VM)
	document.Authentication = VM.ID

	//once blockchain is implemented, will also need to upload the document to the blockchain

	return document, privKey, nil
}

func DocumentToJson(document *models.DIDDocument) ([]byte, error) {
	jsonDoc, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	return jsonDoc, nil
}

func JsonToDocument(jsonDoc []byte) (*models.DIDDocument, error) {
	document := models.CreateDIDDocument()
	err := json.Unmarshal(jsonDoc, document)
	if err != nil {
		return nil, err
	}
	return document, nil
}

//check format of DID string
func prepareDID(did string) ([]string, bool) {
	splitString := strings.Split(did, ":")
	if len(splitString) < 3 {
		fmt.Println("Not enough sections in DID")
		return nil, false
	}
	prefix := splitString[0]
	if prefix != "did" {
		fmt.Println("First section of DID was '" + prefix + "' instead of 'did'")
		return nil, false
	}

	methodName := splitString[1]
	methodNamePattern := `^[A-Za-z0-9]+$`
	match, _ := regexp.MatchString(methodNamePattern, methodName)
	if !match {
		fmt.Println("Method name '" + methodName + "' is formatted incorrectly")
		return nil, false
	}

	identifierPattern := `^([a-zA-Z0-9\._\-]*(%[0-9A-Fa-f][0-9A-Fa-f])*)*$`
	identifierExp, _ := regexp.Compile(identifierPattern)

	for i := 2; i < len(splitString); i++ {
		identifierSection := splitString[i]
		match = identifierExp.MatchString(identifierSection)
		if !match {
			fmt.Println("Identifier section '" + identifierSection + "' is formatted incorrectly")
			return nil, false
		}

		if i == len(splitString)-1 && len(identifierSection) == 0 {
			fmt.Println("Final portion of identifier is empty")
			return nil, false
		}
	}
	return splitString, true
}

func Resolve(did string, options *models.ResolutionOptions) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	splitDID, valid := prepareDID(did)
	if !valid {
		return &models.ResolutionMetadata{Error: "invalid Did"}, nil, &models.DocumentMetadata{}
	}

	methodName := splitDID[1]
	switch methodName {
	case "metablox":
		return MetabloxRead(splitDID[2], options)
	default:
		fmt.Println("Unable to resolve unknown method type '" + methodName + "'")
		return &models.ResolutionMetadata{Error: "methodNotSupported"}, nil, &models.DocumentMetadata{}
	}
}

func ResolveRepresentation(did string, options *models.RepresentationResolutionOptions) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	splitDID, valid := prepareDID(did)
	if !valid {
		return &models.RepresentationResolutionMetadata{Error: "invalid Did"}, nil, &models.DocumentMetadata{}
	}

	methodName := splitDID[1]
	switch methodName {
	case "metablox":
		return MetabloxReadRepresentation(splitDID[2], options)
	default:
		fmt.Println("Unable to resolve unknown method type '" + methodName + "'")
		return &models.RepresentationResolutionMetadata{Error: "methodNotSupported"}, nil, &models.DocumentMetadata{}
	}
}

func AuthenticateDocumentSubject(document *models.DIDDocument, message string, x, y *big.Int) (bool, error) {
	//The subject of the document is the person who has the private key matching the public key in the Authentication verification method

	//Get Authentication VM
	var authenticationMethod *models.VerificationMethod
	for _, vm := range document.VerificationMethod {
		if vm.ID == document.Authentication {
			authenticationMethod = &vm
			break
		}
	}
	if authenticationMethod == nil {
		return false, errors.New("Failed to find authentication method '" + document.Authentication + "'")
	}

	switch authenticationMethod.MethodType {
	case "Secp256k1":
		pubData, err := hex.DecodeString(authenticationMethod.Key)
		if err != nil {
			return false, err
		}
		pubKey, err := secp256k1.ToECDSAPublicKey(pubData)
		if err != nil {
			return false, err
		}

		result := ecdsa.Verify(pubKey, []byte(message), x, y)
		return result, nil

	default:
		return false, errors.New("Unable to resolve unknown verification method type '" + authenticationMethod.MethodType + "'")
	}
}

func MetabloxRead(identifier string, options *models.ResolutionOptions) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	//Once blockchain is implemented, this function will use the identifier to retrieve the matching document from the blockchain (if it exists).
	//Some mechanism should likely be in place to ensure the document was not modified during the transfer, ex. an encrypted hash.
	return nil, nil, nil
}

func MetabloxReadRepresentation(identifier string, options *models.RepresentationResolutionOptions) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	//Should be similar to MetabloxRead, but returns the document in a specific representation format.
	//Representation type is included in options and returned in resolution metadata
	return nil, nil, nil
}
