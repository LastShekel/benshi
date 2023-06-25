package worker

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Conf is configuration structure for both driver and worker
type Conf struct {
	DriverPort string `yaml:"driver-port"` // DriverPort is port for driver in local machine
	WorkerPort string `yaml:"worker-port"` // WorkerPort is port for first (others will be in next open port) worker in local machine
	Files      string `yaml:"files"`       // Files is path for intermediate and files folder
	Inputs     string `yaml:"inputs"`      // Inputs is path for inputs folder
}

func (c *Conf) getConf() *Conf {
	panic("implement me")
}

// LoadConfig loads config from main.yaml
func LoadConfig() Conf {
	var c Conf
	yamlFile, err := ioutil.ReadFile("../../configs/main.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
