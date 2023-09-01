package SSHKeypairED25519PrivateKey

import (
	"crypto/ed25519"
)

func New(ed25519_private_key ed25519.PrivateKey) SSHKeypairED25519PrivateKey {
	__self := &SSHKeypairED25519PrivateKeyStruct{
		ED25519PrivateKey: ed25519_private_key,
	}
	return __self
}
