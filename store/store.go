package store

import (
	"github.com/tokentransfer/interfaces/core"
)

type KvService interface {
	core.Service

	GetData(key []byte) ([]byte, error)
	GetDatas(keys [][]byte) ([][]byte, error)

	PutData(key []byte, value []byte) error
	PutDatas(keys [][]byte, values [][]byte) error

	HasData(key []byte) bool
	RemoveData(key []byte) error

	Flush() error

	ListData(func(key []byte, value []byte) error) error
}
