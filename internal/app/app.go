package app

import (
	"errors"
	"golang-web-app/internal/repository"
	"golang-web-app/internal/server"
	"golang-web-app/internal/service"
	trpHTTP "golang-web-app/internal/transport/http"
	"golang-web-app/pkg/auth"
	"golang-web-app/pkg/database/postgresql"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	config, err := postgresql.NewConfig(
		os.Getenv("POSTGRESQL_USERNAME"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DB_NAME"),
		os.Getenv("POSTGRESQL_SSL_MODE"),
	)
	if err != nil {
		return err
	}

	tokenManager := auth.NewManager(os.Getenv("SIGN_IN_KEY"), os.Getenv("TOKEN_TTL"), os.Getenv("REFRESH_TOKEN_TTL"))

	newRepository := repository.NewRepository(postgresql.NewClient(config.GetConnectionString()))

	newService := service.NewService(service.Dependencies{
		Repository:   newRepository,
		TokenManager: tokenManager,
		PasswordSalt: os.Getenv("PASSWORD_SALT"),
	})

	newHTTP := trpHTTP.NewHTTP(newService, tokenManager)

	srv := server.NewServer(os.Getenv("PORT"), newHTTP.Router())

	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
