package usecase

import (
	"cafe/pkg/models"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"time"
)

func (u *Usecase) EvaluatePlace(ctx context.Context, placeID int, userID int, marks []models.PlaceEvaluationMark, comment string) (models.PlaceEvaluation, []models.PlaceEvaluationMark, error) {
	placeEvaluation := models.PlaceEvaluation{
		PlaceID:  placeID,
		UserID:   userID,
		DateTime: time.Now(),
		Comment:  comment,
	}
	if err := u.dbManager.AddPlaceEvaluationWithMarks(ctx, &placeEvaluation, marks); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("dbManager AddPlaceEvaluationWithMarks"), err)
		return models.PlaceEvaluation{}, nil, err
	}

	// todo: отдавать в nsi

	return placeEvaluation, marks, nil
}

func (u *Usecase) GetPlaceEvaluation(ctx context.Context, placeID int, userID int) (models.PlaceEvaluation, []models.PlaceEvaluationMark, error) {
	placeEvaluation, err := u.nsi.GetPlaceEvaluationByUserIDByPlaceID(ctx, userID, placeID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationByUserIDByPlaceID userID=%d, placeID=%d", userID, placeID), err)
		return models.PlaceEvaluation{}, nil, err
	}

	marks, err := u.nsi.GetPlaceEvaluationMarksByPlaceEvaluationID(ctx, placeEvaluation.ID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationMarksByPlaceEvaluationID placeEvaluationID=%d", placeEvaluation.ID), err)
		return models.PlaceEvaluation{}, nil, err
	}

	return placeEvaluation, marks, nil
}
