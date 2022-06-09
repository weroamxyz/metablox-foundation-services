package contract

import (
	"bytes"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var privateKey, _ = crypto.HexToECDSA("9fd8f6049129527c63830aead266bcf7c53aa82109422da9335aac0c0a36a968")

func TestRegisterDID(t *testing.T) {

	register := models.NewRegisterDID()
	register.Did = "did:metablox:123456"
	register.Account = crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	userAddress := common.HexToAddress(register.Account)

	var messageBytes []byte
	messageBytes = bytes.Join([][]byte{messageBytes, []byte(register.Did), userAddress.Bytes(), []byte(big.NewInt(1).String()) /*nonceBytes[:]*/, []byte("register")}, nil)
	messageHash := crypto.Keccak256Hash(messageBytes)
	comboHash := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n32"), messageHash.Bytes())
	signature, err := crypto.Sign(comboHash[:], privateKey)
	assert.NoError(t, err)
	var r [32]byte
	var s [32]byte
	var v uint8
	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v = signature[64] + 27 //have to increment this manually as the smart contract expects v to be 27 or 28, while the crypto package generates it as 0 or 1

	register.SigR = hexutil.Encode(r[:])
	register.SigS = hexutil.Encode(s[:])
	register.SigV = v

	rr, err := hexutil.Decode(register.SigR)
	ss, err := hexutil.Decode(register.SigS)
	// check user signature
	var signBytes []byte
	signBytes = bytes.Join([][]byte{rr, ss}, nil)
	signBytes = append(signBytes, byte(register.SigV-27))

	userPub, err := crypto.SigToPub(comboHash.Bytes(), signBytes)
	assert.NoError(t, err)
	assert.Equal(t, userAddress, crypto.PubkeyToAddress(*userPub))

}
