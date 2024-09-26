package main

import (
	"log"
	"skylab-book-chameleon/internal/config"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupEnvironment(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *config.Config
		wantMode string
	}{
		{
			name:     "Debug_mode",
			cfg:      &config.Config{Server: config.ServerConfig{Debug: true}},
			wantMode: gin.DebugMode,
		},
		{
			name:     "Release_mode",
			cfg:      &config.Config{Server: config.ServerConfig{Debug: false}},
			wantMode: gin.ReleaseMode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := setupEnvironment(tt.cfg)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantMode, gin.Mode())
		})
	}
}

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  int
	}{
		{"Debug_level", "debug", log.Ldate | log.Ltime | log.Lshortfile},
		{"Info_level", "info", log.Ldate | log.Ltime},
		{"Default_level", "error", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setLogLevel(tt.level)
			assert.Equal(t, tt.want, log.Flags())
		})
	}
}
