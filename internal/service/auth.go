package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
	"golang-web-app/pkg/auth"
	"log"
	"strconv"
	"time"
)

var (
	TokenExpired = errors.New("token expired")
)

type AuthorizationService struct {
	Repository   repository.Authorization
	TokenManager auth.TokenManager
	PasswordSalt string
}

func NewAuthorizationService(ra repository.Authorization, tm auth.TokenManager, passSalt string) *AuthorizationService {
	return &AuthorizationService{
		Repository:   ra,
		TokenManager: tm,
		PasswordSalt: passSalt,
	}
}

func (as *AuthorizationService) CreateUser(user common.User) (int, error) {
	user.Password = as.generateHash(user.Password)

	return as.Repository.CreateUser(user)
}

func (as *AuthorizationService) generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(as.PasswordSalt)))
}

func (as *AuthorizationService) GenerateCredentials(username, password string) (Credentials, error) {
	user, err := as.Repository.GetUser(username, as.generateHash(password))
	if err != nil {
		return Credentials{}, err
	}

	token, err := as.TokenManager.NewToken(strconv.Itoa(user.ID))
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	refreshToken := as.TokenManager.NewRefreshToken()
	refreshTokenTTL := as.TokenManager.CreateRefreshTokenTTL()
	if _, err := as.Repository.UpdateUserSession(user.ID, refreshToken, refreshTokenTTL); err != nil {
		if _, err = as.Repository.CreateUserSession(user.ID, refreshToken, refreshTokenTTL); err != nil {
			return Credentials{}, err
		}
	}

	return Credentials{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

type Credentials struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (as *AuthorizationService) RefreshCredentials(token, refreshToken string) (Credentials, error) {
	userID, err := as.TokenManager.ParseToken(token)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	userSession, err := as.Repository.GetUserSession(userIDint, refreshToken)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	if userSession.RefreshTokenTTL > time.Now().UnixNano() {
		log.Println(TokenExpired)
		return Credentials{}, TokenExpired
	}

	token, err = as.TokenManager.NewToken(userID)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	refreshToken = as.TokenManager.NewRefreshToken()
	refreshTokenTTL := as.TokenManager.CreateRefreshTokenTTL()
	if _, err = as.Repository.UpdateUserSession(userIDint, refreshToken, refreshTokenTTL); err != nil {
		return Credentials{}, err
	}

	return Credentials{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
