package memory

import (
	"cafe/pkg/db_manager"
	"cafe/pkg/models"
	"cafe/pkg/nsi"
	errs "cafe/pkg/nsi/errors"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/kyroy/kdtree/kdrange"
	KDTreePoints "github.com/kyroy/kdtree/points"
)

type NSI struct {
	dbManager db_manager.IDBManager
	context   *NsiContext
}

type NewNSIParams struct {
	DbManager db_manager.IDBManager
}

func NewNSI(params NewNSIParams) (nsi.INSI, error) {
	nsi := NSI{
		dbManager: params.DbManager,
		context:   &NsiContext{},
	}

	ctx := context.Background()

	if err := nsi.Load(ctx); err != nil {
		return nil, wrapErr.NewWrapErr(fmt.Errorf("load data for NSI"), err)
	}

	return &nsi, nil
}

func (n *NSI) GetPlacesInsideBound(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error) {
	points := n.context.placesKDTreeByCoordinates.RangeSearch(kdrange.New(bottomLat, topLat, leftLng, rightLng))
	places := make([]models.Place, 0, len(points))
	for _, v := range points {
		p := v.(*KDTreePoints.Point)
		q := p.Data.(*models.Place)
		places = append(places, *q)
	}

	return places, nil
}

func (n *NSI) GetPlaceByID(ctx context.Context, id int) (models.Place, error) {
	place, has := n.context.placesByID[id]
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("get place with id=%d", id), errs.ErrorEntityNotFound)
		return models.Place{}, err
	}

	return *place, nil
}

func (n *NSI) GetUserSubscriptionsByFollowerID(ctx context.Context, followerID int) ([]models.UserSubscription, error) {
	userSubscriptions, has := n.context.userSubscriptionsByFollowerUserID[followerID]
	if !has {
		return []models.UserSubscription{}, nil
	}

	res := make([]models.UserSubscription, 0, len(userSubscriptions))
	for _, v := range userSubscriptions {
		res = append(res, *v)
	}

	return res, nil
}

func (n *NSI) GetPlaceEvaluationByUserIDByPlaceID(ctx context.Context, userID int, placeID int) (models.PlaceEvaluation, error) {
	v, has := n.context.placeEvaluationsByUserIDByPlaceID[userID]
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("get placeEvaluations with userID=%d", userID), errs.ErrorEntityNotFound)
		return models.PlaceEvaluation{}, err
	}
	v2, has := v[placeID]
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("get placeEvaluation of userID=%d with placeID=%d", userID, placeID), errs.ErrorEntityNotFound)
		return models.PlaceEvaluation{}, err
	}
	return *v2, nil
}

func (n *NSI) GetPlaceEvaluationMarksByPlaceEvaluationID(ctx context.Context, placeEvaluationID int) ([]models.PlaceEvaluationMark, error) {
	v, has := n.context.placeEvaluationMarksByPlaceEvaluationID[placeEvaluationID]
	if !has {
		err := wrapErr.NewWrapErr(fmt.Errorf("placeEvaluationMarksByPlaceEvaluationID placeEvaluationID=%d", placeEvaluationID), errs.ErrorEntityNotFound)
		return nil, err
	}

	res := make([]models.PlaceEvaluationMark, 0, len(v))
	for _, v2 := range v {
		res = append(res, *v2)
	}

	return res, nil
}
