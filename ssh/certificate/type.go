package SSHCertificate

import (
	SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"

	"golang.org/x/crypto/ssh"
)

type SSHCertificate interface {
	GetSSHKeypairED25519() SSHKeypairED25519.SSHKeypairED25519
	GetCertificate() *ssh.Certificate
	SetCertificate(*ssh.Certificate)
	// publicKeyFingerprintSHA256() string
	CaPublicKeyFingerprintSHA256() string
}

type SSHCertificateStruct struct {
	SSHKeypairED25519 SSHKeypairED25519.SSHKeypairED25519
	certificate       *ssh.Certificate
}
