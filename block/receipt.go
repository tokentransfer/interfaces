package block

import (
	"github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/interfaces/crypto"
)

type Receipt interface {
	crypto.Hashable

	GetTransactionResult() TransactionResult
	GetStates() []State

	GetTransactionIndex() uint32
	SetTransactionIndex(uint32)
	GetBlockIndex() uint64
	SetBlockIndex(uint64)
}

type State interface {
	core.Indexable
	crypto.Hashable

	GetStateType() StateType
	GetAccount() core.Address
	GetStateKey() string

	GetBlockIndex() uint64
	SetBlockIndex(uint64)
}
