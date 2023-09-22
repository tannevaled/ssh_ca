package SSHKeypairED25519

import (
	SSHKeypairEd25519PrivateKey "ssh-ca/src/ssh/keypair/ed25519/private"
	SSHKeypairEd25519PublicKey "ssh-ca/src/ssh/keypair/ed25519/public"
)

type Interface interface {
	GetSSHKeypairED25519PrivateKey() SSHKeypairEd25519PrivateKey.SSHKeypairED25519PrivateKey //ed25519.PrivateKey
	GetSSHKeypairED25519PublicKey() SSHKeypairEd25519PublicKey.SSHKeypairED25519PublicKey    //ssh.PublicKey
	SetComment(string)
	GetComment() string
	AUTH_MAGIC() []byte
	MarshalED25519PrivateKey() []byte
	//exportPrivateKey(map[string]string) []byte
	ExportPrivateKey() []byte
	ExportPublicKey() []byte
	MarshaledPubkeyPrefix() []byte
	MarshaledPubkey() []byte
	//saveOnDisk()
}

type SSHKeypairED25519Struct struct {
	SSHKeypairED25519PrivateKey SSHKeypairEd25519PrivateKey.SSHKeypairED25519PrivateKey //ed25519.PrivateKey
	SSHKeypairED25519PublicKey  SSHKeypairEd25519PublicKey.SSHKeypairED25519PublicKey   //ssh.PublicKey
	comment                     string
}

//func LoadSSHKeypairFromFiles(privateFilePath string) SSHKeypair {
//	return nil
//}

/*
func (self SshKeypairEd25519Struct) exportPrivateKey(headers map[string]string) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:    "OPENSSH PRIVATE KEY",
		Headers: headers,
		Bytes:   self.marshalEd25519PrivateKey(),
	})
}
*/
