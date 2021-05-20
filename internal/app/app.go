package app

import (
	"errors"
	"golang-web-app/internal/repository"
	"golang-web-app/internal/server"
	"golang-web-app/internal/service"
	trpHTTP "golang-web-app/internal/transport/http"
	"golang-web-app/pkg/database/postgresql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config := postgresql.NewConfig(
		os.Getenv("POSTGRESQL_USERNAME"),
		os.Getenv("POSTGRESQL_PASSWORD"),
		os.Getenv("POSTGRESQL_HOST"),
		os.Getenv("POSTGRESQL_PORT"),
		os.Getenv("POSTGRESQL_DB_NAME"),
		os.Getenv("POSTGRESQL_SSL_MODE"),
	)

	repository := repository.NewRepository(postgresql.NewClient(config.GetConnectionString()))
	service := service.NewService(service.Dependencies{
		Repository: repository,
	})
	newHTTP := trpHTTP.NewHTTP(service)
	srv := server.NewServer(os.Getenv("PORT"), newHTTP.Init())
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Println(err.Error())
	}
}
