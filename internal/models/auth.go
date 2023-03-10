package models

type AuthInfo struct {
	AdvertiserId uint64
	UserId       uint64
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
