package resources

import (
	"embed"
)

//go:embed *
var EmbedFS embed.FS

const WifiCertFileName = "ca.der"

func GetWifiCert() ([]byte, error) {
	res, err := EmbedFS.ReadFile(WifiCertFileName)
	return res, err
}
