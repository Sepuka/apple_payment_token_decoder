package eccv1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

func generateSharedSecret(privateKey *ecdsa.PrivateKey, key *ecdsa.PublicKey) []byte {
	c := elliptic.P256()
	x, _ := c.ScalarMult(key.X, key.Y, privateKey.D.Bytes())

	return x.Bytes()
}
