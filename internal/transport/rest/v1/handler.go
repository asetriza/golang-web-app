package v1

import (
	"github.com/asetriza/golang-web-app/internal/service"
	"github.com/asetriza/golang-web-app/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Middleware
}

func NewHandler(serv *service.Service, tm auth.TokenManager) *Handler {
	return &Handler{
		Service:    serv,
		Middleware: NewMiddleware(tm),
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoute(v1)
		h.initTodoRoute(v1)
	}
}
