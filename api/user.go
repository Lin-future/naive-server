package api

type SignInReq struct { //Body Params
	Username string
	Password string
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
	Username string
	Password string
}

type SignUpResp struct {
	Access_token string `json:"access_token"` // 身份验证token
}
