package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configuration struct {
	Sql Sql `yaml:"Sql"`
}

type Sql struct {
	Port     int    `yaml:"Port"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
	User     string `yaml:"User"`
	Database string `yaml:"Db"`
	Password string `yaml:"Pass"`
}

func (c *Configuration) GetConf() *Configuration {

	pwd, _ := os.Getwd()
	yamlFile, err := os.ReadFile(pwd + "/config/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
