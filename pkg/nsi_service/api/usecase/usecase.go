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
