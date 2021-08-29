package models

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type ReqLogin struct {
	Email    string `json:"email"`
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
	RestaurateurID int  `json:"restaurateur_id"`
	HasAccess      bool `json:"has_access"`
}
