package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sepuka/apple-pay-token-decoder/decoder"
	"github.com/sepuka/apple-pay-token-decoder/domain"
)

var pkpToken *domain.PKPaymentToken

func main() {
	file, err := ioutil.ReadFile("resources/token.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &pkpToken)
	if err != nil {
		panic(err)
	}

	cert, err := ioutil.ReadFile("resources/certificate.pem")
	if err != nil {
		panic(err)
	}

	privKey, err := ioutil.ReadFile("resources/flatpriv.pem")
	if err != nil {
		panic(err)
	}

	decoder1 := decoder.NewDecoder(pkpToken, privKey, cert)

	text, err := decoder1.Decode()

	if err == nil {
		fmt.Printf("decrypted text: %s", text)
	} else {
		fmt.Printf("occurred error %s while decoding token", err.Error())
	}
}
