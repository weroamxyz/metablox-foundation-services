package models

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/common"
)

type RegisterDID struct {
	Did     string `json:"did"`
	Account string `json:"account"`
	SigV    uint8  `json:"sigV"`
	SigR    string `json:"sigR"`
	SigS    string `json:"sigS"`
}

func NewRegisterDID() *RegisterDID {
	return &RegisterDID{}
}

func CreateMinerInfo() *MinerInfo {
	return &MinerInfo{}
}

func (c *RegisterDID) Address() common.Address {
	return common.HexToAddress(c.Account)
}

func (c *RegisterDID) SigRBytes32() [32]byte {
	var t [32]byte
	r, _ := hexutil.Decode(c.SigR)
	copy(t[:], r[:32])
	return t
}

func (c *RegisterDID) SigSBytes32() [32]byte {
	var t [32]byte
	s, _ := hexutil.Decode(c.SigS)
	copy(t[:], s[:32])
	return t
}

func (c *RegisterDID) ToSigBytes() []byte {
	var signBytes []byte
	r, _ := hexutil.Decode(c.SigR)
	s, _ := hexutil.Decode(c.SigS)
	signBytes = bytes.Join([][]byte{r, s}, nil)
	signBytes = append(signBytes, byte(c.SigV-27))
	return signBytes
}

type WifiUserInfo struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"value"`
}
