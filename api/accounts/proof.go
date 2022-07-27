package accounts

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/vechain/thor/api/utils"
	"github.com/vechain/thor/chain"
	"github.com/vechain/thor/thor"
)

type merkleProof struct {
	StateRoot    string   `json:"stateRoot"`
	Account      string   `json:"account"`
	Slot         string   `json:"slot"`
	Value        string   `json:"value"`
	AccountProof []string `json:"accountProof"`
	StorageProof []string `json:"storageProof"`
}

func (a *Accounts) proofStorage(addr thor.Address, slot thor.Bytes32, summary *chain.BlockSummary) (p *merkleProof, err error) {
	state := a.stater.
		NewState(summary.Header.StateRoot(), summary.Header.Number(), summary.Conflicts, summary.SteadyNum)

	proof, storageProofBytes, value, err := state.ProveStorage(addr, slot)
	if err != nil {
		return nil, err
	}

	var accountProof []string
	for _, p := range proof {
		accountProof = append(accountProof, hexutil.Encode(p))
	}

	var storageProof []string
	for _, p := range storageProofBytes {
		storageProof = append(storageProof, hexutil.Encode(p))
	}

	return &merkleProof{
		StateRoot:    summary.Header.StateRoot().String(),
		Account:      hexutil.Encode(addr[:]),
		Slot:         hexutil.Encode(slot[:]),
		Value:        hexutil.Encode(value[:]),
		AccountProof: accountProof,
		StorageProof: storageProof,
	}, nil
}

func (a *Accounts) handleProofStorage(w http.ResponseWriter, req *http.Request) error {
	hexAddr := mux.Vars(req)["address"]
	addr, err := thor.ParseAddress(hexAddr)
	if err != nil {
		return utils.BadRequest(errors.WithMessage(err, "address"))
	}
	slot, err := thor.ParseBytes32(mux.Vars(req)["key"])
	if err != nil {
		return utils.BadRequest(errors.WithMessage(err, "storage"))
	}
	summary, err := a.handleRevision(req.URL.Query().Get("revision"))
	if err != nil {
		return err
	}
	proof, err := a.proofStorage(addr, slot, summary)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, proof)
}
