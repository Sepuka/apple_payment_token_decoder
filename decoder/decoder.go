package decoder

import (
	"github.com/sepuka/apple-pay-token-decoder/decoder/eccv1"
	"github.com/sepuka/apple-pay-token-decoder/decoder/rsav1"
	"github.com/sepuka/apple-pay-token-decoder/domain"
)

const (
	VersionEccV1 = "EC_v1"
	VersionRsaV1 = "RSA_v1"
)

func NewDecoder(token *domain.PKPaymentToken, privateKey []byte, certificate []byte) domain.Decoder {
	switch token.Version {
	case VersionEccV1:
		return eccv1.NewTokenDecoder(token, privateKey, certificate)
	case VersionRsaV1:
		return &rsav1.Decoder{}
	default:
		panic("unknown token version")
	}
}
