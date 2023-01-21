package account

type GetUsernameRepeatPayload struct {
	Name string `json:"name"`
}

type User struct {
	Username string `json:"username"`
}

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

type GetUsernameRepeatResponse struct {
	Repeat bool `json:"repeat"`
}
