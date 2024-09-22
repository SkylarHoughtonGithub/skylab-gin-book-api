// internal/config/config.go

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `json:"server,omitempty"`
	Database DatabaseConfig `json:"database,omitempty"`
	Log      LogConfig      `json:"log,omitempty"`
}

type ServerConfig struct {
	Port  int  `json:"port,omitempty"`
	Debug bool `json:"debug,omitempty"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
	SSLMode  string `json:"ssl_mode,omitempty"`
}

type LogConfig struct {
	Level string `json:"level,omitempty"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
		c.Driver, c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}
