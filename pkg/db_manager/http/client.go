package http

import (
	errs "cafe/pkg/db_manager/errors"
	"cafe/pkg/models"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

type HttpDbManager struct {
	httpClient *http.HttpClient
}

func NewHttpDbManager(uri string) (*HttpDbManager, error) {
	httpClient := http.NewHttpClient(http.HttpClientParams{
		BaseUrl:               uri,
		Headers:               map[string]string{},
		CodeToErrorMapping:    codeToError,
		RequestPayloadEncoder: http.GobEncoder,
		RequestPayloadDecoder: http.GobDecoder,
	})

	return &HttpDbManager{httpClient: httpClient}, nil
}

var codeToError = map[int]error{
	1: errs.ErrorEntityNotFound,
}

func (h HttpDbManager) GetAllPlaces(ctx context.Context) ([]models.Place, error) {
	var resp []models.Place
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/places",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceMedias(ctx context.Context) ([]models.PlaceMedia, error) {
	var resp []models.PlaceMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var resp []models.Category
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceCategories(ctx context.Context) ([]models.PlaceCategory, error) {
	var resp []models.PlaceCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllKitchenCategories(ctx context.Context) ([]models.KitchenCategory, error) {
	var resp []models.KitchenCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/kitchen-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceKitchenCategories(ctx context.Context) ([]models.PlaceKitchenCategory, error) {
	var resp []models.PlaceKitchenCategory
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-kitchen-categories",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceSchedules(ctx context.Context) ([]models.PlaceSchedule, error) {
	var resp []models.PlaceSchedule
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-schedules",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllAdverts(ctx context.Context) ([]models.Advert, error) {
	var resp []models.Advert
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/adverts",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllAdvertMedias(ctx context.Context) ([]models.AdvertMedia, error) {
	var resp []models.AdvertMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/advert-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllEvaluationCriterions(ctx context.Context) ([]models.EvaluationCriterion, error) {
	var resp []models.EvaluationCriterion
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/evaluation-criterions",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceEvaluations(ctx context.Context) ([]models.PlaceEvaluation, error) {
	var resp []models.PlaceEvaluation
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-evaluations",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllPlaceEvaluationMarks(ctx context.Context) ([]models.PlaceEvaluationMark, error) {
	var resp []models.PlaceEvaluationMark
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/place-evaluation-marks",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllReviews(ctx context.Context) ([]models.Review, error) {
	var resp []models.Review
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/reviews",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllReviewMedias(ctx context.Context) ([]models.ReviewMedia, error) {
	var resp []models.ReviewMedia
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/review-medias",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var resp []models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/users",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}

func (h HttpDbManager) GetUserByUserName(ctx context.Context, userName string) (models.User, error) {
	var resp models.User
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/user-by-name/%s", userName),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.User{}, err
	}
	return resp, nil
}

func (h HttpDbManager) GetAllUserSubscriptions(ctx context.Context) ([]models.UserSubscription, error) {
	var resp []models.UserSubscription
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    "/user-subscriptions",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, err
	}
	return resp, nil
}