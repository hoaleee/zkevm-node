package sequencesender

import (
	"context"
	"github.com/0xPolygonHermez/zkevm-node/etherman/types"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/rollkit/go-da"
	"time"
)

func (s *SequenceSender) submitCelestiaBlob(ctx context.Context, sequences []types.Sequence) error {

	var blobs []da.Blob
	for _, seq := range sequences {
		if len(seq.BatchL2Data) > 0 {
			blobs = append(blobs, seq.BatchL2Data)
		}
	}
	var blobIDs [][]byte
	var err error
	if len(blobs) == 0 {
		log.Warn("No BatchData to post to Celestia")
	} else {
		blobIDs, err = s.celestiaDA.Client.Submit(ctx, blobs, float64(s.cfg.CelestiaDAGasPrice), s.celestiaDA.Namespace)
		if err != nil {
			log.Warn("Blob Submission error", "err", err)
			time.Sleep(12 * time.Second)
			return err
		}
	}

	var blobIndex = 0
	for _, seq := range sequences {
		var blobId []byte
		if len(seq.BatchL2Data) > 0 && blobIndex < len(blobIDs) {
			blobId = blobIDs[blobIndex]
			blobIndex++
		}
		log.Debug("xxx: db insert: blobId :", seq.BatchNumber, "data: ", seq.BatchL2Data, "blobId: ", blobId)
		err = s.ethTxManager.AddCommitment(ctx, seq.BatchNumber, blobId, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
