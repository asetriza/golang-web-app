package rest

import (
	"github.com/asetriza/golang-web-app/internal/service"
	v1 "github.com/asetriza/golang-web-app/internal/transport/rest/v1"
	"github.com/asetriza/golang-web-app/pkg/auth"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	Service      *service.Service
	TokenManager auth.TokenManager
}

func NewHTTP(serv *service.Service, tm auth.TokenManager) *HTTP {
	return &HTTP{
		Service:      serv,
		TokenManager: tm,
	}
}

func (h *HTTP) Router() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		corsMiddleware,
	)

	h.initAPI(router)

	return router
}

func (h *HTTP) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.Service, h.TokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
