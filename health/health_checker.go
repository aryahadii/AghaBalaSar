package health

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/aryahadii/AghaBalaSar/model"
	"github.com/jasonlvhit/gocron"
)

// StartScheduler starts health checking for all services
func StartScheduler() {
	for _, service := range model.Services.ServicesList {
		gocron.Every(uint64(service.Period/service.RequestsCount)).Seconds().Do(healthCheck, service)
		model.ServicesHealthiness[service.Name] = 1
	}
	<-gocron.Start()
}

func healthCheck(service model.Service) {
	if model.ServicesRequests[service.Name] >= service.RequestsCount {
		if model.ServicesFailures[service.Name] >= service.FailureCount {
			alert(service)
		}

		currentHealthiness := 1.0 - float32(model.ServicesFailures[service.Name])/float32(model.ServicesRequests[service.Name])
		model.ServicesHealthiness[service.Name] = (model.ServicesHealthiness[service.Name] + currentHealthiness) / 2
		model.ServicesFailures[service.Name] = 0
		model.ServicesRequests[service.Name] = 0
	}

	if service.POST {
		// TODO Support POST
	} else {
		timeout := time.Duration(service.Timeout * int(time.Second))
		client := http.Client{
			Timeout: timeout,
		}

		resp, err := client.Get(service.Request)
		model.ServicesRequests[service.Name]++
		if err != nil {
			model.ServicesFailures[service.Name]++
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			model.ServicesFailures[service.Name]++
		}
	}
}

func alert(service model.Service) {
	log.Info(service.Name, " (ALERT!)")
}
