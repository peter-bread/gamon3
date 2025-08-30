package setup

import (
	"fmt"
	"log"
	"peter-bread/gamon3/internal/ghswitch"
)

var (
	config  ghswitch.Config
	mapping ghswitch.Mapping
)

// SetupCmd reads YAML config file and creates a JSON mapping.
func SetupCmd() {
	// TODO: Read from proper location (XDG_CONFIG).
	configPath, err := ghswitch.GetConfigPath()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(configPath)

	loadPath := "examples/config.yaml"
	if err := config.Load(loadPath); err != nil {
		log.Fatalln(err)
	}

	// TODO: Write to proper location (XDG_STATE).
	mappingPath, err := ghswitch.GetMappingPath()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(mappingPath)

	savePath := "examples/mapping.json"
	mapping.Create(&config)
	if err := mapping.Save(savePath); err != nil {
		log.Fatalln(err)
	}
}
