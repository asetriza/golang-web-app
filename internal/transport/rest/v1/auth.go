package v1

import (
	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthRoute(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refresh)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var input common.User

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Authorization.CreateUser(c.Request.Context(), input, c.ClientIP())
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	credentials, err := h.Service.Authorization.CreateCredentials(c.Request.Context(), input.Username, input.Password, c.ClientIP())
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"credentials": credentials})
}

type refreshInput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) refresh(c *gin.Context) {
	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		log.Println(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	credentials, err := h.Service.Authorization.RefreshCredentials(c.Request.Context(), input.Token, input.RefreshToken, c.ClientIP())
	if err != nil {
		if err == service.RefreshTokenExpired {
			newErrorResponce(c, http.StatusUnauthorized, err.Error())
			return
		}
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"credentials": credentials})
}
