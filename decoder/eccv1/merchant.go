package eccv1

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
)

const extensionMerchantId = "1.2.840.113635.100.6.32"

func fetchMerchantId(data []byte) ([]byte, error) {
	var (
		roots        = x509.NewCertPool()
		block       *pem.Block
		certificate *x509.Certificate
		err         error
		extension   pkix.Extension
	)
	if ok := roots.AppendCertsFromPEM([]byte(rootPEM)); !ok {
		return nil, errors.New("cannot parse root certificate")
	}

	if block, _ = pem.Decode(data); block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}

	if certificate, err = x509.ParseCertificate(block.Bytes); err != nil {
		return nil, err
	}

	for _, extension = range certificate.Extensions {
		if extension.Id.String() == extensionMerchantId {
			return extension.Value[2:], nil
		}
	}

	return nil, errors.New("merchant identifier field OID 1.2.840.113635.100.6.32 is absent")
}
