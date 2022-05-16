# metablox-foundation-services
Server-side component of the Metablox DID system. The foundation service is responsible for uploading DID documents to the Metablox smart contract, as well as issuing, updating, and renewing verifiable credentials.
## Building from source
### Prerequisites
| Components | Version | Description |
|----------|-------------|-------------|
|[Golang](https://golang.org) | >= 1.16| The Go Programming Language |

### Build

1. Checkout repo.

```bash
cd $GOPATH/src
go get -u -v github.com/MetaBloxIO/metablox-foundation-services
```

2. Import dependencies and build.

```bash
cd github.com/MetaBloxIO/metablox-foundation-services
make build
# note: you may be asked to grant super user permission
```

## Running the Service
The executable, ```metabloxDID```, is located in the top-level directory. Run the following command in the top level of the repository to start running the backend.
``` bash
./metabloxDID
```
This should begin running the foundation service on port **8888**.

## Configuration
The file ```config.yaml``` (located in the top-level directory) has a number of configuration settings that should be set before running the program.

### storage
These settings determine where private keys are saved/loaded on the server.
| Setting | Default | Description |
|----------|-------------|-------------|
|key_saving|"./wallet"|The location where private keys should be saved.|
|key_loading|"./wallet"|The location where private keys should be loaded from.|
|issuer_key_file|"issuer"|The name that should be used for the issuer private key file.|

### mysql
These settings determine the parameters used when connecting to the mysql database that will be paired with the Metablox foundation service.
| Setting | Default | Description |
|----------|-------------|-------------|
|host|"127.0.0.1"|The address where the mysql server is being hosted.|
|port|3306|The port that the host is using for the mysql server.|
|user|"tester"|The username that should be used to access the mysql database.|
|password|"omnisolutesting"|The password that should be used to access the mysql database.|
|dbname|"foundationService"|The name of the database that should be opened on the mysql server.|
|testhost|"127.0.0.1"|**optional**: The address where the mysql test server is being hosted.|
|testport|3306|**optional**: The port that the testhost is using for the mysql test server.|
|testuser|"tester"| **optional**: The username that should be used to access the mysql test database.|
|testpassword|"omnisolutesting"|**optional**: The password that should be used to access the mysql test database.|
|testdbname|"foundationservicetest"|**optional**: The name of the database that should be opened on the mysql test server.|

## Database Information
The foundation service is designed to be run alongside a mysql database, which will store miner information as well as details about any credentials issued by the backend. For the foundation service to run properly, it is important that the database is formatted correctly.

### Credential
|Column Name|Type|
|------------|----|
|ID|int AI PK|
|Type|varchar(100)|
|Issuer|varchar(100)|
|IssuanceDate|timestamp|
|ExpirationDate|timestamp|
|Description|varchar(100)|
|Revoked|tinyint|

### MiningLicenseInfo
|Column Name|Type|
|------------|----|
|CredentialID|int PK|
|ID|varchar(100)|
|Name|varchar(100)|
|Model|varchar(100)|
|Serial|varchar(100)|

### WifiAccessInfo
|Column Name|Type|
|------------|----|
|CredentialID|int PK|
|ID|varchar(100)|
|Type|enum('User','Validator')|

### MinerInfo
|Column Name|Type|
|------------|----|
|ID|int AI PK|
|Name|varchar(100)|
|MAC|varchar(100)|
|CreateTime|timestamp|

### MinerManufacturer
|Column Name|Type|
|------------|----|
|Name|varchar(100) PK|
|Email|varchar(100)|
|Address|varchar(100)|

## Data Structures
There are a number of different data structures used by the foundation service, both as inputs and outputs. Examples and explanations of each have been provided below.

### DID Document
```
{
    "@context": [
        "https://w3id.org/did/v1",
        "https://ns.did.ai/suites/secp256k1-2019/v1/"
    ],
    "id": "HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo",
    "created": "2022-03-29T15:35:45-07:00",
    "updated": "2022-03-29T15:35:45-07:00",
    "version": 1,
    "verificationMethod": [
        {
            "id": "did:metablox:mnmbxakwgkwhj#verification",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "controller": "did:metablox:mnmbxakwgkwhj",
            "publicKeyMultibase": "zR4TQJaWaLA3vvYukULRJoxTsRmqCMsWuEJdDE8CJwRFCUijDGwCBP89xVcWdLRQaEM6b7wD294xCs8byy3CdDMYa"
        }
    ],
    "authentication": "did:metablox:mnmbxakwgkwhj#verification"
}
```
|Field Name|Description|
|-----------|-----------|
|@context|Array of strings used to retrieve JSON-LD contexts (see [here](https://json-ld.org/spec/latest/json-ld/#the-context) for more details). The first string must always be "https://w3id.org/did/v1", and "https://ns.did.ai/suites/secp256k1-2019/v1/" must also be present.
|id|DID associated with the document. The entity with the private key matching this document controls this DID.|
|created|Time of creation for the document.|
|updated|Time of last update for the document.|
|version|The document version number. If the model for DID documents is modified in the future, the version numbers of new documents will be increased.
|verificationMethod|Array of verification methods that can be used to interact with the DID document. See "Verification Method" section for more details on formatting.
|authentication|ID of verification method that can be used to authenticate the holder of this DID document.

### Verification Method
```
    {
            "id": "did:metablox:mnmbxakwgkwhj#verification",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "controller": "did:metablox:mnmbxakwgkwhj",
            "publicKeyMultibase": "zR4TQJaWaLA3vvYukULRJoxTsRmqCMsWuEJdDE8CJwRFCUijDGwCBP89xVcWdLRQaEM6b7wD294xCs8byy3CdDMYa"
    }
```
|Field Name|Description|
|-----------|-----------|
|id|ID of the verification method. Should always consist of the controller DID followed by '#*identifier*'|
|type|The type of verification method. Determines how signature verification should be performed. Currently, only "EcdsaSecp256k1VerificationKey2019" is supported.
|controller|DID that is associated with verification method. Completing the verification method proves a degree of ownership over this DID.
|publicKeyMultibase|Multibase public key created from the controller DID holder's private key. If a provided signature can be verified using this key, it proves that the signer knows the private key, and is therefore the holder of the controller DID.

### Verifiable Credential
```
{
        "@context": [
            "https://www.w3.org/2018/credentials/v1",
            "https://ns.did.ai/suites/secp256k1-2019/v1/"
        ],
        "id": "http://metablox.com/credentials/48",
        "type": [
            "VerifiableCredential",
            "MiningLicense"
        ],
        "issuer": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo",
        "issuanceDate": "2022-05-02T16:21:49-07:00",
        "expirationDate": "2032-05-02T16:21:49-07:00",
        "description": "Example Mining License Credential",
        "credentialSubject": {
            "id": "did:metablox:hgsduijgbwxxxxx",
            "name": "check2",
            "model": "check",
            "serial": "check3"
        },
        "proof": {
            "type": "EcdsaSecp256k1Signature2019",
            "created": "2022-05-02T16:21:49-07:00",
            "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
            "proofPurpose": "Authentication",
            "jws": "eyJhbGciOiJFUzI1NiJ9..6ixXRTXTNj6iOa0pKqExzXLaunTKWlNjCA_Ctm_EfOCb9Q6U6j6EaYKCUJTgTxfEHgh2gLDBUagzsWXedh2lag"
        },
        "revoked": false
    }
```
|Field Name|Description|
|-----------|-----------|
|@context|Array of strings used to retrieve JSON-LD contexts (see [here](https://json-ld.org/spec/latest/json-ld/#the-context) for more details). The first string must always be "https://www.w3.org/2018/credentials/v1", and "https://ns.did.ai/suites/secp256k1-2019/v1/" must also be present.|
|id|Identifier for the credential|
|type|Array of strings used to identify the type of credential. The type "VerifiableCredential" should always be present, alongside one or more other types.
|issuer|DID of the entity that issued the credential. This DID must be registered as a valid issuer in order for the credential to be considered valid.|
|issuanceDate|The date when the credential was issued.|
|expirationDate|The date when the credential will expire. If used after this date, the credential should be rejected.|
|description|Additional information provided about the credential.|
|credentialSubject|A json object providing information about the subject of the credential. The contents of this section will depend on the type of credential; see "Credential Subject (Wifi Access)" and "Credential Subject (Mining License)" sections for more details.
|proof|A json object containing the information necessary to verify that the listed issuer actually issued this credential. This is done by completing one of the issuer's verification methods. See "Credential Proof" section for more details.
|revoked|Indicates whether or not the document has been revoked, rendering it unusable.|

### Verifiable Presentation
```
       {
            "@context": [
            "https://www.w3.org/2018/credentials/v1",
            "https://ns.did.ai/suites/secp256k1-2019/v1/"
        ],
        "type": [
            "VerifiablePresentation"
        ],
        "verifiableCredential": [
            {
                "@context": [
                    "https://www.w3.org/2018/credentials/v1",
                    "https://ns.did.ai/suites/secp256k1-2019/v1/"
                ],
                "id": "http://metablox.com/credentials/27",
                "type": [
                    "VerifiableCredential",
                    "WifiAccess"
                ],
                "issuer": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo",
                "issuanceDate": "2022-04-14T12:30:43-07:00",
                "expirationDate": "2032-04-14T12:30:43-07:00",
                "description": "Example Wifi Access Credential",
                "credentialSubject": {
                    "id": "did:metablox:hgsduijgbwu",
                    "type": "User"
                }
            }
        ],
        "holder": "did:metablox:hgsduijgbwu",
        "proof": {
            "type": "EcdsaSecp256k1Signature2019",
            "created": "2022-03-31T12:53:19-07:00",
            "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
            "proofPurpose": "Authentication",
            "jws": "eyJhbGciOiJFUzI1NiJ9..bmj6KhHcBkLOHgAZrLqgweE-StyBXvvj6bmZqC6TqiYVtC_tXf076xDAAXzmx160dAqivTzgX-943ZU-VWXDqw",
            "nonce": "1651007184616875174"
        }
    }
```
|Field Name|Description|
|-----------|-----------|
|@context|Array of strings used to retrieve JSON-LD contexts (see [here](https://json-ld.org/spec/latest/json-ld/#the-context) for more details). The first string must always be "https://www.w3.org/2018/credentials/v1", and "https://ns.did.ai/suites/secp256k1-2019/v1/" must also be present.|
|type|Array of strings used to identify the type of presentation. The type "VerifiablePresentation" should always be present.
|verifiableCredential|Array of verifiable credentials that the presentation is presenting. See "Verifiable Credential" section for more details.
|holder|DID of the entity that created the presentation.|
|proof|A json object containing the information necessary to verify that the listed holder actually created this presentation. This is done by completing one of the holder's verification methods. See "Presentation Proof" section for more details.

### Credential Subject (Wifi Access)
```
{
      "id": "did:metablox:test",
      "type": "User"
}
```
|Field Name|Description|
|-----------|-----------|
|id|DID of the credential subject.|
|type|The type of user that the credential subject is. Can either be 'User' or 'Validator'.

### Credential Subject (Mining License)
```
{
            "id": "did:metablox:hgsduijgbwxxxxx",
            "name": "check2",
            "model": "check",
            "serial": "check3"
}
```
|Field Name|Description|
|-----------|-----------|
|id|DID of the credential subject.|
|name|Name of the miner manufacturer.|
|model|Name of the miner model.|
|serial|Serial number of the miner.|

### Credential Proof
```
{
            "type": "EcdsaSecp256k1Signature2019",
            "created": "2022-05-02T16:21:49-07:00",
            "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
            "proofPurpose": "Authentication",
            "jws": "eyJhbGciOiJFUzI1NiJ9..6ixXRTXTNj6iOa0pKqExzXLaunTKWlNjCA_Ctm_EfOCb9Q6U6j6EaYKCUJTgTxfEHgh2gLDBUagzsWXedh2lag"
}
```
|Field Name|Description|
|-----------|-----------|
|type|The type of verification method. Determines how signature verification should be performed. Currently, only "EcdsaSecp256k1Signature2019" is supported.
|created|The time of creation for the proof.|
|verificationMethod|The address of the verification method that matches this proof.
|proofPurpose|The purpose of the proof. If the proof is verified, the credential holder is authorized to perform this action. Must match the verification method.
|jws|The signature of the proof, in JWS format. Created using a hash of the credential and the issuer's private key. If verified using the issuer's verification method, then it means the credential was issued by the issuer and that it has not been altered.

### Presentation Proof
```
{
            "type": "EcdsaSecp256k1Signature2019",
            "created": "2022-03-31T12:53:19-07:00",
            "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
            "proofPurpose": "Authentication",
            "jws": "eyJhbGciOiJFUzI1NiJ9..bmj6KhHcBkLOHgAZrLqgweE-StyBXvvj6bmZqC6TqiYVtC_tXf076xDAAXzmx160dAqivTzgX-943ZU-VWXDqw",
            "nonce": "1651007184616875174"
}
```
|Field Name|Description|
|-----------|-----------|
|type|The type of verification method. Determines how signature verification should be performed. Currently, only "EcdsaSecp256k1Signature2019" is supported.
|created|The time of creation for the proof.|
|verificationMethod|The address of the verification method that matches this proof.
|proofPurpose|The purpose of the proof. If the proof is verified, the credential holder is authorized to perform this action. Must match the verification method.
|jws|The signature of the proof, in JWS format. Created using a hash of the credential and the holder's private key. If verified using the holder's verification method, then it means the presentation was created by the holder and that it has not been altered.
|nonce|An arbitrary value. The foundation service will require users to include specific random nonces in their presentations for them to be considered valid, with each nonce only being valid for a single operation. This prevents the same presentation from being used repeatedly to mitigate replay attacks.



## Uploading a DID Document
The POST API '/registry/storedoc' is used to upload a DID document to the Metablox smart contract. To use it, send a request to the url with the document you want to upload in the body of the request.

Example of uploading a document to a foundation service running on localhost (port 8888):
url: localhost:8888/registry/storedoc
body:
```
{
    "@context": [
        "https://w3id.org/did/v1",
        "https://ns.did.ai/suites/secp256k1-2019/v1/"
    ],
    "id": "HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo",
    "created": "2022-03-29T15:35:45-07:00",
    "updated": "2022-03-29T15:35:45-07:00",
    "version": 1,
    "verificationMethod": [
        {
            "id": "did:metablox:mnmbxakwgkwhj#verification",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "controller": "did:metablox:mnmbxakwgkwhj",
            "publicKeyMultibase": "zR4TQJaWaLA3vvYukULRJoxTsRmqCMsWuEJdDE8CJwRFCUijDGwCBP89xVcWdLRQaEM6b7wD294xCs8byy3CdDMYa"
        }
    ],
    "authentication": "did:metablox:mnmbxakwgkwhj#verification"
}
```
If successful, the output will be a success message. The smart contract should be updated within a short time with the new DID.

## Issuing a Credential
The POST APIs '/vc/wifi/issue/*:did*' and '/vc/mining/issue/*:did*' are used to issue new credentials to DID holders. To use them, simply replace *:did* in the url with the identifier (the section after "did:metablox:") of the did issuing the credential, and include a json object with information about the credential subject in the body of the request (see "Credential Subject (Wifi Access)" and "Credential Subject (Mining License)" in the "Data Structures" sections for the correct formats to use).

Example of issuing a wifi access credential to the DID "did:metablox:test" from a foundation service running on localhost (port 8888), assuming that "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo' is a valid issuer:
url: localhost:8888/vc/wifi/issue/HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo
body:
```
{
      "id": "did:metablox:test",
      "type": "User"
}
```
If successful, the output will be a verifiable credential created with the provided information.

## Renewing/Revoking Credentials
The POST APIs '/vc/wifi/renew/*:did*' and '/vc/mining/renew/*:did*' are both used to renew a credential. Meanwhile, the POST APIs '/vc/wifi/revoke/*:did*' and '/vc/mining/revoke/*:did*' are used to revoke a credential.
Both of these APIs are similar to use. They require a valid DID identifier to be provided in the url (replacing the '*:did*' portion), and a valid verifiable presentation in the body of the request. 
This presentation should include the credential being renewed/revoked as the first in the 'verifiableCredential' field; any additional credentials provided will be ignored. To avoid ambiguity, try to avoid including more than one credential here.

To make the presentation valid, it must include a valid nonce in its presentation proof. A nonce can be gotten by using the GET API '/nonce'; this will assign a nonce to the user's IP address for a short time and return it. This nonce can then be included in the user's presentation. Remember that each nonce is only valid for a single operation, so a new one will be needed if another operation is performed later. 

Example of getting a nonce from a foundation service running on localhost (port 8888):
url: localhost:8888/nonce

If successful, the output will be a nonce. This nonce will be stored internally, causing any other nonces to be treated as invalid for that IP address.

Example of renewing a wifi access credential with a foundation service running on localhost (port 8888):
url: localhost:8888/vc/wifi/renew/HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo
body:
```
{
        "@context": [
            "https://www.w3.org/2018/credentials/v1",
            "https://ns.did.ai/suites/secp256k1-2019/v1/"
        ],
        "type": [
            "VerifiablePresentation"
        ],
        "verifiableCredential": [
            {
                "@context": [
                    "https://www.w3.org/2018/credentials/v1",
                    "https://ns.did.ai/suites/secp256k1-2019/v1/"
                ],
                "id": "http://metablox.com/credentials/27",
                "type": [
                    "VerifiableCredential",
                    "WifiAccess"
                ],
                "issuer": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo",
                "issuanceDate": "2022-04-14T12:30:43-07:00",
                "expirationDate": "2032-04-14T12:30:43-07:00",
                "description": "Example Wifi Access Credential",
                "credentialSubject": {
                    "id": "did:metablox:hgsduijgbwu",
                    "placeholderParameter": "check"
                },
                "proof": {
                     "type": "EcdsaSecp256k1Signature2019",
                    "created": "2022-04-21T14:36:35-07:00",
                    "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
                    "proofPurpose": "Authentication",
                    "jws": "eyJhbGciOiJFUzI1NiJ9..uhWLPyNoILjCj_6HiFbMFu09xgo_YfRls6mDl3a7Qt4VSYfb4uSAaxbx12HZ4QA8J6gkdJHBB866JDqQ_o6B6Q"
                },
                "revoked": "false"
            }
        ],
        "holder": "did:metablox:hgsduijgbwu",
        "proof": {
            "type": "EcdsaSecp256k1Signature2019",
            "created": "2022-03-31T12:53:19-07:00",
            "verificationMethod": "did:metablox:HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo#verification",
            "proofPurpose": "Authentication",
            "jws": "eyJhbGciOiJFUzI1NiJ9..bmj6KhHcBkLOHgAZrLqgweE-StyBXvvj6bmZqC6TqiYVtC_tXf076xDAAXzmx160dAqivTzgX-943ZU-VWXDqw",
            "nonce": "2022-04-21 13:07:53.894053734 -0700 PDT"
        }
}
```
If successful, the output will be a new copy of the credential with an updated expiration date. This change will be reflected in the database.

If the url 'localhost:8888/vc/wifi/revoke/HFXPiudexfvsJBqABNmBp785YwaKGjo95kmDpBxhMMYo' is used with the same input, the credential will be revoked instead of renewed. The output in that case will be a copy of the credential that has had its 'revoked' field set to false. This change will be reflected in the database and smart contract.

