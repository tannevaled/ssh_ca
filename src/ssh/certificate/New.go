package SSHCertificate

import (
	SSHKeypairED25519 "ssh-ca/src/ssh/keypair/ed25519"
)

func New() Interface {
	__self := &Struct{
		SSHKeypairED25519: SSHKeypairED25519.New(),
	}
	return __self
}

//func (self SshCertificateStruct) publicKeyFingerprintSHA256() string {
//
//	return ssh.FingerprintSHA256(self.getKeypair().getPublicKey())
//}
