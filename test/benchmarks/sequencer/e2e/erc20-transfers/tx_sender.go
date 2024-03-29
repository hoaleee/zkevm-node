package erc20_transfers

import (
	"fmt"
	"math/big"
	"time"

	"github.com/hoaleee/go-ethereum/accounts/abi/bind"
	"github.com/hoaleee/go-ethereum/core/types"
	"github.com/hoaleee/go-ethereum/ethclient"
	"github.com/hoaleee/zkevm-node/test/benchmarks/sequencer/common/params"
	"github.com/hoaleee/zkevm-node/test/benchmarks/sequencer/common/transactions"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/ERC20"
	uniswap "github.com/hoaleee/zkevm-node/test/scripts/uniswap/pkg"
)

const (
	mintAmount = 1000000000000000000
)

var (
	sleepTime     = 1 * time.Second
	mintAmountBig = big.NewInt(mintAmount)
	countTxs      = 0
)

// TxSender sends ERC20 transfer to the sequencer
func TxSender(l2Client *ethclient.Client, gasPrice *big.Int, auth *bind.TransactOpts, erc20SC *ERC20.ERC20, uniswapDeployments *uniswap.Deployments) ([]*types.Transaction, error) {
	fmt.Printf("sending tx num: %d\n", countTxs+1)
	var actualTransferAmount *big.Int
	if countTxs%2 == 0 {
		actualTransferAmount = big.NewInt(0)
	} else {
		actualTransferAmount = big.NewInt(1)
	}
	tx, err := erc20SC.Transfer(auth, params.To, actualTransferAmount)
	if transactions.ShouldRetryError(err) {
		time.Sleep(sleepTime)
		tx, err = erc20SC.Transfer(auth, params.To, actualTransferAmount)
	}

	if err == nil {
		countTxs += 1
	}

	return []*types.Transaction{tx}, err
}
