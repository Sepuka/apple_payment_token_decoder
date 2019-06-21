package eccv1

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sepuka/apple-pay-token-decoder/domain"
)

const (
	tagLength             = 16
	keyFormat             = "%s\n%s\n%s"
	publicKeyBeginMarker  = `-----BEGIN PUBLIC KEY-----`
	publicKeyEndMarker    = `-----END PUBLIC KEY-----`
	privateKeyBeginMarker = `-----BEGIN PRIVATE KEY-----`
	privateKeyEndMarker   = `-----END PRIVATE KEY-----`

	rootPEM = `
-----BEGIN CERTIFICATE-----
MIICQzCCAcmgAwIBAgIILcX8iNLFS5UwCgYIKoZIzj0EAwMwZzEbMBkGA1UEAwwS
QXBwbGUgUm9vdCBDQSAtIEczMSYwJAYDVQQLDB1BcHBsZSBDZXJ0aWZpY2F0aW9u
IEF1dGhvcml0eTETMBEGA1UECgwKQXBwbGUgSW5jLjELMAkGA1UEBhMCVVMwHhcN
MTQwNDMwMTgxOTA2WhcNMzkwNDMwMTgxOTA2WjBnMRswGQYDVQQDDBJBcHBsZSBS
b290IENBIC0gRzMxJjAkBgNVBAsMHUFwcGxlIENlcnRpZmljYXRpb24gQXV0aG9y
aXR5MRMwEQYDVQQKDApBcHBsZSBJbmMuMQswCQYDVQQGEwJVUzB2MBAGByqGSM49
AgEGBSuBBAAiA2IABJjpLz1AcqTtkyJygRMc3RCV8cWjTnHcFBbZDuWmBSp3ZHtf
TjjTuxxEtX/1H7YyYl3J6YRbTzBPEVoA/VhYDKX1DyxNB0cTddqXl5dvMVztK517
IDvYuVTZXpmkOlEKMaNCMEAwHQYDVR0OBBYEFLuw3qFYM4iapIqZ3r6966/ayySr
MA8GA1UdEwEB/wQFMAMBAf8wDgYDVR0PAQH/BAQDAgEGMAoGCCqGSM49BAMDA2gA
MGUCMQCD6cHEFl4aXTQY2e3v9GwOAEZLuN+yRhHFD/3meoyhpmvOwgPUnPWTxnS4
at+qIxUCMG1mihDK1A3UT82NQz60imOlM27jbdoXt2QfyFMm+YhidDkLF1vLUagM
6BgD56KyKA==
-----END CERTIFICATE-----
`
)

type Decoder struct {
	data        string
	signature   string
	trxId       string
	publicKey   []byte
	keyHash     string
	privateKey  []byte
	certificate []byte
}

func NewTokenDecoder(
	token *domain.PKPaymentToken,
	privateKey []byte,
	certificate []byte,
) domain.Decoder {
	return &Decoder{
		data:        token.Data,
		signature:   token.Signature,
		trxId:       token.Header.TransactionId,
		publicKey:   []byte(token.Header.EphemeralPublicKey),
		keyHash:     token.Header.PublicKeyHash,
		privateKey:  privateKey,
		certificate: certificate,
	}
}

func (d *Decoder) Decode() (interface{}, error) {
	var (
		publicKey    *ecdsa.PublicKey
		privateKey   *ecdsa.PrivateKey
		merchantId   []byte
		sharedKey    []byte
		symmetricKey []byte
		err          error
		iv           = bytes.NewBuffer([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	)

	if merchantId, err = fetchMerchantId(d.certificate); err != nil {
		return nil, err
	}

	if publicKey, err = d.getPublicKey(); err != nil {
		return nil, err
	}

	if privateKey, err = d.getPrivateKey(); err != nil {
		return nil, err
	}

	sharedKey = generateSharedSecret(privateKey, publicKey)
	symmetricKey, err = generateSymmetricKey(merchantId, sharedKey)
	if err != nil {
		return nil, err
	}

	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(d.data))

	block, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil, err
	}

	text, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, tagLength)
	if err != nil {
		return nil, err
	}

	return aesgcm.Open(nil, iv.Bytes(), text, nil)
}

func (d *Decoder) getPublicKey() (*ecdsa.PublicKey, error) {
	key := fmt.Sprintf(keyFormat, publicKeyBeginMarker, d.publicKey, publicKeyEndMarker)
	publicKey, _ := pem.Decode([]byte(key))

	pub, err := x509.ParsePKIXPublicKey(publicKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("cannot build public key: %#v", err)
	}

	return pub.(*ecdsa.PublicKey), nil
}

func (d *Decoder) getPrivateKey() (*ecdsa.PrivateKey, error) {
	key := fmt.Sprintf(keyFormat, privateKeyBeginMarker, d.privateKey, privateKeyEndMarker)
	privateKey, _ := pem.Decode([]byte(key))

	private, err := x509.ParsePKCS8PrivateKey(privateKey.Bytes)
	if err != nil {
		return nil, fmt.Errorf("cannot build private key: %#v", err)
	}

	return private.(*ecdsa.PrivateKey), nil
}
