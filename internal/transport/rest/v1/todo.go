package v1

import (
	"github.com/asetriza/golang-web-app/internal/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initTodoRoute(api *gin.RouterGroup) {
	todo := api.Group("/todo", h.Middleware.IdentifyUser)
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

	userID, err := getUserID(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.UserID = userID

	id, err := h.Service.Todo.Create(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type getTodoInput struct {
	todoID int `binding:"required"`
}

func (h *Handler) getTodo(c *gin.Context) {
	var input getTodoInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := h.Service.Todo.Get(c.Request.Context(), input.todoID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todo": todo})
}

func (h *Handler) getTodos(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	todos, err := h.Service.Todo.GetAll(c.Request.Context(), userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todos": todos})
}

func (h *Handler) updateTodo(c *gin.Context) {
	var input common.Todo
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Todo.Update(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type deleteTodoInput struct {
	todoID int `binding:"required"`
}

func (h *Handler) deleteTodo(c *gin.Context) {
	var input deleteTodoInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Todo.Delete(c.Request.Context(), input.todoID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
