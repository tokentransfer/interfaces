package block

import (
	"github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/interfaces/crypto"
)

type Receipt interface {
	crypto.Hashable

	GetTransactionResult() TransactionResult
	GetTransactionIndex() uint32
	SetTransactionIndex(uint32)
	GetStates() []State
}

type State interface {
	core.Indexable
	crypto.Hashable

	GetStateType() StateType
	GetStateKey() string
	GetBlockIndex() uint64
	SetBlockIndex(uint64)
}
