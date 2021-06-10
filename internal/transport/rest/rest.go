package rest

import (
	"errors"
	"strconv"

	"github.com/asetriza/golang-web-app/internal/service"
	"github.com/asetriza/golang-web-app/pkg/auth"

	"github.com/gin-gonic/gin"
)

type REST struct {
	Service      *service.Service
	TokenManager auth.TokenManager
}

func NewREST(s *service.Service, tm auth.TokenManager) *REST {
	return &REST{
		Service:      s,
		TokenManager: tm,
	}
}

func (r *REST) Router() *gin.Engine {
	router := gin.New()
	router.Use(
		SetRqIDToCtx,
		loggerMiddleware,
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

func parseIdFromPath(c *gin.Context) (int, error) {
	idParam := c.Param("id")
	if idParam == "" {
		return 0, errors.New("empty id param")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errors.New("id param must be int")
	}

	return id, nil
}
