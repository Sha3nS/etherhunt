package clients

import (
	"fmt"
	client "github.com/shawncles/etherhunt/module/clients/types"
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
