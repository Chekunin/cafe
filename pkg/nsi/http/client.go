package http

import (
	"cafe/pkg/common"
	"cafe/pkg/models"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
)

type HttpNSI struct {
	httpClient *http.HttpClient
}

func NewHttpNSI(uri string) (*HttpNSI, error) {
	httpClient := http.NewHttpClient(http.HttpClientParams{
		BaseUrl: uri,
		Headers: map[string]string{},
		ErrorHandler: func(reader io.Reader) error {
			var err common.Err
			if err2 := http.GobDecoder(reader, &err); err2 != nil {
				err2 = wrapErr.NewWrapErr(fmt.Errorf("http GobDecoder"), err2)
				return err2
			}
			if err, has := codeToError[err.Code]; has {
				return err
			}
			return common.ErrInternalServerError
		},
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
