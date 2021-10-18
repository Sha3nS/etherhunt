package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

var (
	ChainType = "MainNet"
	Config = HostConfig{}
)

var (
	configFile = "./config.yml"

)

type HostConfig struct {
	InvokeHeight 	uint64
	ContractAddress []byte

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