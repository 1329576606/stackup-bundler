package modules

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stackup-wallet/stackup-bundler/pkg/userop"
)

// BatchHandlerCtx is the object passed to BatchHandler functions.
type BatchHandlerCtx struct {
	Batch          []*userop.UserOperation
	PendingRemoval []*userop.UserOperation
	EntryPoint     common.Address
	ChainID        *big.Int
	Data           map[string]any
}

// MarkOpIndexForRemoval will remove the op by index from the batch and add it to the pending removal set.
// This should be used for ops that are not to be included on-chain and dropped from the mempool.
func (c *BatchHandlerCtx) MarkOpIndexForRemoval(index int) {
	batch := []*userop.UserOperation{}
	var op *userop.UserOperation
	for i, curr := range c.Batch {
		if i == index {
			op = curr
		} else {
			batch = append(batch, curr)
		}
	}
	if op == nil {
		return
	}

	c.Batch = batch
	c.PendingRemoval = append(c.PendingRemoval, op)
}

// UserOpHandlerCtx is the object passed to UserOpHandler functions.
type UserOpHandlerCtx struct {
	UserOp     *userop.UserOperation
	EntryPoint common.Address
	ChainID    *big.Int
}

// BatchHandlerFunc is an interface to support modular processing of UserOperation batches.
type BatchHandlerFunc func(ctx *BatchHandlerCtx) error

// OpHandlerFunc is an interface to support modular processing of single UserOperations.
type UserOpHandlerFunc func(ctx *UserOpHandlerCtx) error
