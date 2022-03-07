package eth_contract_watcher

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shawncles/etherhunt/module/clients"
	ctypes "github.com/shawncles/etherhunt/module/clients/types"
	localcontext "github.com/shawncles/etherhunt/module/contexts"
	"io/ioutil"
	"time"
)

type ContractHunt struct {
	ctx context.Context
	client *ctypes.EtherClient
}

func NewContractHunt(ctx context.Context) *ContractHunt {
	c := clients.CreateClient(clients.ClientType_ETH)
	ethClient, ok := c.(*ctypes.EtherClient)
	if !ok {
		return nil
	}
	return &ContractHunt{
		ctx: ctx,
		client: ethClient,
	}
}

func (h *ContractHunt) Watch(url string, privkeyPath string, height uint64, invokeContext localcontext.InvokeContext) error {
	err := h.client.Dial(url)
	if err != nil {
		return err
	}
	// parse private key
	privkeyHex, err := ioutil.ReadFile(privkeyPath)
	if err != nil {
		return err
	}
	privkey, err := crypto.HexToECDSA(string(privkeyHex))
	if err != nil {
		return err
	}
	address := crypto.PubkeyToAddress(privkey.PublicKey)
	// get context
	txHex, err := h.client.BuildTx(address, invokeContext.Contract, []byte{}, 1)
	if err != nil {
		return nil
	}
	signedHex, err := h.client.SignTx([]byte(txHex), privkey)
	if err != nil {
		return err
	}

	go func() {
		for {
			currentHeight, err := h.client.BestBlockNumber()
			if err != nil {
				currentHeight = 0
				fmt.Println(err)
			}
			if currentHeight >= height {
				txHash, err := h.client.BroadcastTx(signedHex)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("TxHash: ", txHash)
			}
			select {
			case <- time.After(5 * time.Second):
				continue
			case <- h.ctx.Done():
				fmt.Println(h.ctx.Err())
			}
		}
	}()

	h.client.BroadcastTx()

	invokeContext.Contract
	invokeContext.Method
	invokeContext.Params
	return nil
}

func (h *ContractHunt) Hold() {

}
