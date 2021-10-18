package ether_contract_hunt

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/etherhunt/config"
)

const (
	MainNet = "MainNet"
	Ropsten = "Ropsten"
	Rinkeby = "Rinkeby"
	Goerli 	= "Goerli"
)

func GetChainConfig() (*params.ChainConfig, error) {
	switch config.ChainType {
	case MainNet:
		return params.MainnetChainConfig, nil
	case Ropsten:
		return params.RopstenChainConfig, nil
	case Rinkeby:
		return params.RinkebyChainConfig, nil
	case Goerli:
		return params.GoerliChainConfig, nil
	default:
		return nil, fmt.Errorf("unknown chain type: %s, please select from\n MainNet、Ropsten、Rinkeby、Goerli", config.ChainType)
	}
}


