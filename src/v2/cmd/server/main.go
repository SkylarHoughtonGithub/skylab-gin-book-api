// cmd/server/main.go

package main

import (
	"fmt"
	"log"

	"skylab/internal/config"
	"skylab/internal/database"
	"skylab/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode based on debug setting
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set log level
	setLogLevel(cfg.Log.Level)

	db, err := database.NewDB(cfg.Database.ConnectionString())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := routes.SetupRouter(db)

	log.Printf("Starting server on :%d", cfg.Server.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
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
