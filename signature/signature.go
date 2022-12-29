package signature

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

type Signature struct {
	PublicKey *ecdsa.PublicKey
	Message   string
	Signature string
}

func (s *Signature) Verify() bool {
	signatureBytes, err := base64.StdEncoding.DecodeString(s.Signature)
	if err != nil {
		return false
	}
	hash := sha256.Sum256([]byte(s.Message))

	ecdsaSignature := &ecdsaSignature{}
	if _, err = asn1.Unmarshal(signatureBytes, ecdsaSignature); err != nil {
		return false
	}

	return ecdsa.Verify(s.PublicKey, hash[:], ecdsaSignature.R, ecdsaSignature.S)
}
