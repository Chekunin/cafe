package errors

import "cafe/pkg/common"

var ErrorEntityNotFound = common.NewErr(1, "Сущность не найдена", nil)
var ErrorPhoneNotCorrect = common.NewErr(2, "Указан неверный номер телефона", nil)
var ErrorPhoneCodeNotCorrect = common.NewErr(3, "Неправильный код подтверждения номера телефона", nil)
var ErrorFieldNotUnique = common.NewErr(4, "Нарушена уникальность поля", nil)
