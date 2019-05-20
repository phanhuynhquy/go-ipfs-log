package keystore

import "github.com/libp2p/go-libp2p-crypto"

type Interface interface {
	HasKey(id string) (bool, error)

	CreateKey(id string) (*crypto.Secp256k1PrivateKey, error)

	GetKey(id string) (*crypto.Secp256k1PrivateKey, error)
}
