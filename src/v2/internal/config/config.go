// internal/config/config.go

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Log      LogConfig
}

type ServerConfig struct {
	Port  int
	Debug bool
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type LogConfig struct {
	Level string
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
