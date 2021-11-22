package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ChainType 	= "MainNet"
	configFile 	= "./config.yml"
	Config 		= HostConfig{}
)

type HostConfig struct {
	InvokeHeight 	uint64		`yaml:"Height"`
	ContractAddress []byte		`yaml:"Address"`
	ContractParams 	[]string 	`yaml:"Params"`
	PrivateKeyPath  string		`yaml:"KeyPath"`
}

func InitConfig() error {
	fileBz, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileBz, Config)
	if err != nil {
		return err
	}
	return nil
}
