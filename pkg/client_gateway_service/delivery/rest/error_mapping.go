package rest

import (
	errs "cafe/pkg/client_gateway_service/errors"
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

var errMappingToHttpCode = map[int]int{
	errs.ErrorEntityNotFound.Code:      400,
	errs.ErrorPhoneNotCorrect.Code:     400,
	errs.ErrorPhoneCodeNotCorrect.Code: 400,
	errs.ErrorFieldNotUnique.Code:      400,
	errs.ErrorForbidden.Code:           403,
}
