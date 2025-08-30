package setup

import (
	"log"
	"peter-bread/gamon3/internal/ghswitch"
)

// Setup reads YAML config file and creates a JSON mapping.
func Setup() {
	var config ghswitch.Config
	loadPath := "examples/config.yaml"
	if err := config.Load(loadPath); err != nil {
		log.Fatalln(err)
	}

	var writeMap ghswitch.Mapping
	savePath := "examples/mapping.json"
	writeMap.Create(&config)
	if err := writeMap.Save(savePath); err != nil {
		log.Fatalln(err)
	}
}
