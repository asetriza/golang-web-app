package rest

import (
	"github.com/asetriza/golang-web-app/internal/service"
	"github.com/asetriza/golang-web-app/pkg/auth"

	"github.com/gin-gonic/gin"
)

type REST struct {
	Service      *service.Service
	TokenManager auth.TokenManager
}

func NewREST(serv *service.Service, tm auth.TokenManager) *REST {
	return &REST{
		Service:      serv,
		TokenManager: tm,
	}
}

func (r *REST) Router() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		corsMiddleware,
	)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", r.signUp)
			auth.POST("/sign-in", r.signIn)
			auth.POST("/refresh", r.refresh)
		}

		todo := api.Group("/todo", r.SetUserIDToCtx)
		{
			todo.POST("/", r.createTodo)
			todo.GET("/", r.getTodos)
			todo.GET("/:id", r.getTodo)
			todo.PUT("/:id", r.updateTodo)
			todo.DELETE("/:id", r.deleteTodo)
		}
	}

	return router
}