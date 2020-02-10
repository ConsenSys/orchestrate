package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	models "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/chain-registry/store/types"
)

type PostRequest struct {
	Name     string    `json:"name,omitempty" validate:"required"`
	URLs     []string  `json:"urls,omitempty" pg:"urls,array" validate:"min=1,unique,dive,url"`
	Listener *Listener `json:"listener,omitempty"`
}

type PostResponse struct {
	UUID string `json:"uuid"`
}

func (h Handler) postChain(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	chainRequest := &PostRequest{Listener: &Listener{}}
	err := UnmarshalBody(request.Body, chainRequest)
	if err != nil {
		writeError(rw, err.Error(), http.StatusBadRequest)
		return
	}

	chain := &models.Chain{
		Name:                      chainRequest.Name,
		TenantID:                  mux.Vars(request)["tenantID"],
		URLs:                      chainRequest.URLs,
		ListenerDepth:             chainRequest.Listener.Depth,
		ListenerBlockPosition:     chainRequest.Listener.BlockPosition,
		ListenerFromBlock:         chainRequest.Listener.BlockPosition,
		ListenerBackOffDuration:   chainRequest.Listener.BackOffDuration,
		ListenerExternalTxEnabled: chainRequest.Listener.ExternalTxEnabled,
	}

	err = h.store.RegisterChain(request.Context(), chain)
	if err != nil {
		handleChainRegistryStoreError(rw, err)
		return
	}

	_ = json.NewEncoder(rw).Encode(&PostResponse{
		UUID: chain.UUID,
	})
}
