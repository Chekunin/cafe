package rest

import (
	errs "cafe/pkg/client_gateway_service/errors"
	"cafe/pkg/common"
	"errors"
)

func GetHttpCode(e error) int {
	var err common.Err
	if errors.As(e, &err) {
		if code, has := errMappingToHttpCode[err]; has {
			return code
		}
	}
	return 500
}

var errMappingToHttpCode = map[error]int{
	errs.ErrorEntityNotFound: 400,
}
