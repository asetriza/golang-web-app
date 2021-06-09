package rest

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/service"

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
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		default:
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
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
		switch err {
		case sql.ErrNoRows:
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		default:
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
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

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := getUserIDFromCtx(c)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.ID = todoID
	input.UserID = userID

	err = r.Service.Todo.Update(c.Request.Context(), input)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		case service.AccessDenied:
			newErrorResponce(c, http.StatusForbidden, err.Error())
			return
		default:
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": input.ID})
}

func (r *REST) deleteTodo(c *gin.Context) {
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

	err = r.Service.Todo.Delete(c.Request.Context(), userID, todoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponce(c, http.StatusNotFound, err.Error())
			return
		case service.AccessDenied:
			newErrorResponce(c, http.StatusForbidden, err.Error())
			return
		default:
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": todoID})
}
