package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

// GhHost represents data for a single GitHub CLI host.
//
// The only fields it requires are User and Users. Any other fields, for example
// git_protocol or oauth_token, are ignored.
type GhHost struct {
	Users map[string]any `yaml:"users" validate:"required"`
	User  string         `yaml:"user" validate:"required"`
	// GitProtocol string
	// OAuthToken string
}

// GhHosts represents a GitHub CLI hosts.yml.
//
// It is a mapping of hostnames to host data.
type GhHosts map[string]GhHost

// LocalConfig represents a local configuration that can override
// the [MainConfig].
//
// Account specifies the account that should be used.
type LocalConfig struct {
	Account string `yaml:"account" validate:"required"`
}

// MainConfig represents the primary application configuration.
//
// Default specifies the default account name.
// Accounts is a map from account names to lists of filepaths.
type MainConfig struct {
	Default  string              `yaml:"default" validate:"required"`
	Accounts map[string][]string `yaml:"accounts"`
}

type configType interface {
	GhHosts | LocalConfig | MainConfig
}

var validate *validator.Validate

func load[T configType](path string, strict bool) (*T, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading %v: %w", path, err)
	}

	var cfg T

	validate = validator.New(validator.WithRequiredStructEnabled())

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

// LoadGhHosts reads a GH CLI hosts.yml file at path and returns a validated
// [GhHosts].
//
// The file is decoded non-strictly, allowing fields that are not represented
// in [GhHosts]. The resulting struct is validated according to the struct tags.
// An error is returned if the file cannot be read, decoded, or validated.
func LoadGhHosts(path string) (*GhHosts, error) {
	// Does not use strict decoder as the [GhHosts] type only contains a subset of
	// the data that may be in a hosts.yml file.
	cfg, err := load[GhHosts](path, false)
	if err != nil {
		return nil, fmt.Errorf("failed to load gh hosts: %w", err)
	}
	return cfg, nil
}

// LoadLocalConfig reads the YAML file at path and returns a validated [LocalConfig].
//
// The file is decoded strictly: unknown fields are rejected.
// The resulting configuration is validated according to the struct tags.
// An error is returned if the file cannot be read, decoded, or validated.
func LoadLocalConfig(path string) (*LocalConfig, error) {
	cfg, err := load[LocalConfig](path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load local config: %w", err)
	}
	return cfg, nil
}

// LoadMainConfig reads the YAML file at path and returns a validated [MainConfig].
//
// The file is decoded in strict mode: unknown fields are rejected.
// The resulting configuration is validated according to the struct tags.
// An error is returned if the file cannot be read, decoded, or validated.
func LoadMainConfig(path string) (*MainConfig, error) {
	cfg, err := load[MainConfig](path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load main config: %w", err)
	}
	return cfg, nil
}
