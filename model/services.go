package model

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	// ServicesFilePath is path of config file of services
	ServicesFilePath = "services.yaml"
	// Services contains list of all services
	Services ServicesListYAML
)

// Service is container of properties of services
type Service struct {
	Name          string `yaml:"name"`
	Request       string `yaml:"request-url"`
	POST          bool   `yaml:"post"`
	Timeout       int    `yaml:"timeout"`
	Period        int    `yaml:"period"`
	ContentType   string `yaml:"content-type"`
	RequestsCount int    `yaml:"requests-count"`
	FailureCount  int    `yaml:"failures-count"`
}

// ServicesListYAML contains array of services
type ServicesListYAML struct {
	ServicesList []Service `yaml:"services"`
}

// LoadServices loads services from file
func LoadServices() error {
	var servicesYaml ServicesListYAML
	servicesFile, err := ioutil.ReadFile(ServicesFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(servicesFile, &servicesYaml)
	if err != nil {
		return err
	}

	Services = servicesYaml
	return nil
}
