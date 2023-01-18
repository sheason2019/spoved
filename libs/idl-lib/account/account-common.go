package account

type AccountInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type AccountCrypto struct {
	Salt   string `json:"salt"`
	PubKey string `json:"pubKey"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
