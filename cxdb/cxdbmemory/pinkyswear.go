package cxdbmemory

import (
	"sync"

	"github.com/mit-dci/lit/coinparam"
	"github.com/mit-dci/opencx/match"
)

// PinkySwearEngine is a settlement engine that does not actually do any settlement.
// It just sort of whitelists pubkeys that can / can not make orders
type PinkySwearEngine struct {
	// the pubkey whitelist
	whitelist    map[[33]byte]bool
	whitelistMtx *sync.Mutex

	// this coin
	coin *coinparam.Params
}

// CreatePinkySwearEngine creates a "pinky swear" engine for a specific coin
func CreatePinkySwearEngine(coin *coinparam.Params, whitelist [][33]byte) (engine match.SettlementEngine, err error) {
	pe := &PinkySwearEngine{
		coin:         coin,
		whitelist:    make(map[[33]byte]bool),
		whitelistMtx: new(sync.Mutex),
	}
	pe.whitelistMtx.Lock()
	for _, pubkey := range whitelist {
		pe.whitelist[pubkey] = true
	}
	pe.whitelistMtx.Unlock()
	return
}

// ApplySettlementExecution applies the settlementExecution, this assumes that the settlement execution is
// valid
func (pe *PinkySwearEngine) ApplySettlementExecution(setExec *match.SettlementExecution) (setRes *match.SettlementResult, err error) {
	// this highlights potentially how not generic the SettlementResult struct is, what to put in the NewBal field???
	setRes = &match.SettlementResult{
		// this is where we have sort of undefined on what we do
		NewBal:         uint64(0),
		SuccessfulExec: setExec,
	}
	return
}

// CheckValid returns true if the settlement execution would be valid
func (pe *PinkySwearEngine) CheckValid(setExec *match.SettlementExecution) (valid bool, err error) {
	pe.whitelistMtx.Lock()
	var ok bool
	if valid, ok = pe.whitelist[setExec.Pubkey]; !ok {
		// just being really explicit here, all this does is check if you're in the whitelist and your value
		// is true. There's no reason for it to be false.
		valid = false
		pe.whitelistMtx.Unlock()
		return
	}
	pe.whitelistMtx.Unlock()
	return
}