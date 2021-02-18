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
	GetReceiptHash() core.Hash
	GetTransactions() []TransactionWithData
	GetReceipts() []Receipt
}
