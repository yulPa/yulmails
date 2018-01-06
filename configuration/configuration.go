package configuration

import 	(
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "path/filepath"
  "fmt"
)
type Configuration struct {
  dbHost string `yaml:"db_host"`
  dbName string `yaml:"db_name"`
  DbUser string `yaml:"db_user"`
  dbPass string `yaml:"db_pass"`
  dbPort int `yaml:"db_port"`
}

func NewConfiguration(dbHost string, dbName string, dbUser string, dbPass string, dbPort int) *Configuration {
  /*
    Create a new configuration
    parameter: <string> Database hostname
    parameter: <string> Database name
    parameter: <string> Database username
    parameter: <string> Database password
    parameter: <int> Database port
    returm: <Configuration> A check mail configuration
  */
  return &Configuration{
    dbHost: dbHost,
    dbName: dbName,
    DbUser: dbUser,
    dbPass: dbPass,
    dbPort: dbPort,
  }
}

func NewConfigurationFromConfFile() *Configuration {

  /*
  Create a new Configuration from a conf file
  return: <Configuration> A check mail configuration
*/
  var conf Configuration
	absFilePath, _ := filepath.Abs("check_mail.yaml")
	raw, err := ioutil.ReadFile(absFilePath)
	if err != nil {
		fmt.Println(err)
	}
	yaml.Unmarshal(raw, &conf)
  return &conf
}
