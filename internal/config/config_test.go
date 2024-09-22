// internal/config/config.go

package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name    string  `json:"name,omitempty"`
		want    *Config `json:"want,omitempty"`
		wantErr bool    `json:"want_err,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseConfig_ConnectionString(t *testing.T) {
	tests := []struct {
		name string          `json:"name,omitempty"`
		c    *DatabaseConfig `json:"c,omitempty"`
		want string          `json:"want,omitempty"`
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ConnectionString(); got != tt.want {
				t.Errorf("DatabaseConfig.ConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
