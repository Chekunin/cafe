package errors

import "cafe/pkg/common"

var MapErrorsByCode = map[int]common.Err{
	1: ErrorEntityNotFound,
	//1001: ErrIntegrityConstraintViolation,
	//1002: ErrRestrictViolation,
	//1003: ErrNotNullViolation,
	//1004: ErrForeignKeyViolation,
	//1005: ErrUniqueViolation,
	//1006: ErrCheckViolation,
	//1007: ErrExclusionViolation,
}

var ErrorEntityNotFound = common.NewErr(1, "Сущность не найдена", nil)
var ErrorEntityAlreadyExists = common.NewErr(2, "Сущность уже существует", nil)

// ошибки из БД
var ErrIntegrityConstraintViolation = func(meta interface{}) common.Err {
	return common.NewErr(1001, "Нарушение целостности данных", meta)
}
var ErrRestrictViolation = func(meta interface{}) common.Err {
	return common.NewErr(1002, "Restrict violation", meta)
}
var ErrNotNullViolation = func(meta interface{}) common.Err {
	return common.NewErr(1003, "Значение не может быть пустым", meta)
}
var ErrForeignKeyViolation = func(meta interface{}) common.Err {
	return common.NewErr(1004, "Указан несуществующий ключ", meta)
}
var ErrUniqueViolation = func(meta interface{}) common.Err {
	return common.NewErr(1005, "Значение должно быть уникальным", meta)
}
var ErrCheckViolation = func(meta interface{}) common.Err {
	return common.NewErr(1006, "Check violation", meta)
}
var ErrExclusionViolation = func(meta interface{}) common.Err {
	return common.NewErr(1007, "Exclusion violation", meta)
}
