package schema

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReqActivateUserPhone struct {
	UserPhoneCodeID int `json:"user_phone_code_id"`
}
