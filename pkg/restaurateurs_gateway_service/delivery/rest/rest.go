package rest

import (
	"cafe/pkg/restaurateurs_gateway_service/usecase"
	"github.com/gin-gonic/gin"
)

const modelTag string = "restaurateurs"

type rest struct {
	usecase *usecase.Usecase
}

func NewRest(router *gin.RouterGroup, usecase *usecase.Usecase) *rest {
	rest := &rest{
		usecase: usecase,
	}
	rest.routes(router)
	return rest
}

func (r *rest) routes(router *gin.RouterGroup) {
	router.POST("/auth/login", nil)
	router.POST("/auth/refresh-token", nil)
	router.POST("/auth/signup", nil)
	router.POST("/auth/approve-email", nil)
}
