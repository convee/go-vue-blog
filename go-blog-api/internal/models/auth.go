package models

type AuthInfo struct {
	Name   string   `json:"name"`
	Avatar string   `json:"avatar"`
	Role   []string `json:"role"`
}

type TokenInfo struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Id     uint   `json:"id"`
}
type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
