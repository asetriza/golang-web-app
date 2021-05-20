package http

import (
	"golang-web-app/internal/service"
	v1 "golang-web-app/internal/transport/http/v1"

	"github.com/gin-gonic/gin"
)

type HTTP struct {
	service *service.Service
}

func NewHTTP(s *service.Service) *HTTP {
	return &HTTP{
		service: s,
	}
}

func (h *HTTP) Init() *gin.Engine {
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
	handlerV1 := v1.NewHandler(h.service)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
