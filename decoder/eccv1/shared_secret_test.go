package eccv1

import (
	"fmt"
	"testing"

	"github.com/sepuka/apple-pay-token-decoder/internal"
	"github.com/stretchr/testify/assert"
)

const expectedSharedKey = `16d99f9792dac26fc029ded192df9d2261fac090a0f1a5be2760de86afe068dc`

func TestGenerateSharedSecret(t *testing.T) {
	privateKey, err := internal.BuildEcPrivateKey("../../resources/private.pem")
	assert.NoError(t, err)

	publicKey, err := internal.BuildEcPublicKey("../../resources/pubkey.pem")
	assert.NoError(t, err)

	assert.Equal(t, expectedSharedKey, fmt.Sprintf("%x", generateSharedSecret(privateKey, publicKey)))
}
