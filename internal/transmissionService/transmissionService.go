package transmissionService

import (
	"artubric/transmission-rest/config"
	"context"
	"fmt"

	trpc "github.com/hekmon/transmissionrpc/v2"
)

type TransmissionService struct {
	client *trpc.Client
}

func New(conf *config.TransmissionServer) *TransmissionService {
	advancedConfig := &trpc.AdvancedConfig{
		Port: conf.Port,
	}
	client, err := trpc.New(
		conf.Host,
		conf.Username,
		conf.Password,
		advancedConfig,
	)
	if err != nil {
		panicError := fmt.Errorf("failed to connect to remote transmission server with error:\n%+v", err)
		panic(panicError)
	}

	ts := &TransmissionService{
		client: client,
	}

	return ts
}

func (ts *TransmissionService) CreateNewTorrent(ctx context.Context, payload trpc.TorrentAddPayload) (trpc.Torrent, error) {

	torrent, err := ts.client.TorrentAdd(ctx, payload)
	if err != nil {
		return torrent, err
	}

	return torrent, nil
}
