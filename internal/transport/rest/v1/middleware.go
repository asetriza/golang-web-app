package v1

import (
	"errors"
	"github.com/asetriza/golang-web-app/pkg/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type Middleware struct {
	tokenManager auth.TokenManager
}

func NewMiddleware(tm auth.TokenManager) Middleware {
	return Middleware{
		tokenManager: tm,
	}
}

func (m *Middleware) parseAuthHeader(c *gin.Context) (int, error) {
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

	return m.tokenManager.ParseToken(headerParts[1])
}

func (m *Middleware) IdentifyUser(c *gin.Context) {
	id, err := m.parseAuthHeader(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, id)
}

func getUserID(c *gin.Context) (int, error) {
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
