package http

import (
	"cafe/pkg/common"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	wrapErr "github.com/Chekunin/wraperr"
	"moul.io/http2curl"
)

type DataEncoder func(payload interface{}) (io.Reader, error)
type DataDecoder func(reader io.Reader, res interface{}) error

type HttpClient struct {
	baseUrl               string
	codeToErrorMapping    map[int]error
	headers               map[string]string // заголовки, которые всегда будут передаваться
	httpClient2           *http.Client
	requestPayloadEncoder DataEncoder
	requestPayloadDecoder DataDecoder
}

type HttpClientParams struct {
	BaseUrl               string
	CodeToErrorMapping    map[int]error
	Headers               map[string]string
	Timeout               time.Duration // максимальное время, которое будет длиться запрос
	RequestPayloadEncoder DataEncoder
	RequestPayloadDecoder DataDecoder
	MaxIdleConnsPerHost   int
}

func NewHttpClient(params HttpClientParams) *HttpClient {
	transport := getDefaultHttpTransport()
	transport.MaxIdleConnsPerHost = params.MaxIdleConnsPerHost
	if params.MaxIdleConnsPerHost == 0 {
		transport.MaxIdleConnsPerHost = 100
	}

	client := HttpClient{
		baseUrl:            params.BaseUrl,
		codeToErrorMapping: params.CodeToErrorMapping,
		headers:            params.Headers,
		httpClient2: &http.Client{
			Timeout:   params.Timeout,
			Transport: transport,
		},
		requestPayloadEncoder: params.RequestPayloadEncoder,
		requestPayloadDecoder: params.RequestPayloadDecoder,
	}
	if client.requestPayloadEncoder == nil {
		client.requestPayloadEncoder = JsonEncoder
	}
	if client.requestPayloadDecoder == nil {
		client.requestPayloadDecoder = JsonDecoder
	}
	return &client
}

func getDefaultHttpTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

func (c *HttpClient) SetHeaders(headers map[string]string) {
	c.headers = headers
}

func (c *HttpClient) PostRequest(ctx context.Context,
	url string,
	headers map[string]string,
	payload interface{},
	result interface{}) (*http.Response, error) {
	return c.DoRequest(ctx, "POST", url, headers, payload, result)
}

func (c *HttpClient) GetRequest(ctx context.Context,
	url string,
	headers map[string]string,
	result interface{}) (*http.Response, error) {
	return c.DoRequest(ctx, "GET", url, headers, nil, result)
}

type RequestOptions struct {
	Ctx                   context.Context
	Method                string
	Url                   string
	Headers               map[string]string
	Payload               interface{}
	Result                interface{}
	RequestPayloadEncoder DataEncoder
	RequestPayloadDecoder DataDecoder
	UrlParams             map[string]string
}

func (c HttpClient) setDefaultOptions(opt *RequestOptions) {
	if opt == nil {
		return
	}

	if opt.Ctx == nil {
		opt.Ctx = context.Background()
	}
	if opt.Method == "" {
		opt.Method = "GET"
	}
	if opt.RequestPayloadEncoder == nil {
		opt.RequestPayloadEncoder = c.requestPayloadEncoder
	}
	if opt.RequestPayloadDecoder == nil {
		opt.RequestPayloadDecoder = c.requestPayloadDecoder
	}
}

func (c *HttpClient) DoRequestWithOptions(options RequestOptions) (*http.Response, error) {
	c.setDefaultOptions(&options)
	payloadReader, err := options.RequestPayloadEncoder(options.Payload)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("requestPayloadEncoder"), err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(
		options.Ctx,
		options.Method,
		fmt.Sprintf("%s%s", c.baseUrl, options.Url),
		payloadReader,
	)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("new request with context"), err)
		return nil, err
	}

	q := req.URL.Query()
	for key, val := range options.UrlParams {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	if options.Ctx != nil {
		requestID, _ := common.FromContextRequestId(options.Ctx)
		req.Header.Set(string(common.HeaderKeyRequestID), requestID)
	}
	for i, v := range c.headers {
		req.Header.Set(i, v)
	}
	for i, v := range options.Headers {
		req.Header.Set(i, v)
	}

	curl, _ := http2curl.GetCurlCommand(req)
	t := time.Now()
	resp, err := c.httpClient2.Do(req)
	requestTakes := fmt.Sprintf("request takes %f microseconds", float64(time.Now().UnixNano()-t.UnixNano())/float64(time.Microsecond))
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("do http request curl --- %s --- %s", curl, requestTakes), err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		err := wrapErr.NewWrapErr(fmt.Errorf("http status code=%d curl --- %s --- %s", resp.StatusCode, curl, requestTakes), c.mapPaymentErrorToLocal(resp.Body))
		return nil, err
	}
	if options.Result != nil {
		if err := options.RequestPayloadDecoder(resp.Body, options.Result); err != nil {
			err = wrapErr.NewWrapErr(fmt.Errorf("decode resp.Body curl --- %s --- %s", curl, requestTakes), err)
			return nil, err
		}
	}

	return resp, nil
}

func (c *HttpClient) DoRequest(ctx context.Context,
	method string,
	url string,
	headers map[string]string,
	payload interface{},
	result interface{}) (*http.Response, error) {
	return c.DoRequestWithOptions(RequestOptions{
		Ctx:     ctx,
		Method:  method,
		Url:     url,
		Headers: headers,
		Payload: payload,
		Result:  result,
	})
}

// todo: подумать какие ошибки здесь отдавать, свои или чужие
func (c *HttpClient) mapPaymentErrorToLocal(respBody io.ReadCloser) error {
	type Err2 struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    interface{} `json:"meta"`
	}

	var e Err2
	if err := json.NewDecoder(respBody).Decode(&e); err != nil {
		return common.ErrInternalServerError
	}
	if err, has := c.codeToErrorMapping[e.Code]; has {
		return err
	} else {
		return common.ErrInternalServerError
	}
}
