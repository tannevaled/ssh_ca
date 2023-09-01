package SSHCertificate

import (
	SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"

	"golang.org/x/crypto/ssh"
)

func New() SSHCertificate {
	__self := &SSHCertificateStruct{
		SSHKeypairED25519: SSHKeypairED25519.New(),
	}
	return __self
}

func (self SSHCertificateStruct) SetCertificate(certificate *ssh.Certificate) {
	self.certificate = certificate
}

//func (self SshCertificateStruct) publicKeyFingerprintSHA256() string {
//
//	return ssh.FingerprintSHA256(self.getKeypair().getPublicKey())
//}

func (self SSHCertificateStruct) CaPublicKeyFingerprintSHA256() string {
	return ssh.FingerprintSHA256(self.GetCertificate().SignatureKey)
}
