package framework

import (
	"encoding/json"
	"os"
)

// Config represents the configuration for the framework
type Config struct {
	Routes    []Route        `json:"routes"`
	Static    StaticConfig   `json:"static"`
	Templates TemplateConfig `json:"templates"`
}

// Route represents a route configuration
type Route struct {
	Path     string                 `json:"path"`
	Template string                 `json:"template"`
	Title    string                 `json:"title"`
	Data     map[string]interface{} `json:"data"`
}

// StaticConfig represents the configuration for static file serving
type StaticConfig struct {
	Path string `json:"path"`
	Dir  string `json:"dir"`
}

// TemplateConfig represents the configuration for templates
type TemplateConfig struct {
	Common string `json:"common"`
	Pages  string `json:"pages"`
}

// LoadConfig loads the configuration from a JSON file
func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
