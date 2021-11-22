package eth_contract

import (
	"context"
	"github.com/shawncles/etherhunt/module/clients"
	ctypes "github.com/shawncles/etherhunt/module/clients/types"
	localcontext "github.com/shawncles/etherhunt/module/contexts"
)

type ContractHunt struct {
	ctx context.Context
	client *ctypes.EtherClient
}

func NewContractHunt() *ContractHunt {
	c := clients.CreateClient(clients.ClientType_ETH)
	ethClient, ok := c.(*ctypes.EtherClient)
	if !ok {
		return nil
	}
	return &ContractHunt{
		ctx: context.Background(),
		client: ethClient,
	}
}

func (h *ContractHunt) Watch(url string, invokeContext localcontext.InvokeContext) error {
	err := h.client.Dial(url)
	if err != nil {
		return err
	}
	h.client
	invokeContext.Contract
	invokeContext.Method
	invokeContext.Params
	return nil
}
