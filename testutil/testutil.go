package testutil

import (
	"encoding/json"
	"os"

	"github.com/BaristaVentures/errand-boy/config"
)

// CreateConfigFile creates a dummy config file.
func CreateConfigFile(conf *config.Config, configPath string) error {
	configBytes, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	file.Write(configBytes)
	file.Close()
	return nil
}
