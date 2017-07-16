package alerting

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/aryahadii/AghaBalaSar/model"
	"github.com/jasonlvhit/gocron"
)

// StartScheduler starts health checking for all services
func StartScheduler() {
	for i, service := range model.Services.ServicesList {
		model.Services.ServicesList[i].LastRequestsResult = make([]bool, service.RequestsCount)
		for j := 0; j < service.RequestsCount; j++ {
			model.Services.ServicesList[i].LastRequestsResult[j] = true
		}

		gocron.Every(uint64(service.Period/service.RequestsCount)).Seconds().Do(healthCheck, &model.Services.ServicesList[i])
		model.ServicesHealthiness[service.Name] = 1
	}
	<-gocron.Start()
}

func healthCheck(service *model.Service) {
	// Compute healthiness
	var okRequests int
	for i := 0; i < service.RequestsCount; i++ {
		if service.LastRequestsResult[i] == true {
			okRequests++
		}
	}
	model.ServicesHealthiness[service.Name] = float32(okRequests) / float32(service.RequestsCount)

	// Alert if it's not OK
	if okRequests <= service.RequestsCount-service.FailureCount {
		alert(service)
	}

	if service.POST {
		// TODO Support POST
	} else {
		timeout := time.Duration(service.Timeout * int(time.Second))
		client := http.Client{
			Timeout: timeout,
		}

		resp, err := client.Get(service.Request)
		if err != nil {
			submitFailedRequest(service)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			submitFailedRequest(service)
			return
		}

		submitSucceedRequest(service)
	}
}

func submitFailedRequest(service *model.Service) {
	service.LastRequestsResult = append(service.LastRequestsResult[1:], false)
}

func submitSucceedRequest(service *model.Service) {
	service.LastRequestsResult = append(service.LastRequestsResult[1:], true)
}

func alert(service *model.Service) {
	log.Info(service.Name, " (ALERT!)")

	if time.Since(service.LastAlert) > time.Minute*10 {
		if len(service.Slacks) > 0 {
			alertSlack(service)
		}

		if len(service.Webhooks) > 0 {
			alertSlack(service)
		}
	}
}
