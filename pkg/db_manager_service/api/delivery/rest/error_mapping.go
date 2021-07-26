package rest

import (
	"cafe/pkg/common"
	errs "cafe/pkg/db_manager/errors"
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
