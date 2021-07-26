package errors

import "cafe/pkg/common"

var (
	UnknownError                = common.NewErr(0, "Unknown error", nil)
	ErrIncorrectLoginOrPassword = common.NewErr(1, "Неправильный логин или пароль", nil)
	ErrAuthorizationHeader      = common.NewErr(2, "Формат заголовка авторизации должен быть Bearer {token}", nil)
	ErrIncorrectToken           = common.NewErr(3, "Неправильный токен", nil)
	ErrIncorrectRequest         = common.NewErr(400, "Неправильная ошибка", nil)
	ErrInternalServerError      = common.NewErr(500, "Внутренняя ошибка сервера", nil)
)

var codeToError = map[int]error{
	1:   ErrIncorrectLoginOrPassword,
	2:   ErrAuthorizationHeader,
	400: ErrIncorrectRequest,
	500: ErrInternalServerError,
}

func GetErrByCode(code int) error {
	if err, has := codeToError[code]; has {
		return err
	}
	return UnknownError
}
