There is Apple payment token decoder implementation

Payment token structure described over there https://developer.apple.com/library/archive/documentation/PassKit/Reference/PaymentTokenJSON/PaymentTokenJSON.html

The library does not implement check token signature, it does only
1. fetch merchant id from OID 1.2.840.113635.100.6.32
1. parse public (ephemeral) key from token field which named "ephemeralPublicKey"
1. load private key and certificate (which you've signed it from apple developer service)
1. decoding token to json

###Tests
You can run up tests just call `go test ./...`, all need files inside the _resources_ folder already

###Links
implementation on some languages
* nodeJS https://github.com/sidimansourjs/applepay-token
* ruby https://github.com/spreedly/gala
* php https://github.com/PayU/apple-pay