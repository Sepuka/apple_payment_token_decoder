package domain

// see the Payment token format there
// https://developer.apple.com/library/archive/documentation/PassKit/Reference/PaymentTokenJSON/PaymentTokenJSON.html

type Header struct {
	TransactionId string `json:"transactionId"`
	EphemeralPublicKey string `json:"ephemeralPublicKey"`
	PublicKeyHash string `json:"publicKeyHash"`
}

type PKPaymentToken struct {
	Version string `json:"version"`
	Data string `json:"data"`
	Signature string `json:"signature"`
	Header Header `json:"header"`
}
