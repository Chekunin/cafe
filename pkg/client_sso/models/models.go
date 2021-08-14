package models

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type ReqLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ReqLogout struct {
	Token string `json:"token"`
}

type ReqRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type ReqCheckPermission struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Token  string `json:"token"`
}

type RespCheckPermission struct {
	UserID    int  `json:"user_id"`
	HasAccess bool `json:"has_access"`
}
