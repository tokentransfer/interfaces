package store

import (
	block "github.com/tokentransfer/interfaces/block"
	core "github.com/tokentransfer/interfaces/core"
)

type MerkleTree interface {
	KvService

	GetRoot() []byte
	Commit() error
	Cancel() error
	Verify(key []byte) ([]byte, error)
}

type BlockReader interface {
	GetBlockByHash(core.Hash, ...interface{}) (block.Block, error)
	GetBlockByIndex(uint64, ...interface{}) (block.Block, error)

	GetTransactionByHash(core.Hash, ...interface{}) (block.TransactionWithData, error)
	GetTransactionByIndex(core.Address, uint64, ...interface{}) (block.TransactionWithData, error)

	GetStateByHash(core.Hash, ...interface{}) (block.State, error)
	GetStateByTypeAndKey(block.StateType, string, ...interface{}) (block.State, error)
	GetStateByTypeAndAddress(block.StateType, core.Address, ...interface{}) (block.State, error)
	GetStateByAddressAndIndex(core.Address, uint64, ...interface{}) (block.State, error)
	GetStateByAddress(core.Address, ...interface{}) (block.State, error)
}

type BlockWriter interface {
	PutBlock(block.Block, ...interface{}) error
	PutTransaction(block.TransactionWithData, ...interface{}) error
	PutState(block.State, ...interface{}) error
}

type BlockService interface {
	BlockReader
	BlockWriter

	Commit(...interface{}) error
	Cancel(...interface{}) error
	Close(...interface{}) error
	Verify(key []byte, s ...interface{}) ([]byte, error)
}

type MerkleService interface {
	BlockService

	GetTransactionRoot() core.Hash
	GetStateRoot() core.Hash
}
