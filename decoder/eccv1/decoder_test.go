package eccv1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sepuka/apple-pay-token-decoder/domain"
	"github.com/stretchr/testify/assert"
)

const jsonToken = `{
  "version":"EC_v1",
  "data":"8FXb9Qozbg2NMTSABJJykjFayKY+5JZR+DWwM3m+7jyoRt4S+s+K9TIGbJzPe5JfQ2gQ3ilSRqx3gwskT0/NDVuS22blk9x8OkZ2FrxZHQ7/RGwHjdUCgN0i2PIPd6kJBebgl86EvvMyBhw+ZIP7kTD+2fuGmF0flaNwZYRU+oMqamJFpLyL9/dw4L0zjpjBPnIin5CBmXU5KGJN+cqrvx1jrtuKAanStx4xVWVsTvBfk72eyOgWuRqiECIkhaKKAXzoaYrb9HoC0IWkAnLAs4zjGYQxgJKUotaFZ20Ul65DY343cZMJJl9cN62YD2mOQcltyY5tapHyUjyMjVfih9YJ5MvbzHUZ4+MEmZ7ywUU+9QLmYVs4L3ZaDq/wQRmysa0t6qIJP6CUqfE/R2k10/7VehFa1E3QZIUTc0zo8uvH",
  "signature":"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",
  "header": {
    "ephemeralPublicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnttINJoMfcj5IKjJDXsVGYUPa/jPAOZy6nkPYLFZo2+p44Lsax9y4pL0hsQJF1E+5bLnwqWDluyEFY1LLZl42Q==",
    "publicKeyHash":"LD7zsYjsn8KMnLDp6nrrbrysw2Oe0De76yUwD7yykQc=",
    "transactionId":"1ee29d8d26982085b5df05e975a3455de3e52e8663ec59a75c505f73ac6936ed"
  }
}`

type PaymentData struct {
	Cryptogram string `json:"onlinePaymentCryptogram"`
	Eci        string `json:"eciIndicator"`
}

type ExpectedDecryptedToken struct {
	Account      string      `json:"applicationPrimaryAccountNumber"`
	ExpDate      string      `json:"applicationExpirationDate"`
	CurrencyCode string      `json:"currencyCode"`
	Amount       int         `json:"transactionAmount"`
	DeviceId     string      `json:"deviceManufacturerIdentifier"`
	Type         string      `json:"paymentDataType"`
	PaymentData  PaymentData `json:"paymentData"`
}

func TestDecoder(t *testing.T) {
	var (
		expectedToken = &ExpectedDecryptedToken{
			Account:      "4817499130197105",
			ExpDate:      "231231",
			CurrencyCode: "840",
			Amount:       25000,
			DeviceId:     "040010030273",
			Type:         "3DSecure",
			PaymentData: PaymentData{
				Cryptogram: "AoBXOn8AB9NUEIbHc+iTMAABAAA=",
				Eci:        "7",
			},
		}
		buffer                = bytes.NewBufferString(jsonToken)
		encryptedPaymentToken = &domain.PKPaymentToken{}
	)

	dir, err := os.Getwd()
	assert.NoError(t, err, "cannot detect currency directory")

	privateKeyPath := fmt.Sprintf("%s/%s", dir, `../../resources/private.pem`)
	privateKeyContent, err := ioutil.ReadFile(privateKeyPath)
	assert.NoError(t, err, "cannot read private key from file")

	err = json.Unmarshal(buffer.Bytes(), encryptedPaymentToken)
	assert.NoError(t, err, "cannot unmarshal encrypted payment token")

	certificatePath := fmt.Sprintf("%s/%s", dir, `../../resources/certificate.pem`)
	crt, err := ioutil.ReadFile(certificatePath)
	assert.NoError(t, err, "cannot read certificate from file")

	decoder := NewTokenDecoder(encryptedPaymentToken, privateKeyContent, crt)
	text, err := decoder.Decode()
	assert.NoError(t, err)

	jsonToken := &ExpectedDecryptedToken{}
	err = json.Unmarshal(text.([]byte), jsonToken)
	assert.NoError(t, err, "cannot unmarshal decrypted payment token")

	assert.Equal(t, expectedToken, jsonToken)
}
