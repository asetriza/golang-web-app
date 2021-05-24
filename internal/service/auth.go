package service

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
	"golang-web-app/pkg/auth"
	"log"
	"time"
)

var (
	RefreshTokenExpired = errors.New("refreshToken expired")
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

func (as *AuthorizationService) CreateUser(ctx context.Context, user common.User, userIP string) (Credentials, error) {
	user.Password = as.generateHash(user.Password)

	userID, err := as.Repository.CreateUser(ctx, user)
	if err != nil {
		return Credentials{}, err
	}

	return as.createUserSession(ctx, userID, userIP)
}

func (as *AuthorizationService) generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(as.PasswordSalt)))
}

func (as *AuthorizationService) CreateCredentials(ctx context.Context, username, password, userIP string) (Credentials, error) {
	user, err := as.Repository.GetUser(ctx, username, as.generateHash(password))
	if err != nil {
		return Credentials{}, err
	}

	return as.updateUserSession(ctx, user.ID)
}

type Credentials struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (as *AuthorizationService) RefreshCredentials(ctx context.Context, token, refreshToken, userIP string) (Credentials, error) {
	userID, err := as.TokenManager.ParseToken(token)
	if err != nil {
		if err.Error() != "Token is expired" {
			log.Println(err)
			return Credentials{}, err
		}
	}

	userSession, err := as.Repository.GetUserSession(ctx, userID, refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return Credentials{}, RefreshTokenExpired
		}
		log.Println(err)
	}

	if userSession.RefreshTokenTTL > time.Now().UnixNano() && userSession.UserIP == userIP {
		log.Println(RefreshTokenExpired)
		return Credentials{}, RefreshTokenExpired
	}

	return as.updateUserSession(ctx, userID)
}

func (as *AuthorizationService) updateUserSession(ctx context.Context, userID int) (Credentials, error) {
	token, err := as.TokenManager.NewToken(userID)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	refreshToken := as.TokenManager.NewRefreshToken()
	refreshTokenTTL := as.TokenManager.CreateRefreshTokenTTL()
	if _, err = as.Repository.UpdateUserSession(ctx, userID, refreshToken, refreshTokenTTL); err != nil {
		return Credentials{}, err
	}

	return Credentials{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AuthorizationService) createUserSession(ctx context.Context, userID int, userIP string) (Credentials, error) {
	token, err := as.TokenManager.NewToken(userID)
	if err != nil {
		log.Println(err)
		return Credentials{}, err
	}

	refreshToken := as.TokenManager.NewRefreshToken()
	refreshTokenTTL := as.TokenManager.CreateRefreshTokenTTL()
	if _, err = as.Repository.CreateUserSession(ctx, userID, userIP, refreshToken, refreshTokenTTL); err != nil {
		return Credentials{}, err
	}

	return Credentials{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
