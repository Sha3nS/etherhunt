package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shawncles/etherhunt/module/clients"
	"math/big"
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

func (ec *EtherClient) Dial(url string) error {
	c, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	ec.client = c
	return nil
}

func (ec *EtherClient) BestBlockNumber() (uint64, error) {
	if ec.client == nil {
		return 0, clients.NilClient
	}
	height, err := ec.client.BlockNumber(ec.ctx)
	if err != nil {
		return 0, err
	}

	return height, nil
}

func (ec *EtherClient) EstimateGas(msg *ethereum.CallMsg) (uint64, error) {
	if ec.client == nil {
		return 0, clients.NilClient
	}
	gasUsed, err := ec.client.EstimateGas(ec.ctx, *msg)
	if err != nil {
		return 0, err
	}

	return gasUsed, nil
}

func (ec *EtherClient) GasPrice() (*big.Int, error) {
	if ec.client == nil {
		return nil, clients.NilClient
	}
	gasPrice, err := ec.client.SuggestGasPrice(ec.ctx)
	if err != nil {
		return nil, err
	}

	return gasPrice, nil
}


func (ec *EtherClient) GasTipCap() (*big.Int, error) {
	if ec.client == nil {
		return nil, clients.NilClient
	}
	gasTipCap, err := ec.client.SuggestGasTipCap(ec.ctx)
	if err != nil {
		return nil, err
	}
	return gasTipCap, nil
}

func (ec *EtherClient) BalanceAtCurrent(address [20]byte) (*big.Int, error) {
	if ec.client == nil {
		return nil, clients.NilClient
	}
	height, err := ec.BestBlockNumber()
	if err != nil {
		return nil ,err
	}
	gasTipCap, err := ec.client.BalanceAt(ec.ctx, address, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}

	return gasTipCap, nil
}

func (ec *EtherClient) NonceAt(address [20]byte) (uint64, error) {
	if ec.client == nil {
		return 0, clients.NilClient
	}
	height, err := ec.BestBlockNumber()
	if err != nil {
		return 0 ,err
	}
	gasTipCap, err := ec.client.NonceAt(ec.ctx, address, big.NewInt(int64(height)))
	if err != nil {
		return 0, err
	}

	return gasTipCap, nil
}

func (ec *EtherClient) BuildTx(address, to [20]byte, leverage int) (string, error) {
	if ec.client == nil {
		return "", clients.NilClient
	}
	chainID, err := ec.client.ChainID(ec.ctx)
	if err != nil {
		return "", err
	}
	nonce, err := ec.NonceAt(address)
	if err != nil {
		return "", err
	}
	tipCap, err := ec.GasTipCap()
	if err != nil {
		return "", err
	}
	gas, err := ec.GasPrice()
	if err != nil {
		return "", err
	}

	addr := common.Address(to)

	innerTx := &types.DynamicFeeTx{ // TODO
		ChainID: 	chainID,
		Nonce: 		nonce + 1,
		GasTipCap:  tipCap,
		GasFeeCap:  big.NewInt(21000),
		Gas:        gas.Uint64(),
		To:         &addr,
		Value:      big.NewInt(0),
		Data:       nil,
		AccessList: nil,
	}

	transaction := types.NewTx(innerTx)
	bz, err := transaction.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bz), nil
}

func (ec *EtherClient) SignTx(bz []byte, priv *ecdsa.PrivateKey) (string, error) {
	if ec.client == nil {
		return "", clients.NilClient
	}
	transaction := types.NewTx(&types.DynamicFeeTx{})
	err := transaction.UnmarshalBinary(bz)
	if err != nil {
		return "", err
	}
	chainID, err := ec.client.ChainID(ec.ctx)
	if err != nil {
		return "", err
	}
	signer := types.NewEIP155Signer(chainID)

	tx, err := types.SignTx(transaction, signer, priv)
	if err != nil {
		return "", err
	}
	data, err := tx.MarshalBinary()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

