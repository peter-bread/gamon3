package config

// This file contains extra code to validate the GitHubCom struct.
// It ensures that the User field value is a key in the Users field value.

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

type GitHubCom struct {
	Users map[string]any `yaml:"users" validate:"required"`
	User  string         `yaml:"user" validate:"required"`
}

type GhHosts struct {
	GitHubCom GitHubCom `yaml:"github.com" validate:"required"`
}

type LocalConfig struct {
	Account string `yaml:"account" validate:"required"`
}

type MainConfig struct {
	Default  string              `yaml:"default" validate:"required"`
	Accounts map[string][]string `yaml:"accounts"`
}

type configType interface {
	GhHosts | LocalConfig | MainConfig
}

var validate *validator.Validate

func githubComStructValidation(sl validator.StructLevel) {
	cfg := sl.Current().Interface().(GitHubComConfig)

	if cfg.Users == nil {
		return
	}

	if _, ok := cfg.Users[cfg.User]; !ok {
		sl.ReportError(
			cfg.User,
			"User",
			"user",
			"usernotfound",
			"",
		)
	}
}

func load[T configType](path string, strict bool) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading %v: %w", path, err)
	}

	var cfg T

	validate = validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterStructValidation(githubComStructValidation, GitHubCom{})

	opts := []yaml.DecodeOption{yaml.Validator(validate)}
	if strict {
		opts = append(opts, yaml.Strict())
	}

	err = yaml.UnmarshalWithOptions(data, &cfg, opts...)
	if err != nil {
		return nil, fmt.Errorf("error while decoding %v: %v", path, yaml.FormatError(err, true, true))
	}

	return &cfg, nil
}

func LoadGhHosts(path string) (*GhHosts, error) {
	cfg, err := load[GhHosts](path, false)
	if err != nil {
		return nil, fmt.Errorf("failed to load gh hosts: %w", err)
	}
	return cfg, nil
}

func LoadLocalConfig(path string) (*LocalConfig, error) {
	cfg, err := load[LocalConfig](path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load local config: %w", err)
	}
	return cfg, nil
}

func LoadMainConfig(path string) (*MainConfig, error) {
	cfg, err := load[MainConfig](path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load main config: %w", err)
	}
	return cfg, nil
}
