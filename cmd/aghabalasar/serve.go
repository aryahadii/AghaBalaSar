package main

import (
	log "github.com/Sirupsen/logrus"

	"net/http"

	"github.com/aryahadii/AghaBalaSar/configuration"
	"github.com/aryahadii/AghaBalaSar/health"
	"github.com/aryahadii/AghaBalaSar/model"
	"github.com/aryahadii/AghaBalaSar/view"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	logVersion()

	err := model.LoadServices()
	if err != nil {
		log.WithError(err).Fatal("Can't load services")
	}

	go health.StartScheduler()

	http.HandleFunc("/", view.ServicesHealth)
	http.HandleFunc("/healthz", view.Healthz)
	log.Fatal(http.ListenAndServe(configuration.AghabalasarConfig.GetString("address"), nil))
}
