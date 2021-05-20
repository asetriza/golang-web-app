package v1

import (
	"golang-web-app/internal/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthRoute(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var input common.User

	if err := c.BindJSON(&input); err != nil {
		log.Print(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type singInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		log.Print(err)
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.Service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
