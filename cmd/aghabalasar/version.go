package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aryahadii/AghaBalaSar"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of Agha Balasar",
	Run: func(cmd *cobra.Command, args []string) {
		logVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func logVersion() {
	log.Info("version   > ", aghabalasar.Version)
	log.Info("buildtime > ", aghabalasar.BuildTime)
	log.Info("commit    > ", aghabalasar.Commit)
}
