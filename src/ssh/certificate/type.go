package SSHCertificate

import (
	SSHKeypairED25519 "ssh-ca/src/ssh/keypair/ed25519"

	"golang.org/x/crypto/ssh"
)

type Interface interface {
	GetSSHKeypairED25519() SSHKeypairED25519.Interface
	GetCertificate() *ssh.Certificate
	SetCertificate(*ssh.Certificate)
	// publicKeyFingerprintSHA256() string
	GetCertificateAuthorityPublicKeyFingerprintSHA256() string
}

type Struct struct {
	SSHKeypairED25519 SSHKeypairED25519.Interface
	certificate       *ssh.Certificate
}
