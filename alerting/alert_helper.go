package alerting

import (
	"bytes"
	"fmt"
	"time"

	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/aryahadii/AghaBalaSar/model"
)

const (
	slackPostData = `{"text": "%s"}`
)

func alertSlack(service *model.Service) {
	log.Info(service.Name, " (SLACK)")

	service.LastAlert = time.Now()

	for _, webhookURL := range service.Slacks {
		bodyString := fmt.Sprintf(slackPostData, service.Name+" is not OK")
		body := bytes.NewBufferString(bodyString)
		resp, err := http.Post(webhookURL, "application/json", body)
		if err != nil {
			log.WithError(err).Error(fmt.Sprintf("Can't alert by Slack (%s)", webhookURL))
		}
		if resp.StatusCode != 200 {
			log.Error(fmt.Sprintf("Can't alert by Slack. %s returns %d", webhookURL, resp.StatusCode))
		}
	}
}

func alertWebhook(service *model.Service) {
	log.Info(service.Name, " (WEBHOOK)")
	service.LastAlert = time.Now()
}
