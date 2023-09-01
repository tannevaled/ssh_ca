package SSHCertificateAuthority

import (
	"time"

	SSHCertificate "ssh-ca/ssh/certificate"
	SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"

	"golang.org/x/crypto/ssh"
)

type Interface interface {
	//	GetPrivateKeyFilePath() string
	GetName() string
	GetDescription() string
	GetSSHSigner() ssh.Signer
	GetCreationDate() time.Time
	GetSSHKeypairED25519() SSHKeypairED25519.SSHKeypairED25519
	SignUserCertificate(SSHCertificate.SSHCertificate, string, int) error
	PublicKeyFingerprintSHA256() string
	ExportPrivateKey() []byte
	ExportPublicKey() []byte
	// Show()
	// SaveOnDisk()
}

type Struct struct {
	name        string
	description string
	//	privateKeyFilePath string
	creationDate      time.Time
	SSHKeypairED25519 SSHKeypairED25519.SSHKeypairED25519
	sshSigner         ssh.Signer
}
