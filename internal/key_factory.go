package internal

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func BuildEcPublicKey(keyPath string) (*ecdsa.PublicKey, error) {
	key, err := BuildPublicKey(keyPath)

	if err != nil {
		return nil, err
	}

	return key.(*ecdsa.PublicKey), err
}

func BuildPublicKey(publicKeyPath string) (interface{}, error) {
	publicKeyContent, err := ioutil.ReadFile(publicKeyPath)

	if err != nil {
		return nil, err
	}

	pemKey, _ := pem.Decode(publicKeyContent)
	return x509.ParsePKIXPublicKey(pemKey.Bytes)
}

func BuildEcPrivateKey(keyPath string) (*ecdsa.PrivateKey, error) {
	key, err := BuildPrivateKey(keyPath)

	if err != nil {
		return nil, err
	}

	return key.(*ecdsa.PrivateKey), err
}

func BuildPrivateKey(privateKeyPath string) (interface{}, error) {
	privateKeyContent, err := ioutil.ReadFile(privateKeyPath)

	if err != nil {
		return nil, err
	}

	pemKey, _ := pem.Decode(privateKeyContent)

	return x509.ParsePKCS8PrivateKey(pemKey.Bytes)
}
