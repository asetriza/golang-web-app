package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asetriza/golang-web-app/internal/repository"
	"github.com/asetriza/golang-web-app/internal/server"
	"github.com/asetriza/golang-web-app/internal/service"
	"github.com/asetriza/golang-web-app/internal/transport/rest"
	"github.com/asetriza/golang-web-app/pkg/auth"
	"github.com/asetriza/golang-web-app/pkg/database/postgresql"

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

	conn, err := postgresql.NewClient(config.GetConnectionString())
	if err != nil {
		return err
	}

	Repository := repository.NewRepository(conn)

	Service := service.NewService(service.Dependencies{
		Repository:   Repository,
		TokenManager: tokenManager,
		PasswordSalt: os.Getenv("PASSWORD_SALT"),
	})

	REST := rest.NewREST(Service, tokenManager)

	srv := server.NewServer(os.Getenv("PORT"), REST.Router())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()

	log.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}

	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
