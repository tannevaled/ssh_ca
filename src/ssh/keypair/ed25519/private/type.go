package SSHKeypairED25519PrivateKey

import (
	"crypto/ed25519"
)

type SSHKeypairED25519PrivateKeyStruct struct {
	ED25519PrivateKey ed25519.PrivateKey
}

type SSHKeypairED25519PrivateKey interface {
	GetED25519PrivateKey() ed25519.PrivateKey
	GetED25519PublicKey() ed25519.PublicKey
	Export() []byte
	//marshal() []byte
}
