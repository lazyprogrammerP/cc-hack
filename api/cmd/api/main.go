package main

import (
	"context"
	"database/sql"
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
	"github.com/lazyprogrammerP/cc-hack/internal/routes"
	"github.com/lazyprogrammerP/cc-hack/internal/sqlc"
	"github.com/lazyprogrammerP/cc-hack/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("failed to load the environment variables")
		return
	}

	// Read the environment variables
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	DSN := os.Getenv("DSN")
	if DSN == "" {
		log.Error().Msg("DSN is not set")
		return
	}

	DB_DRIVER := os.Getenv("DB_DRIVER")
	if DB_DRIVER == "" {
		log.Error().Msg("DB_DRIVER is not set")
		return
	}

	// Create an instance of the database
	db, err := sql.Open(DB_DRIVER, DSN)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to the database")
		return
	}

	// Check if the database is alive
	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msg("failed to ping the database")
		return
	}

	queries := sqlc.New(db)

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
	r.Mount("/department", routes.RegisterDepartmentRoutes(queries))

	addr := fmt.Sprintf(":%s", PORT)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info().Msg("starting the server on port " + PORT)
		if err := srv.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("failed to start the server")
		}
	}()

	<-interrupt

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Info().Msg("shutting down the server")
	if err = srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shut down the server. forcing shutdown")
	}

	log.Info().Msg("server shut down successfully")
}
