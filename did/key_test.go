package did

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func DeleteTestKeyFile() {
	saveLocation := viper.GetString("storage.key_saving")
	os.Remove(saveLocation + "testKey")
}
func TestSaveAndLoadKey(t *testing.T) {
	t.Cleanup(DeleteTestKeyFile)
	savedPrivKey, err := GenerateNewPrivateKey("testKey")
	assert.Nil(t, err)

	loadedPrivKey, err := LoadPrivateKey("testKey")
	assert.Nil(t, err)
	assert.Equal(t, savedPrivKey, loadedPrivKey)
}
