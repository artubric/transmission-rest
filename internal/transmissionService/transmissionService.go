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
	fmt.Println("Connecting to transmission remote client... ")
	client, _ := trpc.New(
		conf.Host,
		conf.Username,
		conf.Password,
		advancedConfig,
	)

	// use as sanity, since error on client creation is empty
	ok, version, minVersion, err := client.RPCVersion(context.TODO())
	if ok {
		fmt.Printf("server allowed versions: %d - %d", minVersion, version)
	} else {
		panic(fmt.Errorf("failed to fetch transmission rpc version with error:\n%+v", err))
	}
	fmt.Println("Done. ")

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
