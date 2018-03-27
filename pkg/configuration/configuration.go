package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"

	"github.com/yulPa/yulmails/pkg/logger"
	"github.com/yulPa/yulmails/pkg/mongo"
)

var log = logger.GetLogger("configuration-ym")

type Configuration struct {
	S Services `yaml:"services",json:"services"`
	V string   `yaml:"version",json:"version"`
}

type Services struct {
	Archiving   OptsArchiving `yaml:"archiving_db",json:"archiving_db"`
	Senders     []Node        `yaml:"senders",json:"senders"`
	Computes    []Node        `yaml:"computes",json:"computes"`
	Entrypoints []Node        `yaml:"entrypoints",json:"entrypoints"`
}

type OptsArchiving struct {
	Name     string `yaml:"name",json:"name"`
	Host     string `yaml:"host",json:"host"`
	Port     string `yaml:"port",json:"port"`
	Username string `yaml:"username",json:"username"`
	Password string `yaml:"password",json:"password"`
}

type Node struct {
	Name string `yaml:"name",json:"name"`
	Host string `yaml:"host",json:"host"`
	Port string `yaml:"port",json:"port"`
}

func NewConfigurationFromConfFile(session mongo.Session) error {
	/*
	  Create a new Configuration from a conf file
	  return: <Configuration> A check mail configuration
	*/
	var conf Configuration
	absFilePath, _ := filepath.Abs("./conf/yulmails.yaml")
	raw, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		log.Errorln(err)
	}
	yaml.Unmarshal(raw, &conf)
	return conf.saveConfiguration(session)
}

func (this Configuration) saveConfiguration(session mongo.Session) error {
	/*
		This private method will save configuration into an archivingdb
	*/

	sess := session.Copy()
	defer sess.Close()
	db := sess.DB("global")
	col := db.C("configuration")

	err := col.Insert(this)

	return err
}
