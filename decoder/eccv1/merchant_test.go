package eccv1

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegative(t *testing.T) {
	var testCases = map[string]struct{
		certificate []byte
		err error
	}{
		"empty certificate": {
			[]byte{},
			errors.New("failed to parse certificate PEM"),
		},
		"invalid certificate": {
			[]byte(`-----BEGIN CERTIFICATE-----
blah blah blah
-----END CERTIFICATE-----`),
			errors.New(`asn1: structure error: tags don't match (16 vs {class:1 tag:14 length:86 isCompound:true}) {optional:false explicit:false application:false private:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} certificate @2`),
		},
	}

	for testName, testCase := range testCases {
		_, err := fetchMerchantId(testCase.certificate)
		assert.EqualError(t, testCase.err, err.Error(), fmt.Sprintf("test '%s' faield", testName))
	}
}
