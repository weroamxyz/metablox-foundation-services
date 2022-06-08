package regutil

import "regexp"

const (
	RegETHAddress = "^0[xX][0-9a-zA-Z]{40}$"
)

func IsETHAddress(address string) bool {
	if address == "" {
		return false
	}
	flag, err := regexp.MatchString(RegETHAddress, address)
	return err == nil && flag
}
