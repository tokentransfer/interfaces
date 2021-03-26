package block

import (
	"github.com/tokentransfer/interfaces/core"
	"github.com/tokentransfer/interfaces/crypto"
)

type Transaction interface {
	core.Indexable
	crypto.Signable

	GetTransactionType() TransactionType
}

type TransactionWithData interface {
	crypto.Hashable

	GetTransaction() Transaction
	GetReceipt() Receipt
}
