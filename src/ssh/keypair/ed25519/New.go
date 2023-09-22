package SSHKeypairED25519

import (
	cryptoRand "crypto/rand"
	"fmt"
	"os"

	"golang.org/x/crypto/ed25519"
)

/*
 *
 * https://stackoverflow.com/questions/71850135/generate-ed25519-key-pair-compatible-with-openssh
 *
 */
func New() Interface {
	ed25519_public_key, ed25519_private_key, err := ed25519.GenerateKey(cryptoRand.Reader)
	if err != nil {
		fmt.Println(fmt.Errorf("%s\n", err))
		os.Exit(1)
	}
	self := Load(
		ed25519_public_key,
		ed25519_private_key,
	)
	return self
}
