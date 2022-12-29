package challenge

type Registration struct {
	UserID string `json:"userId"`
	Pk     string `json:"pk"`
}

type Request struct {
	UserID string `json:"userId"`
}

type Verify struct {
	UserID    string `json:"userId"`
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

type Challenge struct {
	Challenge string `json:"challenge"`
}
