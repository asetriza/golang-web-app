package auth

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing
type TokenManager interface {
	NewToken(userId string) (string, error)
	ParseToken(accessToken string) (string, error)
	NewRefreshToken() string
	CreateRefreshTokenTTL() int64
}

type Token struct {
	signingKey      string
	tokenTTL        int
	refreshToken    string
	refreshTokenTTL int64
}

func NewManager(signingKey string, tokenTTL, refreshTokenTTL string) *Token {
	if signingKey == "" {
		log.Fatalf("signing key parsing error")
	}

	tokenttl, err := strconv.Atoi(tokenTTL)
	if err != nil {
		log.Fatalf("tokenTTL parsing error")
	}

	refreshTokenttl, err := strconv.ParseInt(refreshTokenTTL, 10, 64)
	if err != nil {
		log.Fatalf("refreshTokenTTL parsing error")
	}

	return &Token{
		signingKey:      signingKey,
		tokenTTL:        tokenttl,
		refreshTokenTTL: refreshTokenttl,
	}
}

func (t *Token) NewToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(t.tokenTTL) * time.Minute).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *Token) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (t *Token) NewRefreshToken() string {
	t.refreshToken = uuid.NewString()
	return t.refreshToken
}

func (t *Token) CreateRefreshTokenTTL() int64 {
	return t.refreshTokenTTL*60 + time.Now().Unix()
}
