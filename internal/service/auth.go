package service

import (
	"crypto/sha1"
	"fmt"
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "asdfksjkaSDFASDF341223tF@#314f$23sdfas"
	signInKey = "asdf4rfh39qg4h3t93q$#Q#@rf2buwei"
	tokenTTL  = 12 * time.Hour
)

type AuthorizationService struct {
	Repository repository.Authorization
}

func NewAuthorizationService(ra repository.Authorization) *AuthorizationService {
	return &AuthorizationService{
		Repository: ra,
	}
}

func (as *AuthorizationService) CreateUser(user common.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return as.Repository.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (as *AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := as.Repository.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.StandardClaims{
			Subject:   strconv.Itoa(user.ID),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		})

	return token.SignedString([]byte(signInKey))
}
