package core

import (
	"bytes"
	"encoding"
	"encoding/hex"
	"errors"
)

type Binariable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type Textable interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type Indexable interface {
	GetIndex() uint64
}

type Config interface {
	GetId() string
	GetType() int
	GetGasAccount() Address
	GetDataDir() string
	GetBootstraps() []string
	GetAddress() string
	GetPort() int64
	GetThreadCount() uint32
	GetTimeout() int64
	GetSystemCode() string
	GetBlockDuration() uint32
	GetBackend() string
}

// Service name and version
type Service interface {
	Init(c Config) error
	Start() error
	Close() error
}

type Hash []byte

func (h Hash) String() string {
	return EncodeToString(h)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) Equals(a Hash) bool {
	return bytes.Equal(h, a)
}

func (h Hash) IsZero() bool {
	for i := 0; i < len(h); i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h Hash) MarshalText() ([]byte, error) {
	return []byte(EncodeToString(h)), nil
}

func (h Hash) UnmarshalText(b []byte) error {
	var err error
	h, err = hex.DecodeString(string(b))
	return err
}

type PublicKey []byte

func (h PublicKey) String() string {
	return EncodeToString(h)
}

func (h PublicKey) MarshalText() ([]byte, error) {
	return []byte(EncodeToString(h)), nil
}

func (h PublicKey) UnmarshalText(b []byte) error {
	var err error
	h, err = hex.DecodeString(string(b))
	return err
}

type Signature []byte

func (h Signature) String() string {
	return EncodeToString(h)
}

func (h Signature) MarshalText() ([]byte, error) {
	return []byte(EncodeToString(h)), nil
}

func (h Signature) UnmarshalText(b []byte) error {
	var err error
	h, err = hex.DecodeString(string(b))
	return err
}

type Bytes []byte

func (h Bytes) String() string {
	return EncodeToString(h)
}

func (h Bytes) MarshalText() ([]byte, error) {
	return []byte(EncodeToString(h)), nil
}

func (h Bytes) UnmarshalText(b []byte) error {
	var err error
	h, err = hex.DecodeString(string(b))
	return err
}

type Symbol string

func NewSymbol(c string) (*Symbol, error) {
	if len(c) > 0 {
		currency := Symbol(c)
		return &currency, nil
	}
	return nil, errors.New("null symbol")
}

func (c *Symbol) Equals(b *Symbol) bool {
	ca := c.String()
	ba := b.String()
	return ca == ba
}

func (c *Symbol) String() string {
	if c != nil && len(*c) > 0 {
		return string(*c)
	}
	return ""
}

type Address interface {
	Textable
	Binariable

	String() string
}
