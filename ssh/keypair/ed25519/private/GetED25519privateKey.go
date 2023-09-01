package SSHKeypairED25519PrivateKey

import (
	"crypto/ed25519"
)

func (self SSHKeypairED25519PrivateKeyStruct) GetED25519PrivateKey() ed25519.PrivateKey {
	return self.ED25519PrivateKey
}
