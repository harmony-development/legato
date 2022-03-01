package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	_ "embed"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port int16 `yaml:"port" validate:"required,gte=1,lte=65535"`
}

//go:embed default-config.yml
var defaultConfig string

var (
	errCreatedDefaultConfig = errors.New("config file not found, default config created")
	errValidationFailed     = errors.New("config validation failed")
)

// ReadConfig reads the config file and returns a Config struct.
func ReadConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := ioutil.WriteFile(path, []byte(defaultConfig), 0o600); err != nil {
				return nil, fmt.Errorf("failed to write default config: %w", err)
			}

			return nil, errCreatedDefaultConfig
		}

		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if err := validator.New().Struct(cfg); err != nil {
		var validationErr validator.ValidationErrors

		errors.As(err, &validationErr)

		acc := ""
		for _, err := range validationErr {
			acc += fmt.Sprintf("%s failed to meet requirement %s %s\n", err.Field(), err.Tag(), err.Param())
		}

		return nil, fmt.Errorf("%w: %s", errValidationFailed, acc)
	}

	return &cfg, nil
}
