package SSHKeypairED25519PublicKey

import (
	"crypto/ed25519"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type SSHKeypairED25519PublicKey interface {
	GetED25519PublicKey() ed25519.PublicKey
	GetSSHPublicKey() ssh.PublicKey
	Export() []byte
	//	marshal() []byte
	//marshaledPrefix() []byte
}

type SSHKeypairED25519PublicKeyStruct struct {
	//key ssh.PublicKey
	ed25519PublicKey ed25519.PublicKey
}

func New(ed25519_public_key ed25519.PublicKey) SSHKeypairED25519PublicKey {
	__self := &SSHKeypairED25519PublicKeyStruct{
		ed25519PublicKey: ed25519_public_key,
	}
	return __self
}

func (self SSHKeypairED25519PublicKeyStruct) GetED25519PublicKey() ed25519.PublicKey {
	return self.ed25519PublicKey
}
func (self SSHKeypairED25519PublicKeyStruct) GetSSHPublicKey() ssh.PublicKey {
	__sshPublicKey, __err := ssh.NewPublicKey(self.GetED25519PublicKey())
	if __err != nil {
		fmt.Println(fmt.Errorf("failed to create public key: %s\n", __err))
		os.Exit(1)
	}
	return __sshPublicKey
}

//func (self SshKeypairEd25519PublicKeyStruct) marshal() []byte {
//	return self
//}

func (self SSHKeypairED25519PublicKeyStruct) Export() []byte {
	return ssh.MarshalAuthorizedKey(self.GetSSHPublicKey())
}
