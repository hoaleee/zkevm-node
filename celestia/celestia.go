package celestia

import (
	"encoding/hex"
	"github.com/rollkit/go-da"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rollkit/go-da/proxy"
)

type CelestiaConfig struct {
	Rpc         string `mapstructure:"Rpc"`
	NamespaceId string `mapstructure:"NamespaceId"`
}

type CelestiaDA struct {
	Cfg       CelestiaConfig
	Client    *proxy.Client
	Namespace da.Namespace
}

func NewCelestiaDA(cfg CelestiaConfig) (*CelestiaDA, error) {
	daClient := proxy.NewClient()
	err := daClient.Start(cfg.Rpc, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	namespace, err := hex.DecodeString(cfg.NamespaceId)

	if err != nil {
		return nil, err
	}

	return &CelestiaDA{
		Cfg:       cfg,
		Client:    daClient,
		Namespace: namespace,
	}, nil
}
