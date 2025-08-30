package setup

import (
	"log"
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

	// TODO: Create directory if not exists.
	// e.g. $HOME/.config/gamon3

	if err := config.Load(configPath); err != nil {
		log.Fatalln(err)
	}

	// TODO: Create directory if not exists.
	// e.g. $HOME/.local/state/gamon3

	mappingPath, err := ghswitch.GetMappingPath()
	if err != nil {
		log.Fatalln(err)
	}

	mapping.Create(&config)
	if err := mapping.Save(mappingPath); err != nil {
		log.Fatalln(err)
	}
}
