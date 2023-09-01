package SSHKeypairED25519PrivateKey

import (
	"crypto/ed25519"
)

func (self SSHKeypairED25519PrivateKeyStruct) GetED25519PublicKey() ed25519.PublicKey {
	return self.GetED25519PrivateKey().Public().(ed25519.PublicKey)
}
