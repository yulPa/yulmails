package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"

	"github.com/yulPa/yulmails/logger"
)

var log = logger.GetLogger()

type configuration struct {
	S services `yaml:"services"`
	V string   `yaml:"version"`
}

type services struct {
	Archiving   optsArchiving `yaml:"archiving_db"`
	Senders     []node        `yaml:"senders"`
	Computes    []node        `yaml:"computes"`
	Entrypoints []node        `yaml:"entrypoints"`
}

type optsArchiving struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type node struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfigurationFromConfFile() *configuration {

	/*
	  Create a new Configuration from a conf file
	  return: <Configuration> A check mail configuration
	*/
	var conf configuration
	absFilePath, _ := filepath.Abs("yulmails.yaml")
	raw, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		log.Errorln(err)
	}
	yaml.Unmarshal(raw, &conf)
	return &conf
}
