package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
)

// Config stores config data. This data is read from a YAML file.
type Config struct {
	Accounts map[string][]string `yaml:"accounts"`
	Default  string              `yaml:"default"`
}

type Mapping struct {
	Paths   map[string]string `json:"paths"`
	Default string            `json:"default"`
}

func normalise(path string) string {
	// TODO: Use `filepath.Abs`?
	return filepath.Clean(os.ExpandEnv(path))
}

// MapAccounts runs the function `f` on all acounts in a `Config`.
// This function modifies `Config` in-place.
func (c *Config) MapAccounts(f func(string) string) {
	for account, paths := range c.Accounts {
		for i, p := range paths {
			paths[i] = f(p)
		}
		c.Accounts[account] = paths
	}
}

// Load reads `Config` data from a YAML file located at `path`.
func (c *Config) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) printYAML() {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}

func (c *Config) printRaw() {
	fmt.Println(c)
}

func (m *Mapping) Create(c *Config) {
	m.Default = c.Default
	m.Paths = make(map[string]string)

	for account, paths := range c.Accounts {
		for _, path := range paths {
			m.Paths[normalise(path)] = account
		}
	}
}

func (m *Mapping) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(m)
}

func (m *Mapping) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, m); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) GetAccount(dir string) string {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return m.Default
	}

	if account, ok := m.Paths[abs]; ok {
		return account
	}

	for path, account := range m.Paths {
		if strings.HasPrefix(abs, path) {
			return account
		}
	}

	return m.Default
}

// Setup reads YAML config file and creates a JSON mapping.
func Setup() {
	var config Config
	loadPath := "examples/config.yaml"
	if err := config.Load(loadPath); err != nil {
		log.Fatalln(err)
	}

	var writeMap Mapping
	savePath := "examples/mapping.json"
	writeMap.Create(&config)
	if err := writeMap.Save(savePath); err != nil {
		log.Fatalln(err)
	}
}

// Run reads from a JSON mapping file and determines the account based on
// current working directory.
func Run() {
	var readMap Mapping
	readPath := "examples/mapping.json"
	if err := readMap.Load(readPath); err != nil {
		log.Fatalln(err)
	}

	pwd := os.Getenv("PWD")
	account := readMap.GetAccount(pwd)

	// TODO: get current account from ~/.config/gh/hosts.yml
	// Compare this to `account`. If they are different, then
	// use exec.Command to switch account.

	fmt.Println("Switching to account:", account)
	// cmd := exec.Command("echo", "gh", "auth", "switch", "--user", account)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()
}

func main() {
	Setup()
	Run()
}
