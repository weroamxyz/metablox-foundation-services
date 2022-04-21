package did

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metabloxDID/contract"
	"github.com/metabloxDID/key"
	"github.com/metabloxDID/models"
	"github.com/mr-tron/base58"
	"github.com/multiformats/go-multibase"
)

func GenerateDIDString(privKey *ecdsa.PrivateKey) string {
	privData := crypto.FromECDSA(privKey)

	hash := sha256.New()
	hash.Write(privData)
	hashData := hash.Sum(nil)
	didString := base58.Encode(hashData)
	returnString := "did:metablox:" + didString
	return returnString
}

func CreateDID(privKey *ecdsa.PrivateKey) (*models.DIDDocument, error) {

	document := new(models.DIDDocument)

	var err error
	document.ID = GenerateDIDString(privKey)
	if err != nil {
		return nil, err
	}
	document.Context = make([]string, 0)
	document.Context = append(document.Context, "https://w3id.org/did/v1")
	document.Context = append(document.Context, "https://ns.did.ai/suites/secp256k1-2019/v1/")
	document.Created = time.Now().Format(time.RFC3339)
	document.Updated = document.Created
	document.Version = 1

	pubData := crypto.FromECDSAPub(&privKey.PublicKey)

	VM := models.VerificationMethod{}
	VM.ID = document.ID + "#verification"
	VM.MultibaseKey, err = multibase.Encode(multibase.Base58BTC, pubData)
	if err != nil {
		return nil, err
	}
	VM.Controller = document.ID
	VM.MethodType = "EcdsaSecp256k1VerificationKey2019"

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
func IsDIDValid(did []string) bool {
	if len(did) != 3 {
		fmt.Println("Not exactly 3 sections in DID")
		return false
	}
	prefix := did[0]
	if prefix != "did" {
		fmt.Println("First section of DID was '" + prefix + "' instead of 'did'")
		return false
	}

	methodName := did[1]
	if methodName != "metablox" {
		fmt.Println("Second section of DID was '" + methodName + "'instead of 'metablox'")
		return false
	}

	identifierPattern := `^([a-zA-Z0-9\._\-]*(%[0-9A-Fa-f][0-9A-Fa-f])*)*$`
	identifierExp, _ := regexp.Compile(identifierPattern)

	identifierSection := did[2]
	match := identifierExp.MatchString(identifierSection)
	if !match {
		fmt.Println("Identifier section is formatted incorrectly")
		return false
	}

	if len(identifierSection) == 0 {
		fmt.Println("Identifier is empty")
		return false
	}

	return true
}

func SplitDIDString(did string) []string {
	return strings.Split(did, ":")
}

func PrepareDID(did string) ([]string, bool) {
	splitString := SplitDIDString(did)
	valid := IsDIDValid(splitString)
	return splitString, valid
}

func Resolve(did string, options *models.ResolutionOptions) (*models.ResolutionMetadata, *models.DIDDocument, *models.DocumentMetadata) {
	splitDID, valid := PrepareDID(did)
	if !valid {
		return &models.ResolutionMetadata{Error: "invalid Did"}, nil, &models.DocumentMetadata{}
	}

	generatedDocument, generatedHash, err := contract.GetDocument(did)
	if err != nil {
		return &models.ResolutionMetadata{Error: err.Error()}, nil, nil
	}

	docID, success := PrepareDID(generatedDocument.ID)
	if !success {
		return &models.ResolutionMetadata{Error: "document DID is invalid"}, nil, nil
	}

	if docID[2] != splitDID[2] {
		return &models.ResolutionMetadata{Error: "generated document DID does not match provided DID"}, nil, nil
	}

	comparisonHash := sha256.Sum256(ConvertDocToBytes(*generatedDocument))
	if comparisonHash != generatedHash {
		return &models.ResolutionMetadata{Error: "document failed hash check"}, nil, nil
	}
	return &models.ResolutionMetadata{}, generatedDocument, nil
}

func ResolveRepresentation(did string, options *models.RepresentationResolutionOptions) (*models.RepresentationResolutionMetadata, []byte, *models.DocumentMetadata) {
	//Should be similar to Resolve, but returns the document in a specific representation format.
	//Representation type is included in options and returned in resolution metadata
	readOptions := models.CreateResolutionOptions()
	readResolutionMeta, document, readDocumentMeta := Resolve(did, readOptions)
	if readResolutionMeta.Error != "" {
		return &models.RepresentationResolutionMetadata{Error: readResolutionMeta.Error}, nil, nil
	}

	switch options.Accept {
	case "application/did+json":
		fallthrough
	default: //default to JSON format if options.Accept is empty/invalid
		byteStream, err := json.Marshal(document)
		if err != nil {
			return &models.RepresentationResolutionMetadata{Error: "failed to convert document into JSON"}, nil, nil
		}
		return &models.RepresentationResolutionMetadata{ContentType: "application/did+json"}, byteStream, readDocumentMeta
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

		result := crypto.VerifySignature(pubData, message, signature)
		return result, nil

	default:
		return false, errors.New("Unable to resolve unknown verification method type '" + authenticationMethod.MethodType + "'")
	}
}

func AuthenticateDocumentHolder(doc *models.DIDDocument, signature, nonce string) (bool, error) {
	authenticationMethod, err := doc.RetrieveVerificationMethod(doc.Authentication)
	if err != nil {
		return false, err
	}

	switch authenticationMethod.MethodType {
	case "EcdsaSecp256k1VerificationKey2019":
		return AuthenticateSecp256k1(signature, nonce, authenticationMethod)
	default:
		return false, errors.New("unable to verify unknown proof type " + authenticationMethod.MethodType)
	}
}

func AuthenticateSecp256k1(signature, nonce string, vm models.VerificationMethod) (bool, error) {
	_, pubData, err := multibase.Decode(vm.MultibaseKey)
	if err != nil {
		return false, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubData)
	if err != nil {
		return false, err
	}
	return key.VerifyJWSSignature(signature, pubKey, []byte(nonce))
}

func ConvertDocToBytes(doc models.DIDDocument) []byte {
	var convertedBytes []byte
	for _, item := range doc.Context {
		convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(doc.ID), []byte(doc.Created), []byte(doc.Updated), []byte(strconv.Itoa(doc.Version))}, []byte{})
	for _, item := range doc.VerificationMethod {
		convertedBytes = bytes.Join([][]byte{convertedBytes, ConvertVMToBytes(item)}, []byte{})
	}

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(doc.Authentication)}, []byte{})

	for _, item := range doc.Service {
		convertedBytes = bytes.Join([][]byte{convertedBytes, ConvertServiceToBytes(item)}, []byte{})
	}
	return convertedBytes
}

func ConvertVMToBytes(vm models.VerificationMethod) []byte {
	var convertedBytes []byte

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(vm.ID), []byte(vm.MethodType), []byte(vm.Controller), []byte(vm.MultibaseKey)}, []byte{})
	return convertedBytes
}

func ConvertServiceToBytes(service models.Service) []byte {
	var convertedBytes []byte

	convertedBytes = bytes.Join([][]byte{convertedBytes, []byte(service.ID), []byte(service.Type), []byte(service.ServiceEndpoint)}, []byte{})
	return convertedBytes
}
