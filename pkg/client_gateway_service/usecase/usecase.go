package usecase

import (
	"cafe/pkg/client_sso"
	"cafe/pkg/db_manager"
	"cafe/pkg/models"
	"cafe/pkg/nsi"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"math/rand"
	"time"
)

type Usecase struct {
	dbManager db_manager.IDBManager
	nsi       nsi.INSI
	clientSso client_sso.IClientSso
}

type NewUsecaseParams struct {
	DbManager db_manager.IDBManager
	Nsi       nsi.INSI
	ClientSso client_sso.IClientSso
}

func NewUsecase(params NewUsecaseParams) *Usecase {
	rand.Seed(time.Now().UnixNano())

	return &Usecase{
		dbManager: params.DbManager,
		nsi:       params.Nsi,
		clientSso: params.ClientSso,
	}
}

func (u *Usecase) GetPlaces(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error) {
	places, err := u.nsi.GetPlacesInsideBound(ctx, leftLng, rightLng, topLat, bottomLat)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlacesInsideBound leftLng=%f, rightLng=%f, topLat=%f, bottomLat=%f", leftLng, rightLng, topLat, bottomLat), err)
		return nil, err
	}

	return places, nil
}

func (u *Usecase) GetPlaceByID(ctx context.Context, placeID int) (models.Place, error) {
	place, err := u.nsi.GetPlaceByID(ctx, placeID)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("nsi GetPlaceByID id=%d", placeID), err)
		return place, err
	}

	return place, nil
}
