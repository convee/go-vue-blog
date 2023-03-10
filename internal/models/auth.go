package models

type AuthInfo struct {
	Username string
}

type TokenInfo struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Id       uint   `json:"id"`
}
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
