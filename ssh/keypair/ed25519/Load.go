package SSHKeypairED25519

import (
	SSHKeypairEd25519PrivateKey "ssh-ca/ssh/keypair/ed25519/private"
	SSHKeypairEd25519PublicKey "ssh-ca/ssh/keypair/ed25519/public"

	"golang.org/x/crypto/ed25519"
)

func Load(
	ed25519_public_key ed25519.PublicKey,
	ed25519_private_key ed25519.PrivateKey,
) SSHKeypairED25519 {
	__self := &SSHKeypairED25519Struct{
		SSHKeypairED25519PrivateKey: SSHKeypairEd25519PrivateKey.New(ed25519_private_key),
		SSHKeypairED25519PublicKey:  SSHKeypairEd25519PublicKey.New(ed25519_public_key),
	}
	return __self
}
