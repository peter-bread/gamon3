package setup

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"peter-bread/gamon3/internal/ghswitch"
)

var (
	config  ghswitch.Config
	mapping ghswitch.Mapping
)

// SetupCmd reads YAML config file and creates a JSON mapping.
func SetupCmd() {
	configPath, err := ghswitch.GetConfigPath()
	if err != nil {
		log.Fatalln(err)
	}

	if err := config.Load(configPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mappingPath, err := ghswitch.GetMappingPath()
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.MkdirAll(filepath.Dir(mappingPath), 0755); err != nil {
		log.Fatalln(err)
	}

	mapping.Create(&config)
	if err := mapping.Save(mappingPath); err != nil {
		log.Fatalln(err)
	}
}
