package http

import (
	"cafe/pkg/transport/http"
	"context"
)

type HttpClientSso struct {
	httpClient *http.HttpClient
}

func NewHttpNSI(uri string) (*HttpClientSso, error) {
	httpClient := http.NewHttpClient(http.HttpClientParams{
		BaseUrl:               uri,
		Headers:               map[string]string{},
		CodeToErrorMapping:    codeToError,
		RequestPayloadEncoder: http.GobEncoder,
		RequestPayloadDecoder: http.GobDecoder,
	})

	return &HttpClientSso{httpClient: httpClient}, nil
}

var codeToError = map[int]error{}

func (h HttpClientSso) Login(ctx context.Context, userName string, password string) (string, error) {
	panic("implement me")
}

func (h HttpClientSso) CheckPermission(ctx context.Context, method, path, token string) (bool, error) {
	panic("implement me")
}

func (h HttpClientSso) GetUserID(ctx context.Context, token string) (int, error) {
	panic("implement me")
}
