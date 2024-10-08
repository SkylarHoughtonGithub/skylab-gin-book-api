package main

import (
	"fmt"
	"log"

	"skylab-gin-book-api/internal/config"
	"skylab-gin-book-api/internal/database"
	"skylab-gin-book-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if err := setupEnvironment(cfg); err != nil {
		return fmt.Errorf("failed to setup environment: %w", err)
	}

	db, err := setupDatabase(cfg)
	if err != nil {
		return fmt.Errorf("failed to setup database: %w", err)
	}

	r := routes.SetupRouter(db)

	log.Printf("Starting server on :%d", cfg.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func setupEnvironment(cfg *config.Config) error {
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	setLogLevel(cfg.Log.Level)
	return nil
}

func setupDatabase(cfg *config.Config) (*database.DB, error) {
	db, err := database.NewDB(cfg.Database.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := database.CreateTable(db); err != nil {
		return nil, fmt.Errorf("error creating table: %w", err)
	}
	fmt.Println("Table created or already exists.")

	return db, nil
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	case "info":
		log.SetFlags(log.Ldate | log.Ltime)
	default:
		log.SetFlags(0)
	}
}
