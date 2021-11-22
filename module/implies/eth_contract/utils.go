package eth_contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shawncles/etherhunt/config"
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
		return nil, fmt.Errorf("unknown chaintype: %s, surrently suppport:\n MainNet、Ropsten、Rinkeby、Goerli", config.ChainType)
	}
}


