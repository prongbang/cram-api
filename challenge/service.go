package challenge

import (
	"fmt"

	"github.com/prongbang/cram-api/endcode"
	"github.com/prongbang/cram-api/signature"
)

type Service interface {
	Registration(r Registration) bool
	Request(r Request) string
	Verify(v Verify) bool
}

type service struct {
}

// Verify implements Service
func (s *service) Verify(v Verify) bool {
	pkBase64 := RegistrationMap[v.UserID].(string)
	challengeText := ChallengeMap[v.UserID].(string)

	decode := endcode.Decode{
		PublicKey: pkBase64,
	}
	pk := decode.ToECDSAPublicKey()

	textToSign := fmt.Sprintf("%s%s", v.Challenge, v.Nonce)
	sig := signature.Signature{
		PublicKey: pk,
		Message:   textToSign,
		Signature: v.Signature,
	}
	verified := sig.Verify()

	return verified && challengeText == v.Challenge
}

// Request implements Service
func (s *service) Request(r Request) string {

	challengeText := "577cb5f5-34d2-43ea-8a41-c67cdbf80a4a"

	ChallengeMap[r.UserID] = challengeText

	return challengeText
}

// Registration implements Service
func (s *service) Registration(r Registration) bool {

	RegistrationMap[r.UserID] = r.Pk

	return true
}

func NewService() Service {
	return &service{}
}
