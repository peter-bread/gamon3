package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
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

type GHHosts struct {
	GitHubCom struct {
		User string `yaml:"user"`
	} `yaml:"github.com"`
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

func (g *GHHosts) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, g); err != nil {
		return err
	}
	return nil
}

func getGHConfigPath() (string, error) {
	if configDir, found := os.LookupEnv("GH_CONFIG_DIR"); found {
		return filepath.Join(configDir, "hosts.yml"), nil
	}

	if xdgConfigDir, found := os.LookupEnv("XDG_CONFIG_HOME"); found {
		return filepath.Join(xdgConfigDir, "gh", "hosts.yml"), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "gh", "hosts.yml"), nil
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

	var ghHosts GHHosts
	ghHostsPath, err := getGHConfigPath()
	if err != nil {
		log.Fatalln(err)
	}

	if err := ghHosts.Load(ghHostsPath); err != nil {
		log.Fatalln(err)
	}

	currentAccount := ghHosts.GitHubCom.User
	fmt.Println("Current: ", currentAccount)

	if account != currentAccount {
		// cmd := exec.Command("gh", "auth", "switch", "--user", account).Run()
		cmd := exec.Command("echo", "gh", "auth", "switch", "--user", account)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func main() {
	Setup()
	Run()
}
