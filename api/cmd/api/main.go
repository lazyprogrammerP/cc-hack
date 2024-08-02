package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/lazyprogrammerP/cc-hack/pkg/logger"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Error().Err(err).Msg("failed to load the environment variables")
		return
	}

	// Read the environment variables
	PORT := os.Getenv("PORT")

	r := chi.NewRouter()

	// CORS Configuration
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}

	r.Use(cors.Handler(corsOptions))

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Heartbeat("/"))

	// Register the routes
	// ...

	addr := fmt.Sprintf(":%s", PORT)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info().Msg("starting the server on port " + PORT)
		if err := srv.ListenAndServe(); err != nil {
			logger.Error().Err(err).Msg("failed to start the server")
		}
	}()

	<-interrupt

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger.Info().Msg("shutting down the server")
	if err = srv.Shutdown(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to shut down the server. forcing shutdown")
	}

	logger.Info().Msg("server shut down successfully")
}
