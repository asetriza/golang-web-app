package v1

import (
	"golang-web-app/internal/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTodoRoute(api *gin.RouterGroup) {
	todo := api.Group("/todo", h.Middleware.Identity)
	{
		todo.POST("/", h.createTodo)
		todo.GET("/", h.getTodos)
		todo.GET("/:id", h.getTodo)
		todo.PUT("/:id", h.updateTodo)
		todo.DELETE("/:id", h.deleteTodo)
	}
}

func (h *Handler) createTodo(c *gin.Context) {
	var input common.Todo

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Todo.Create(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getTodos(c *gin.Context) {

}

func (h *Handler) getTodo(c *gin.Context) {

}

func (h *Handler) updateTodo(c *gin.Context) {

}

func (h *Handler) deleteTodo(c *gin.Context) {

}
