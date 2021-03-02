package block

import (
	"github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/interfaces/crypto"
)

type Block interface {
	core.Indexable
	crypto.Hashable

	GetParentHash() core.Hash
	GetTransactionHash() core.Hash
	GetStateHash() core.Hash
	GetTransactions() []TransactionWithData
	GetStates() []State
}
