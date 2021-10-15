package client

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EtherClient struct {
	ctx context.Context
	client *ethclient.Client
}

func NewEtherClient() *EtherClient {
	return &EtherClient{
		ctx: context.Background(),
		client: nil,
	}
}

func (c *EtherClient) Dial(url string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	c.client = client
	return nil
}


