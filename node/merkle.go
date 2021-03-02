package node

import (
	"github.com/tokentransfer/interfaces/block"
	"github.com/tokentransfer/interfaces/core"
)

type MerkleTree interface {
	GetRoot() []byte
	Commit() error
	Cancel() error
	GetData(key []byte) ([]byte, error)
	PutData(key, value []byte) error
}

type MerkleService interface {
	core.Service

	PutBlock(block.Block) error
	GetBlockByIndex(uint64) (block.Block, error)
	GetBlockByHash(core.Hash) (block.Block, error)

	PutTransaction(block.TransactionWithData) error
	GetTransaction(core.Hash) (block.TransactionWithData, error)
	GetTransactionByHash(core.Hash) (block.TransactionWithData, error)
	GetTransactionByIndex(core.Address, uint64) (block.TransactionWithData, error)
	GetTransactionRoot() core.Hash

	PutState(block.State) error
	GetState(core.Hash) (block.State, error)
	GetStateByKey(string) (block.State, error)
	GetStateByIndex(string, uint64) (block.State, error)
	GetStateRoot() core.Hash

	Commit() error
	Cancel() error

	VerifyTransaction(tx block.Transaction) (bool, error)
	ProcessTransaction(tx block.Transaction) (block.TransactionWithData, error)
}
