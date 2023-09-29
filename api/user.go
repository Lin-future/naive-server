package api

type SignInReq struct { //Body Params
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResp struct {
	Access_token string `json:"access_token"` // 身份验证token
}

type CheckInReq struct {
	Access_token string `json:"access_token"`
	Checkword    string `json:"checkword"`
}

type CheckInResp struct {
	Point int `json:"point"`
}

type SignUpReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpResp struct {
	Access_token string `json:"access_token"` // 身份验证token
}
