package sequencesender

import (
	"context"
	"math/big"
	"time"

	"github.com/hoaleee/go-ethereum/common"
	"github.com/hoaleee/go-ethereum/core/types"
	theEtherman "github.com/hoaleee/zkevm-node/etherman"
	ethmanTypes "github.com/hoaleee/zkevm-node/etherman/types"
	"github.com/hoaleee/zkevm-node/ethtxmanager"
	"github.com/hoaleee/zkevm-node/state"
	"github.com/jackc/pgx/v4"
)

// Consumer interfaces required by the package.

// etherman contains the methods required to interact with ethereum.
type etherman interface {
	BuildSequenceBatchesTxData(sender common.Address, sequences []ethmanTypes.Sequence, l2Coinbase common.Address, committeeSignaturesAndAddrs []byte) (to *common.Address, data []byte, err error)
	EstimateGasSequenceBatches(sender common.Address, sequences []ethmanTypes.Sequence, l2Coinbase common.Address, committeeSignaturesAndAddrs []byte) (*types.Transaction, error)
	GetLastBatchTimestamp() (uint64, error)
	GetLatestBlockTimestamp(ctx context.Context) (uint64, error)
	GetLatestBatchNumber() (uint64, error)
	GetCurrentDataCommittee() (*theEtherman.DataCommittee, error)
}

// stateInterface gathers the methods required to interact with the state.
type stateInterface interface {
	GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	IsBatchClosed(ctx context.Context, batchNum uint64, dbTx pgx.Tx) (bool, error)
	GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error)
	GetForcedBatch(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) (*state.ForcedBatch, error)
	GetTimeForLatestBatchVirtualization(ctx context.Context, dbTx pgx.Tx) (time.Time, error)
	GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
}

type ethTxManager interface {
	Add(ctx context.Context, owner, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64, dbTx pgx.Tx) error
	AddCommitment(ctx context.Context, batchNum uint64, blobID []byte, dbTx pgx.Tx) error
	ProcessPendingMonitoredTxs(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx)
}
