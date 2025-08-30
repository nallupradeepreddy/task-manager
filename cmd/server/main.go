package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nallupradeepreddy/task-manager/internal/adapters/api"
	"github.com/nallupradeepreddy/task-manager/internal/adapters/repository"
	"github.com/nallupradeepreddy/task-manager/internal/core/service"
)

func main() {

	// Set up Viper to read from .env file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Err(err).Msg("No .env file found, using environment variables only")
	}

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	// Set up Zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// Database connection using GORM
	dbURL := viper.GetString("DATABASE_URL")
	if dbURL == "" {
		log.Fatal().Msg("DATABASE_URL not set in .env file")
	}
	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database with GORM")
	}

	// Wire up repository, service, and handler
	userRepo := repository.NewUserPostgresRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)
	userHandler.RegisterRoutes(r)

	// Example route
	r.GET("/ping", func(c *gin.Context) {
		log.Info().Msg("Received /ping request")
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Protected example route
	r.GET("/protected", api.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authorized!"})
	})

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	go func() {
		log.Info().Msg("Starting server on :8080")
		serverErrors <- srv.ListenAndServe()
	}()

	// Listen for interrupt signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Error().Err(err).Msg("Server error")
	case sig := <-quit:
		log.Info().Str("signal", sig.String()).Msg("Shutting down server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	} else {
		log.Info().Msg("Server exited gracefully")
	}
}
