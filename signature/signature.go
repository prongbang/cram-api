package signature

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
)

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

	return ecdsa.VerifyASN1(s.PublicKey, hash[:], signatureBytes)
}
