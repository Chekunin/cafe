package errors

import "cafe/pkg/common"

var ErrorEntityNotFound = common.NewErr(1, "Сущность не найдена", nil)
