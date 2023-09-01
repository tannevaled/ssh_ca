package SSHKeypairED25519

import (
	cryptoRand "crypto/rand"
	"fmt"
	"os"
	SSHKeypairEd25519PrivateKey "ssh-ca/ssh/keypair/ed25519/private"
	SSHKeypairEd25519PublicKey "ssh-ca/ssh/keypair/ed25519/public"

	"golang.org/x/crypto/ed25519"
)

/*
 *
 * https://stackoverflow.com/questions/71850135/generate-ed25519-key-pair-compatible-with-openssh
 *
 */
func New() SSHKeypairED25519 {
	__ed25519_public_key, __ed25519_private_key, __err := ed25519.GenerateKey(cryptoRand.Reader)
	if __err != nil {
		fmt.Println(fmt.Errorf("%s\n", __err))
		os.Exit(1)
	}
	//__ssh_public_key, __err := ssh.NewPublicKey(__ed25519_public_key)
	//	if __err != nil {
	//		fmt.Println(fmt.Errorf("failed to create public key: %s\n", __err))
	//		os.Exit(1)
	//	}
	__self := &SSHKeypairED25519Struct{
		SSHKeypairED25519PrivateKey: SSHKeypairEd25519PrivateKey.New(__ed25519_private_key),
		SSHKeypairED25519PublicKey:  SSHKeypairEd25519PublicKey.New(__ed25519_public_key),
	}
	return __self
}
