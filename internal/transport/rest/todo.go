package rest

import (
	"log"
	"net/http"

	"github.com/asetriza/golang-web-app/internal/common"

	"github.com/gin-gonic/gin"
)

func (r *REST) createTodo(c *gin.Context) {
	var input common.Todo
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.UserID = userID

	id, err := r.Service.Todo.Create(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type getTodoInput struct {
	todoID int `binding:"required"`
}

func (r *REST) getTodo(c *gin.Context) {
	var input getTodoInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := r.Service.Todo.Get(c.Request.Context(), input.todoID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todo": todo})
}

func (r *REST) getTodos(c *gin.Context) {
	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	todos, err := r.Service.Todo.GetAll(c.Request.Context(), userID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todos": todos})
}

func (r *REST) updateTodo(c *gin.Context) {
	var input common.Todo
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.Service.Todo.Update(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type deleteTodoInput struct {
	todoID int `binding:"required"`
}

func (r *REST) deleteTodo(c *gin.Context) {
	var input deleteTodoInput
	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.Service.Todo.Delete(c.Request.Context(), input.todoID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
