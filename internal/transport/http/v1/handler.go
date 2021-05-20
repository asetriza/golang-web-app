package v1

import (
	"golang-web-app/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		Service: serv,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoute(v1)
		h.initUserRoute(v1)
	}
}
