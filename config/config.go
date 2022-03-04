package config

import (
	_ "embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	strutil "github.com/harmony-development/legato/util"
	"gopkg.in/yaml.v2"
)

// RedisConfig is the config for the Redis database.
type RedisConfig struct {
	URL string `yaml:"url" validate:"required"`
}

type PostgresConfig struct {
	URL string `vaml:"url"`
}

type AuthConfig struct {
	DisableRegistration bool `yaml:"disableRegistration"`
	TokenLength         int  `yaml:"tokenLength" validate:"required,min=4"`
}

// Config is the root config structure.
type Config struct {
	Port         int16          `yaml:"port" validate:"required,gte=1,lte=65535"`
	UseLocalCORS bool           `yaml:"useLocalCORS" validate:"required"`
	CleanLogs    bool           `yaml:"cleanLogs" validate:"required"`
	Redis        RedisConfig    `yaml:"redis" validate:"required"`
	Postgres     PostgresConfig `yaml:"postgres" validate:"required"`
	Auth         AuthConfig     `yaml:"auth" validate:"required"`
}

//go:embed default-config.yml
var defaultConfig string

var (
	errCreatedDefaultConfig = errors.New("config file not found, default config created")
	errValidationFailed     = errors.New("config validation failed")
)

// ReadConfig reads the config file and returns a Config struct.
func ReadConfig() (*Config, error) {
	path := filepath.Clean(strutil.FirstNonEmpty(os.Getenv("CONFIG_PATH"), "./config.yml"))

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
			acc += fmt.Sprintf("%s failed to meet requirement %s %s\n", err.Namespace(), err.Tag(), err.Param())
		}

		return nil, fmt.Errorf("%w: \n%s", errValidationFailed, acc)
	}

	return &cfg, nil
}
