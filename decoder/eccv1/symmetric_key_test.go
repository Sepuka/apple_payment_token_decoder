package eccv1

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/sepuka/apple-pay-token-decoder/internal"
	"github.com/stretchr/testify/assert"
)

const expectedSymmetricKey = `af0abdeff68f12ae489faa55a6d66d3b00ea7b48df9bfd70d8228e7b8d04348f`

func TestGenerateSymmetricKey(t *testing.T) {
	privateKey, err := internal.BuildEcPrivateKey("../../resources/private.pem")
	assert.NoError(t, err)

	publicKey, err := internal.BuildEcPublicKey("../../resources/pubkey.pem")
	assert.NoError(t, err)

	cert, err := ioutil.ReadFile("../../resources/certificate.pem")
	if err != nil {
		assert.NoError(t, err)
	}

	merchantId, err := fetchMerchantId(cert)
	if err != nil {
		assert.NoError(t, err)
	}

	sharedKey := generateSharedSecret(privateKey, publicKey)

	symmetricKey, err := generateSymmetricKey(merchantId, sharedKey)
	assert.NoError(t, err)

	assert.Equal(t, expectedSymmetricKey, fmt.Sprintf("%x", symmetricKey))
}
