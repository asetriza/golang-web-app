package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (r *REST) parseAuthHeader(c *gin.Context) (int, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, errors.New("token is empty")
	}

	return r.TokenManager.ParseToken(headerParts[1])
}

func (r *REST) SetUserIDToCtx(c *gin.Context) {
	id, err := r.parseAuthHeader(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, id)
}

func getUserIDFromCtx(c *gin.Context) (int, error) {
	idFromCtx, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("userCtx not found")
	}

	id, ok := idFromCtx.(int)
	if !ok {
		return 0, errors.New("userCtx is invalid type")
	}

	return id, nil
}
