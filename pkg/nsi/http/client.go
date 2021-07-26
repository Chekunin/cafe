package http

import (
	"cafe/pkg/models"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
)

type HttpNSI struct {
	httpClient *http.HttpClient
}

func NewHttpNSI(uri string) (*HttpNSI, error) {
	httpClient := http.NewHttpClient(http.HttpClientParams{
		BaseUrl:               uri,
		Headers:               map[string]string{},
		CodeToErrorMapping:    codeToError,
		RequestPayloadEncoder: http.GobEncoder,
		RequestPayloadDecoder: http.GobDecoder,
	})

	return &HttpNSI{httpClient: httpClient}, nil
}

var codeToError = map[int]error{}

func (n HttpNSI) GetPlacesInsideBound(ctx context.Context, leftLng float64, rightLng float64, topLat float64, bottomLat float64) ([]models.Place, error) {
	var resp []models.Place
	_, err := n.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/places-inside-bound"),
		UrlParams: map[string]string{
			"left_lng":   fmt.Sprintf("%f", leftLng),
			"right_lng":  fmt.Sprintf("%f", rightLng),
			"top_lat":    fmt.Sprintf("%f", topLat),
			"bottom_lat": fmt.Sprintf("%f", bottomLat),
		},
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return resp, err
	}
	return resp, nil
}

func (n HttpNSI) GetPlaceByID(ctx context.Context, id int) (models.Place, error) {
	var resp models.Place
	_, err := n.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/place-by-id/%d", id),
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return resp, err
	}
	return resp, nil
}
