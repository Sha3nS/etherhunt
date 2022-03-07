package eth_contract_watcher

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
)

const (
	MainNet = "MainNet"
	Ropsten = "Ropsten"
	Rinkeby = "Rinkeby"
	Goerli 	= "Goerli"
)

func GetChainConfig(chainType string) (*params.ChainConfig, error) {
	switch chainType {
	case MainNet:
		return params.MainnetChainConfig, nil
	case Ropsten:
		return params.RopstenChainConfig, nil
	case Rinkeby:
		return params.RinkebyChainConfig, nil
	case Goerli:
		return params.GoerliChainConfig, nil
	default:
		return nil, fmt.Errorf("unknown chaintype: %s, surrently suppport:\n MainNet、Ropsten、Rinkeby、Goerli", chainType)
	}
}


