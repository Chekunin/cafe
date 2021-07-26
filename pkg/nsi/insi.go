package nsi

import (
	"cafe/pkg/models"
	"context"
)

type INSI interface {
	GetPlacesInsideBound(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error)
	GetPlaceByID(ctx context.Context, id int) (models.Place, error)
}
