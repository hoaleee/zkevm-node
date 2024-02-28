package pkg

import (
	"github.com/hoaleee/go-ethereum/common"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/ERC20"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/uniswap/v2/core/UniswapV2Factory"
	"github.com/hoaleee/zkevm-node/test/contracts/bin/uniswap/v2/periphery/UniswapV2Router02"
)

type Deployments struct {
	ACoin     *ERC20.ERC20
	ACoinAddr common.Address
	BCoin     *ERC20.ERC20
	BCoinAddr common.Address
	CCoin     *ERC20.ERC20
	CCoinAddr common.Address
	Router    *UniswapV2Router02.UniswapV2Router02
	Factory   *UniswapV2Factory.UniswapV2Factory
}
