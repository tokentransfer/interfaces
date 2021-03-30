package account

import (
	"github.com/tokentransfer/interfaces/core"
)

type Signable interface {
	Sign(hash core.Hash, msg []byte) (core.Signature, error)
}

type Verifiable interface {
	Verify(hash core.Hash, msg []byte, signature core.Signature) (bool, error)
}

type PublicKey interface {
	core.Textable
	core.Binariable

	Verifiable

	GenerateAddress() (core.Address, error)
}

type PrivateKey interface {
	core.Textable
	core.Binariable

	Signable

	GeneratePublic() (PublicKey, error)
	GetSecret() (string, error)
}

type Key interface {
	core.Textable
	core.Binariable

	Signable
	Verifiable

	GetPrivate() (PrivateKey, error)
	GetPublic() (PublicKey, error)
	GetAddress() (core.Address, error)
}

type AccountService interface {
	GenerateFamilySeed(password string) (KeyType, Key, error)

	NewKeyFromSecret(string) (KeyType, Key, error)
	NewKeyFromBytes([]byte) (KeyType, Key, error)

	NewPublicFromHex(string) (KeyType, PublicKey, error)
	NewPublicFromBytes([]byte) (KeyType, PublicKey, error)

	NewAccountFromAddress(string) (KeyType, core.Address, error)
	NewAccountFromBytes([]byte) (KeyType, core.Address, error)
}
