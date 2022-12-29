package endcode

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
)

type Decode struct {
	PublicKey string
}

func (d *Decode) ToECDSAPublicKey() *ecdsa.PublicKey {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(d.PublicKey)
	if err != nil {
		return nil
	}

	pub, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil
	}

	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil
	}

	return publicKey
}
