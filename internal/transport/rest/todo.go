package rest

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/asetriza/golang-web-app/internal/common"

	"github.com/gin-gonic/gin"
)

func (r *REST) createTodo(c *gin.Context) {
	var input common.Todo
	if err := c.BindJSON(&input); err != nil {
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
	todoIDParam := c.Param("id")
	if todoIDParam == "" {
		newErrorResponce(c, http.StatusBadRequest, "empty id param")
		return
	}

	todoID, err := strconv.Atoi(todoIDParam)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "id param must be int")
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	todo, err := r.Service.Todo.Get(c.Request.Context(), userID, todoID)
	switch err {
	case sql.ErrNoRows:
		newErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todo": todo})
}

func (r *REST) getTodos(c *gin.Context) {
	var input common.Pagination
	if err := c.BindQuery(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	todos, err := r.Service.Todo.GetAll(c.Request.Context(), userID, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"todos": todos})
}

func (r *REST) updateTodo(c *gin.Context) {
	var input common.Todo

	todoIDParam := c.Param("id")
	if todoIDParam == "" {
		newErrorResponce(c, http.StatusBadRequest, "empty id param")
		return
	}

	todoID, err := strconv.Atoi(todoIDParam)
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "id param must be int")
		return
	}

	input.ID = todoID

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.UserID = userID

	err = r.Service.Todo.Update(c.Request.Context(), input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": input.ID})
}

type deleteTodoInput struct {
	TodoID int `json:"todoId" binding:"required"`
}

func (r *REST) deleteTodo(c *gin.Context) {
	var input deleteTodoInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = r.Service.Todo.Delete(c.Request.Context(), userID, input.TodoID)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": input.TodoID})
}
