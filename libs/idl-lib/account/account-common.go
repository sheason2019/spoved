package account

type GetUsernameRepeatPayload struct {
	Name string `json:"name" form:"name"`
}

type GetSshPubKeyResponse struct {
	PubKey string `json:"pubKey" form:"pubKey"`
}

type User struct {
	Username string `json:"username" form:"username"`
}

type AccountInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Salt     string `json:"salt" form:"salt"`
}

type AccountCrypto struct {
	Salt   string `json:"salt" form:"salt"`
	PubKey string `json:"pubKey" form:"pubKey"`
}

type LoginResponse struct {
	Token string `json:"token" form:"token"`
}

type GetUsernameRepeatResponse struct {
	Repeat bool `json:"repeat" form:"repeat"`
}
