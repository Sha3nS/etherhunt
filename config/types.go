package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ConfigPath 	= "./config.yml"
	Config 		= HostConfig{}
)

type HostConfig struct {
	ChainType		string		`yaml:"ChainType"`
	URL				string		`yaml:"URL"`
	InvokeHeight 	uint64		`yaml:"Height"`
	ContractAddress string		`yaml:"Address"`
	ContractParams 	[]string 	`yaml:"Params"`
	PrivateKeyPath  string		`yaml:"KeyPath"`
}

func InitConfig() error {
	fileBz, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileBz, &Config)
	if err != nil {
		return err
	}
	return nil
}
