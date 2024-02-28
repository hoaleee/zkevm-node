package uniswap_transfers

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/hoaleee/go-ethereum/accounts/abi/bind"
	"github.com/hoaleee/go-ethereum/core/types"
	"github.com/hoaleee/go-ethereum/ethclient"
	"github.com/hoaleee/zkevm-node/test/benchmarks/sequencer/common/transactions"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/ERC20"
	uniswap "github.com/hoaleee/zkevm-node/test/scripts/uniswap/pkg"
)

var (
	gasLimit  = 21000
	sleepTime = 1 * time.Second
	countTxs  = 0
	txTimeout = 60 * time.Second
)

// TxSender sends eth transfer to the sequencer
func TxSender(l2Client *ethclient.Client, gasPrice *big.Int, auth *bind.TransactOpts, erc20SC *ERC20.ERC20, uniswapDeployments *uniswap.Deployments) ([]*types.Transaction, error) {
	msg := fmt.Sprintf("# Swap Cycle Number: %d #", countTxs+1)
	delimiter := strings.Repeat("#", len(msg))
	fmt.Println(delimiter)
	fmt.Println(msg)
	fmt.Println(delimiter)
	var err error

	txs := uniswap.SwapTokens(l2Client, auth, *uniswapDeployments)
	for transactions.ShouldRetryError(err) {
		time.Sleep(sleepTime)
		txs = uniswap.SwapTokens(l2Client, auth, *uniswapDeployments)
	}

	if err == nil {
		countTxs += 1
	}

	return txs, err
}
