package common

import (
	"bytes"
	"cafe/pkg/common/catcherr"
	log "cafe/pkg/common/logman"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"time"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(HeaderKeyRequestID)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ContextKeyRequestId, requestId))
		c.Writer.Header().Set(HeaderKeyRequestID, requestId)
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				var err2 error
				switch v := err.(type) {
				case nil:
					err2 = fmt.Errorf("nil")
				case error:
					err2 = v
				case string:
					err2 = fmt.Errorf("%s", v)
				default:
					err2 = fmt.Errorf("%v", v)
				}
				err2 = wrapErr.NewWrapErr(err2, fmt.Errorf(string(debug.Stack())))
				err2 = wrapErr.NewWrapErr(fmt.Errorf("panic"), err2)
				catcherr.AsCritical().Catch(err2)
				c.AbortWithError(http.StatusInternalServerError, err2)
			}
		}()
		c.Next()
	}
}

func ErrorResponder(sendDataFunc func(c *gin.Context, code int, obj interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0]
			var ret interface{}
			var respErr Err
			if errors.As(err.Err, &respErr) {
				ret = respErr
			} else {
				if c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
					ret = ErrBadRequest
				} else {
					ret = ErrInternalServerError
				}
			}
			sendDataFunc(c, 0, ret)
		}
	}
}

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		// for response body log
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()
		for _, err := range c.Errors {
			catcherr.Catch(wrapErr.NewWrapErr(fmt.Errorf("Request error"), err.Err), catcherr.Fields{
				"duration":      float64(time.Now().Sub(startTime)) / float64(time.Second),
				"uri":           c.Request.RequestURI,
				"method":        c.Request.Method,
				"headers":       c.Request.Header,
				"body":          ellipsis(string(body), ellipsisLength),
				"error":         err,
				"status":        c.Writer.Status(),
				"request_id":    c.Request.Header.Get(string(HeaderKeyRequestID)),
				"response_body": ellipsis(blw.body.String(), ellipsisLength),
			})
		}
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		// for response body log
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		log.Info("Request", log.Fields{
			"duration":      time.Now().Sub(startTime),
			"uri":           c.Request.RequestURI,
			"method":        c.Request.Method,
			"headers":       c.Request.Header,
			"body":          ellipsis(string(body), ellipsisLength),
			"status":        c.Writer.Status(),
			"request_id":    c.Request.Header.Get(string(HeaderKeyRequestID)),
			"response_body": ellipsis(blw.body.String(), ellipsisLength),
		})
	}
}

// for response body log
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

const ellipsisLength = 5 * 1024

func ellipsis(s string, l int) string {
	if len(s) > l {
		return s[:l] + "..."
	}
	return s
}
