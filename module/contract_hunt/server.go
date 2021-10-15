package contract_hunt

import (
	"context"
	"github.com/etherhunt/module/client"
)

type ContractHunt struct {
	ctx context.Context
	client *client.EtherClient
}

func (h *ContractHunt)