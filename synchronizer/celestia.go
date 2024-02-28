package synchronizer

import (
	"context"
	"fmt"
	"github.com/hoaleee/zkevm-node/hex"
	"github.com/hoaleee/zkevm-node/log"
	"github.com/rollkit/go-da"
)

func (s *ClientSynchronizer) getBlobData(blobID []byte) ([]byte, error) {

	log.Info("Requesting data from Celestia", "blob id", hex.EncodeToHex(blobID))
	blobs, err := s.celestiaDA.Client.Get(context.Background(), []da.ID{blobID}, s.celestiaDA.Namespace)
	if err != nil {
		return nil, err
	}
	if len(blobs) != 1 {
		log.Warn("celestia: unexpected length for blobs", "expected", 1, "got", len(blobs))
		if len(blobs) == 0 {
			log.Warn("celestia: skipping empty blobs")
			return nil, fmt.Errorf(`Empty blob data`)
		}
	}
	log.Info("Succesfully fetched data from Celestia", "blob id", hex.EncodeToHex(blobID))
	return blobs[0], nil
}
