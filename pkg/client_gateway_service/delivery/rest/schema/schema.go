package schema

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ReqRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type RespRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ReqSignUp struct {
	Name     string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ApprovePhone struct {
	UserID int    `json:"user_id"`
	Phone  string `json:"phone"`
	Code   string `json:"code"`
}
