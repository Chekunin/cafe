package nsi

import (
	"cafe/pkg/models"
	"context"
)

type INSI interface {
	GetPlacesInsideBound(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error)
	GetPlaceByID(ctx context.Context, id int) (models.Place, error)
	GetUserSubscriptionsByFollowerID(ctx context.Context, followerID int) ([]models.UserSubscription, error)
	GetPlaceEvaluationByUserIDByPlaceID(ctx context.Context, userID int, placeID int) (models.PlaceEvaluation, error)
	GetPlaceEvaluationMarksByPlaceEvaluationID(ctx context.Context, placeEvaluationID int) ([]models.PlaceEvaluationMark, error)
}
