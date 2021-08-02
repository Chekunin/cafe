package http

import (
	"cafe/pkg/client_sso/models"
	"cafe/pkg/common"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
)

type HttpClientSso struct {
	httpClient *http.HttpClient
}

func NewHttpClientSso(uri string) (*HttpClientSso, error) {
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

	return &HttpClientSso{httpClient: httpClient}, nil
}

var codeToError = map[int]error{}

func (h HttpClientSso) Login(ctx context.Context, userName string, password string) (models.Tokens, error) {
	var resp models.Tokens
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Payload: struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}{
			UserName: userName,
			Password: password,
		},
		Url:    "/login",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.Tokens{}, err
	}
	return resp, nil
}

func (h HttpClientSso) Logout(ctx context.Context, token string) error {
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Payload: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
		Url: "/logout",
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return err
	}
	return nil
}

func (h HttpClientSso) RefreshToken(ctx context.Context, refreshToken string) (models.Tokens, error) {
	var resp models.Tokens
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Payload: struct {
			RefreshToken string `json:"refresh_token"`
		}{
			RefreshToken: refreshToken,
		},
		Url:    "/refresh-token",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return models.Tokens{}, err
	}
	return resp, nil
}

func (h HttpClientSso) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	var resp bool
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "POST",
		Payload: struct {
			Method string `json:"method"`
			Path   string `json:"path"`
			Token  string `json:"token"`
		}{
			Method: method,
			Path:   path,
			Token:  token,
		},
		Url:    "/check-permission",
		Result: &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return false, err
	}
	return resp, nil
}

func (h HttpClientSso) GetUserID(ctx context.Context, token string) (int, error) {
	panic("implement me")
}
