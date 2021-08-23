package rest

import (
	"cafe/pkg/common"
	errs "cafe/pkg/nsi/errors"
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
	errs.ErrorEntityNotFound.Code: 400,
}
