package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerName string `yaml:"server_name"`
	Address    string `yaml:"address"`
}

var config = &Config{}

func Setup() {
	f, err := ioutil.ReadFile("./conf/conf.yml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(f, config); err != nil {
		panic(err)
	}
}

func GetThemisServer() (string, string) {
	return config.ServerName, config.Address
}
