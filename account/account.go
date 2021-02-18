package account

import (
	"github.com/tokentransfer/interfaces/core"
)

type PublicKey interface {
	core.Binariable

	Verify(hash core.Hash, msg []byte, signature core.Signature) (bool, error)

	GenerateAddress() (core.Address, error)
}

type PrivateKey interface {
	core.Textable
	core.Binariable

	GeneratePublic() (PublicKey, error)
	GetSecret() (string, error)
}

type Key interface {
	core.Textable
	core.Binariable

	Sign(hash core.Hash, msg []byte) (core.Signature, error)

	GetPrivate() (PrivateKey, error)
	GetPublic() (PublicKey, error)
	GetAddress() (core.Address, error)
}
