package rest

import (
	errs "cafe/pkg/client_sso/errors"
	"cafe/pkg/common"
	"errors"
)

func GetHttpCode(e error) int {
	var err common.Err
	if errors.As(e, &err) {
		if code, has := errMappingToHttpCode[err.Code]; has {
			return code
		}
	}
	return 500
}

// todo: заполнить
var errMappingToHttpCode = map[int]int{
	errs.ErrIncorrectLoginOrPassword.Code: 400,
}
