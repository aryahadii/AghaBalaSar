package view

import (
	"net/http"
	"strconv"

	"github.com/aryahadii/AghaBalaSar/model"
)

func ServicesHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	for _, service := range model.Services.ServicesList {
		healthiness := strconv.Itoa(int(model.ServicesHealthiness[service.Name]*100)) + "%"
		w.Write([]byte(service.Name + ": " + healthiness + "\n"))
	}
}
