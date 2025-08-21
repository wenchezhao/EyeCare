package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Config
type Config struct {
	RemindInterval         int      `json:"remind_interval"`
	RestDuration           int      `json:"rest_duration"`
	EnableFullscreenIgnore bool     `json:"enable_fullscreen_ignore"`
	EnableProcessWhitelist bool     `json:"enable_process_whitelist"`
	ProcessWhitelist       []string `json:"process_whitelist"`
	DarkTheme              bool     `json:"dark_theme"`
}

// DefaultConfig
func DefaultConfig() *Config {
	return &Config{
		RemindInterval:         20,
		RestDuration:           20,
		EnableFullscreenIgnore: true,
		EnableProcessWhitelist: false,
		ProcessWhitelist:       []string{},
		DarkTheme:              true,
	}
}

// GetConfigPath
func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	configDir := filepath.Join(homeDir, ".eyecare")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return "config.json"
	}
	return filepath.Join(configDir, "config.json")
}

// LoadConfig
func LoadConfig() (*Config, error) {
	configPath := GetConfigPath()

	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			cfg := DefaultConfig()
			if err := SaveConfig(cfg); err != nil {
				return nil, err
			}
			return cfg, nil
		}
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// SaveConfig
func SaveConfig(cfg *Config) error {
	configPath := GetConfigPath()

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(cfg)
}

// GetRemindInterval
func (c *Config) GetRemindInterval() time.Duration {
	return time.Duration(c.RemindInterval) * time.Minute
}

// GetRestDuration
func (c *Config) GetRestDuration() time.Duration {
	return time.Duration(c.RestDuration) * time.Second
}
