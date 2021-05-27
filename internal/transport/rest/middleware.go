package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/asetriza/golang-web-app/pkg/uuid"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	infoCtx             = "info"
	requestIDKey        = "requestIDKey"
	userIDKey           = "userIDKey"
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

func logRequest(c *gin.Context) string {

	var infoString string
	formatter := func(param gin.LogFormatterParams) string {
		body, _ := ioutil.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		headers := c.Request.Header

		info := struct {
			ClientIP     string        `json:"clientIp"`
			TimeStamp    string        `json:"timeStamp"`
			Headers      interface{}   `json:"headers"`
			Method       string        `json:"method"`
			Path         string        `json:"path"`
			RequestProto string        `json:"requestProto"`
			StatusCode   int           `json:"statusCode"`
			Body         string        `json:"body"`
			Latency      time.Duration `json:"latency"`
			UserAgent    string        `json:"userAgent"`
			ErrorMessage string        `json:"errorMessage"`
		}{
			ClientIP:     param.ClientIP,
			TimeStamp:    param.TimeStamp.Format(time.RFC1123),
			Headers:      headers,
			Method:       param.Method,
			Path:         param.Path,
			RequestProto: param.Request.Proto,
			StatusCode:   param.StatusCode,
			Body:         string(body),
			Latency:      param.Latency,
			UserAgent:    param.Request.UserAgent(),
			ErrorMessage: param.ErrorMessage,
		}

		infoJson, err := json.Marshal(info)
		if err != nil {
			log.Println(err)
		}
		infoString = string(infoJson)
		return string(infoJson)
	}

	log := gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: formatter,
		Output:    ioutil.Discard,
	})

	log(c)

	return infoString
}

func SetRqIDToCtx(c *gin.Context) {
	c.Set(requestIDKey, uuid.Create())
}

func loggerMiddleware(c *gin.Context) {
	ctx := WithInfo(c.Request.Context(), logRequest(c))
	ctx1 := WithRqID(ctx, getRqIDFromCtx(c))
	userID, err := getUserIDFromCtx(c)
	if err != nil {
		userID = -1
	}
	ctx2 := WithUserID(ctx1, userID)
	logMw(ctx2)
}

func logMw(ctx context.Context) {
	logger, err := Logger(ctx)
	if err != nil {
		log.Println(err)
	}
	logger.Info("FROM MIDDLEWARE")
	defer logger.Sync()
}

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
	userID, err := r.parseAuthHeader(c)
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userID)
}

func getUserIDFromCtx(c *gin.Context) (int, error) {
	userIDFromCtx, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("userCtx not found")
	}

	userID, ok := userIDFromCtx.(int)
	if !ok {
		return 0, errors.New("userCtx is invalid type")
	}

	return userID, nil
}

func getRqIDFromCtx(c *gin.Context) string {
	ctxRqIDFromCtx, ok := c.Get(requestIDKey)
	if !ok {
		log.Println("rqCtx not found")
	}

	rqID, ok := ctxRqIDFromCtx.(string)
	if !ok {
		log.Println("rqCtx is invalid type")
	}

	return rqID
}

func getInfoFromCtx(c *gin.Context) string {
	infoFromCtx, ok := c.Get(infoCtx)
	if !ok {
		log.Println("infoCtx not found")
	}
	log.Println(infoFromCtx)

	info, ok := infoFromCtx.(string)
	if !ok {
		log.Println("infoCtx is invalid type")
	}

	return info
}
