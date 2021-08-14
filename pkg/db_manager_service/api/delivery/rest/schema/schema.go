package schema

import "cafe/pkg/models"

type ReqLogin struct {
	Username string
	Password string
}

type ReqActivateUserPhone struct {
	UserPhoneCodeID int
}

type ReqEvaluatePlace struct {
	PlaceEvaluation models.PlaceEvaluation
	Marks           []models.PlaceEvaluationMark
}
