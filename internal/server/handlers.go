package server

import (
	"artubric/transmission-rest/internal/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	trpc "github.com/hekmon/transmissionrpc/v2"
)

func (s *Server) handleTorrentsV1(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusNotImplemented)
	case http.MethodDelete:
		w.WriteHeader(http.StatusNotImplemented)
	case http.MethodPut:
		w.WriteHeader(http.StatusNotImplemented)
	case http.MethodPost:
		// TODO: seperate logic
		// unmarshal JSON => model
		var addTorrentRequest model.AddTorrentRequest
		var addTorrentPayload trpc.TorrentAddPayload

		err = json.Unmarshal(body, &addTorrentRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		addTorrentPayload.Filename = &addTorrentRequest.Filename
		addTorrentPayload.DownloadDir = &addTorrentRequest.DownloadDir

		// pass model to appropriate method
		torrent, err := s.transServ.CreateNewTorrent(r.Context(), addTorrentPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorJSON, _ := json.Marshal(err)
			w.Write(errorJSON)
			return
		}

		// marshal model => JSON
		torrentJSON, err := json.Marshal(torrent)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorJSON, _ := json.Marshal(err)
			w.Write(errorJSON)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(torrentJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
