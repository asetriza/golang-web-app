package rest

import (
	"net/http"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/service"

	"github.com/gin-gonic/gin"
)

func (r *REST) signUp(c *gin.Context) {
	var input common.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	credentials, err := r.Service.Authorization.CreateUser(c.Request.Context(), input, c.ClientIP())
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"credentials": credentials})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *REST) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	credentials, err := r.Service.Authorization.CreateCredentials(c.Request.Context(), input.Username, input.Password, c.ClientIP())
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"credentials": credentials})
}

type refreshInput struct {
	Token        string `json:"token" binding:"required"`
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func (r *REST) refresh(c *gin.Context) {
	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	credentials, err := r.Service.Authorization.RefreshCredentials(c.Request.Context(), input.Token, input.RefreshToken, c.ClientIP())
	if err != nil {
		switch err {
		case service.RefreshTokenExpired:
			newErrorResponce(c, http.StatusUnauthorized, err.Error())
			return
		default:
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"credentials": credentials})
}
