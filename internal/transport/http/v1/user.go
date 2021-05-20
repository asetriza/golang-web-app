package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserRoute(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.POST("/", h.createUser)
		user.GET("/", h.getUsers)
		user.GET("/:id", h.getUser)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}
}

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) getUsers(c *gin.Context) {

}

func (h *Handler) getUser(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
