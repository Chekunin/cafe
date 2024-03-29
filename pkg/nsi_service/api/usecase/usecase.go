package usecase

import (
	"cafe/pkg/models"
	"cafe/pkg/nsi"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

type Usecase struct {
	nsi nsi.INSI
}

func NewUsecase(nsi nsi.INSI) *Usecase {
	return &Usecase{nsi: nsi}
}

func (u *Usecase) GetPlaceByID(ctx context.Context, id int) (models.Place, error) {
	place, err := u.nsi.GetPlaceByID(ctx, id)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceByID id=%d", id), err)
		return models.Place{}, err
	}
	return place, nil
}

func (u *Usecase) GetPlacesInsideBound(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error) {
	places, err := u.nsi.GetPlacesInsideBound(ctx, leftLng, rightLng, topLat, bottomLat)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlacesInsideBound leftLng=%f, rightLng=%f, topLat=%f, bottomLat=%f", leftLng, rightLng, topLat, bottomLat), err)
		return nil, err
	}
	return places, nil
}

func (u *Usecase) GetUserSubscriptionsByFollowerID(ctx context.Context, followerID int) ([]models.UserSubscription, error) {
	userSubscriptions, err := u.nsi.GetUserSubscriptionsByFollowerID(ctx, followerID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetUserSubscriptionsByFollowerID followerID=%d", followerID), err)
		return nil, err
	}
	return userSubscriptions, nil
}

func (u *Usecase) GetPlaceEvaluationByUserIDByPlaceID(ctx context.Context, userID int, placeID int) (models.PlaceEvaluation, error) {
	placeEvaluation, err := u.nsi.GetPlaceEvaluationByUserIDByPlaceID(ctx, userID, placeID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationByUserIDByPlaceID userID=%d placeID=%d", userID, placeID), err)
		return models.PlaceEvaluation{}, err
	}
	return placeEvaluation, nil
}

func (u *Usecase) GetPlaceEvaluationMarksByPlaceEvaluationID(ctx context.Context, placeEvaluationID int) ([]models.PlaceEvaluationMark, error) {
	placeEvaluationMarks, err := u.nsi.GetPlaceEvaluationMarksByPlaceEvaluationID(ctx, placeEvaluationID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceEvaluationMarksByPlaceEvaluationID placeEvaluationID=%d", placeEvaluationID), err)
		return nil, err
	}
	return placeEvaluationMarks, nil
}

func (u *Usecase) GetReviewMediaByID(ctx context.Context, reviewMediaID int) (models.ReviewMedia, error) {
	reviewMedia, err := u.nsi.GetReviewMediaByID(ctx, reviewMediaID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetReviewMediaByID reviewMediaID=%d", reviewMediaID), err)
		return models.ReviewMedia{}, err
	}
	return reviewMedia, nil
}

func (u *Usecase) GetPlaceEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	evaluationCriterions, err := u.nsi.GetPlaceEvaluationCriterions(ctx)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi evaluationCriterions"), err)
		return nil, err
	}
	return evaluationCriterions, nil
}

func (u *Usecase) GetReviewsByUserID(ctx context.Context, userID int) ([]models.Review, error) {
	reviews, err := u.nsi.GetReviewsByUserID(ctx, userID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetReviewsByUserID"), err)
		return nil, err
	}
	return reviews, nil
}
