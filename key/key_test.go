package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndLoadKey(t *testing.T) {
	savedPrivKey, fileName, err := GenerateNewPrivateKey()
	assert.Nil(t, err)

	loadedPrivKey, err := LoadPrivateKey(fileName)
	assert.Nil(t, err)
	assert.Equal(t, savedPrivKey, loadedPrivKey)
}
