package configuration

import (
	"github.com/spf13/viper"
)

var (
	// ConfigFilePath is path of config
	ConfigFilePath = "config.yaml"
	// AghabalasarConfig is main config of service
	AghabalasarConfig viper.Viper
)

// LoadConfig loads configs from ConfigFilePath
func LoadConfig() error {
	config := viper.New()
	config.SetConfigFile(ConfigFilePath)

	if err := config.ReadInConfig(); err != nil {
		return err
	}

	config.SetDefault("address", "localhost:8000")
	config.SetDefault("debug", true)

	AghabalasarConfig = *config
	return nil
}
