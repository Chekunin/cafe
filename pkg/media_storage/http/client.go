package http

import (
	"bytes"
	"cafe/pkg/common"
	"cafe/pkg/media_storage"
	"cafe/pkg/review_media_storage_service/api/delivery/rest/schema"
	"cafe/pkg/transport/http"
	"context"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"io"
	"os"
)

type HttpMediaStorage struct {
	httpClient *http.HttpClient
}

func NewHttpMediaStorage(uri string) (*HttpMediaStorage, error) {
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

	return &HttpMediaStorage{httpClient: httpClient}, nil
}

var codeToError = map[int]error{}

func (h HttpMediaStorage) Get(ctx context.Context, path string) (*os.File, error) {
	panic("implement me")
}

func (h HttpMediaStorage) GetStream(ctx context.Context, path string) (io.ReadCloser, string, error) {
	var respBytes []byte
	response, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:    ctx,
		Method: "GET",
		Url:    fmt.Sprintf("/media"),
		UrlParams: map[string]string{
			"path": path,
		},
		Result: &respBytes,
		RequestPayloadDecoder: func(reader io.Reader, res interface{}) error {
			buf := new(bytes.Buffer)
			if _, err := buf.ReadFrom(reader); err != nil {
				err = wrapErr.NewWrapErr(fmt.Errorf("buf ReadFrom"), err)
				return err
			}
			respBytes = buf.Bytes()
			return nil
		},
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return nil, "", err
	}
	readCloser := io.NopCloser(bytes.NewReader(respBytes))
	return readCloser, response.Header.Get("Content-Type"), nil
}

func (h HttpMediaStorage) Put(ctx context.Context, path string, reader io.Reader) (media_storage.Object, error) {
	var resp media_storage.Object

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	payload := schema.ReqPutMedia{
		Path: path,
		Data: buf.Bytes(),
	}
	_, err := h.httpClient.DoRequestWithOptions(http.RequestOptions{
		Ctx:     ctx,
		Method:  "POST",
		Url:     fmt.Sprintf("/media"),
		Payload: payload,
		Result:  &resp,
	})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request"), err)
		return media_storage.Object{}, err
	}

	return resp, nil
}

func (h HttpMediaStorage) Delete(ctx context.Context, path string) error {
	panic("implement me")
}

func (h HttpMediaStorage) List(ctx context.Context, path string) ([]media_storage.Object, error) {
	panic("implement me")
}
