package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aryahadii/AghaBalaSar/configuration"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aghabalasar <subcommand>",
	Short: "Health-check Service",
	Long: `Aghabalasar can be used as a health-checker for http services
			More info at https://github.com/aryahadii/AghaBalaSar`,
	Run: nil,
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVarP(&configuration.ConfigFilePath,
		"config", "c", configuration.ConfigFilePath, "Path to the config file (e.g. ./config.yaml)")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		err := configuration.LoadConfig()
		if err != nil {
			log.WithError(err).Fatalf("Failed to load the config file")
		}
	}
}
