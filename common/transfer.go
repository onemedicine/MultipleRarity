package common

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func GetMultipleAdventureGasLimit(contractAbi string,fromAddress common.Address, contractAddress common.Address, summoners []*big.Int, client *ethclient.Client) (uint64, error) {

	var err error
	abiContract, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return 0, err
	}
	var iargs []interface{}
	iargs = append(iargs, contractAddress)
	iargs = append(iargs, summoners)

	input, err := abiContract.Pack("multiple_adventure", iargs...)
	if err != nil {
		return 0, err
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &contractAddress,
		Data: input,
	})
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

