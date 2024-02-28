package erc20_transfers

import (
	"context"
	"fmt"
	"time"

	"github.com/hoaleee/go-ethereum/accounts/abi/bind"
	"github.com/hoaleee/go-ethereum/core/types"
	"github.com/hoaleee/go-ethereum/ethclient"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/ERC20"
	"github.com/hoaleee/zkevm-node/test/operations"
)

const (
	txTimeout = 60 * time.Second
)

func DeployERC20Contract(client *ethclient.Client, ctx context.Context, auth *bind.TransactOpts) (*ERC20.ERC20, error) {
	var (
		tx  *types.Transaction
		err error
	)
	fmt.Println("Sending TX to deploy ERC20 SC")
	_, tx, erc20SC, err := ERC20.DeployERC20(auth, client, "Test Coin", "TCO")
	if err != nil {
		panic(err)
	}
	err = operations.WaitTxToBeMined(ctx, client, tx, txTimeout)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending TX to do a ERC20 mint")
	tx, err = erc20SC.Mint(auth, mintAmountBig)
	if err != nil {
		panic(err)
	}
	err = operations.WaitTxToBeMined(ctx, client, tx, txTimeout)
	if err != nil {
		panic(err)
	}
	return erc20SC, err
}
