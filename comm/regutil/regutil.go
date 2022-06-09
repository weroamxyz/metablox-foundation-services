package regutil

import (
	"github.com/ethereum/go-ethereum/common"
)

//const (
//	RegETHAddress = "^0[xX][0-9a-zA-Z]{40}$"
//)

func IsETHAddress(address string) bool {
	if address == "" {
		return false
	}
	return common.IsHexAddress(address)
}
