package model

var (
	ServicesFailures    = make(map[string]int)
	ServicesRequests    = make(map[string]int)
	ServicesHealthiness = make(map[string]float32)
)
