package ether_contract_hunt

import (
	"context"
	"github.com/etherhunt/module/client"
	ctypes "github.com/etherhunt/module/client/types"
	mcontext "github.com/etherhunt/module/context"
)

type ContractHunt struct {
	ctx context.Context
	client *ctypes.EtherClient
}

func NewContractHunt() *ContractHunt {
	c := client.CreateClient(client.ClientType_ETH)
	ethClient, ok := c.(*ctypes.EtherClient)
	if !ok {
		return nil
	}
	return &ContractHunt{
		ctx: context.Background(),
		client: ethClient,
	}
}

func (h *ContractHunt) Watch(url string, invokeContext mcontext.InvokeContext) error {
	err := h.client.Dial(url)
	if err != nil {
		return err
	}
	h.client.
	invokeContext.Contract
	invokeContext.Method
	invokeContext.Params
	return nil
}
