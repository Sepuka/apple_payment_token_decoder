package eccv1

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	algoName = "\x0Did-aes256-GCMApple"
)

func generateSymmetricKey(merchantId []byte, sharedSecret []byte) ([]byte, error) {
	partV, _ := hex.DecodeString(fmt.Sprintf("%s", merchantId))

	f := append([]byte(algoName), partV...)

	key := sha256.New()
	if _, err := key.Write([]byte{0, 0, 0}); err != nil {
		return nil, err
	}
	if _, err := key.Write([]byte{1}); err != nil {
		return nil, err
	}
	if _, err := key.Write(sharedSecret); err != nil {
		return nil, err
	}
	if _, err := key.Write(f); err != nil {
		return nil, err
	}

	return key.Sum(nil), nil
}
