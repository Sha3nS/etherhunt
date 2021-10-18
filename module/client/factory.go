package client

import (
	"fmt"
	client "github.com/etherhunt/module/client/types"
)

const (
	ClientType_ETH = "eth"
)

func CreateClient(ty string) Client {
	switch ty {
	case ClientType_ETH:
		return client.NewEtherClient()
	default:
		panic(fmt.Sprintf("type: %s not implied", ty))
	}
}
