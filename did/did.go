package did

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/dappley/go-dappley/crypto/keystore/secp256k1"
	"github.com/metabloxDID/models"
	"github.com/mr-tron/base58"
	"github.com/multiformats/go-multibase"
)

func CreateDID(privKey *ecdsa.PrivateKey) (*models.DIDDocument, error) {

	document := new(models.DIDDocument)

	privData, err := secp256k1.FromECDSAPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()
	hash.Write(privData)
	hashData := hash.Sum(nil)
	didString := base58.Encode(hashData)
	document.ID = "did:metablox:" + didString
	document.Context = make([]string, 0)
	document.Context = append(document.Context, "https://w3id.org/did/v1")
	document.Context = append(document.Context, "https://ns.did.ai/suites/secp256k1-2019/v1/")
	document.Created = time.Now().Format(time.RFC3339)
	document.Updated = document.Created
	document.Version = 1

	pubData, err := secp256k1.FromECDSAPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}

	VM := models.VerificationMethod{}
	VM.ID = document.ID + "#verification"
	VM.MultibaseKey, err = multibase.Encode(multibase.Base58BTC, pubData)
	if err != nil {
		return nil, err
	}
	VM.Controller = document.ID
	VM.MethodType = "EcdsaSecp256k1VerificationKey2019"
	VM.Expires = time.Now().AddDate(10, 0, 0).Format(time.RFC3339) //set to expire 10 years from now as a placeholder

	document.VerificationMethod = append(document.VerificationMethod, VM)
	document.Authentication = VM.ID

	//once blockchain is implemented, will also need to upload the document to the blockchain

	return document, nil
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

func AuthenticateDocumentSubject(document *models.DIDDocument, message, signature []byte) (bool, error) {
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
	case "EcdsaSecp256k1VerificationKey2019":
		_, pubData, err := multibase.Decode(authenticationMethod.MultibaseKey)
		if err != nil {
			return false, err
		}

		result, err := secp256k1.Verify(message, signature, pubData)
		if err != nil {
			return false, err
		}
		return result, nil

	default:
		return false, errors.New("Unable to resolve unknown verification method type '" + authenticationMethod.MethodType + "'")
	}
}

func MetabloxRead(identifier string, options *models.ResolutionOptions) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	//Once blockchain is implemented, this function will use the identifier to retrieve the matching document from the blockchain (if it exists).
	//Some mechanism should likely be in place to ensure the document was not modified during the transfer, ex. an encrypted hash.

	placeholderKey, _ := secp256k1.NewECDSAPrivateKey()
	placeholderDoc, _ := CreateDID(placeholderKey)
	placeholderDoc.ID = "did:metablox:sampleIssuer"
	return &models.ResolutionMetadata{}, placeholderDoc, nil
}

func MetabloxReadRepresentation(identifier string, options *models.RepresentationResolutionOptions) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	//Should be similar to MetabloxRead, but returns the document in a specific representation format.
	//Representation type is included in options and returned in resolution metadata
	return nil, nil, nil
}
