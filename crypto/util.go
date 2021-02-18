package crypto

import (
	"bytes"

	"github.com/tokentransfer/interfaces/core"
)

func ZeroHash(cs CryptoService) core.Hash {
	l := cs.GetSize()
	return core.Hash(bytes.Repeat([]byte{0x00}, l))
}
