package resources

import (
	"embed"
)

//go:embed *
var EmbedFS embed.FS

const WifiCAFileName = "ca.der"
const WifiServerFileName = "server.der"

func GetWifiCAFile() ([]byte, error) {
	res, err := EmbedFS.ReadFile(WifiCAFileName)
	return res, err
}

func GetWifiServerFile() ([]byte, error) {
	res, err := EmbedFS.ReadFile(WifiServerFileName)
	return res, err
}
