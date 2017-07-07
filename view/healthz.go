package view

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aryahadii/AghaBalaSar"
	"github.com/aryahadii/AghaBalaSar/configuration"
)

const healthzTemplate = `
Version: %s
Build Time: %s
Commit: %s
Uptime: %s
`

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("I'm OK"))
	if !configuration.AghabalasarConfig.GetBool("debug") {
		return
	}

	uptime := time.Since(aghabalasar.StartTime)
	w.Write([]byte(fmt.Sprintf(healthzTemplate,
		aghabalasar.Version, aghabalasar.BuildTime, aghabalasar.Commit, uptime)))
}
