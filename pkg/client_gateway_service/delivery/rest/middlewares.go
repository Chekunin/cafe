package rest

import (
	clientSsoErrs "cafe/pkg/client_sso/errors"
	"cafe/pkg/common"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func (r *rest) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("CLIENT_NOSIGN") != "" {
			c.Next()
			return
		}

		token, err := getTokenFromHeader(c.Request)
		if err != nil {
			err = wrapErr.NewWrapErr(fmt.Errorf("getTokenFromHeader"), err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), common.ContextKeyToken, token))

		if _, err := r.clientSso.CheckPermission(c.Request.Context(), c.Request.Method, c.Request.URL.Path, token); err != nil {
			switch {
			case errors.Is(err, clientSsoErrs.ErrIncorrectToken):
				c.AbortWithError(http.StatusForbidden, common.ErrPermissionDenied)
			default:
				c.AbortWithError(http.StatusInternalServerError, common.ErrInternalServerError)
			}
			return
		}

		c.Next()
	}
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", wrapErr.NewWrapErr(fmt.Errorf("Authorization header is empty"), nil)
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", wrapErr.NewWrapErr(fmt.Errorf("Authorization header is incorrect"), nil)
	}

	return authHeaderParts[1], nil
}
