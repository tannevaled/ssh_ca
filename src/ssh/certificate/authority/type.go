package SSHCertificateAuthority

import (
	"time"

	SSHCertificate "ssh-ca/src/ssh/certificate"
	SSHKeypairED25519 "ssh-ca/src/ssh/keypair/ed25519"

	"golang.org/x/crypto/ssh"
)

type Interface interface {
	//	GetPrivateKeyFilePath() string
	GetUUID() string
	GetDescription() string
	GetSSHSigner() ssh.Signer
	GetCreationDate() time.Time
	GetSSHKeypairED25519() SSHKeypairED25519.Interface
	SignUserCertificate(*SSHCertificate.Interface, string, int) *SSHCertificate.Interface
	PublicKeyFingerprintSHA256() string
	ExportPrivateKey() []byte
	ExportPublicKey() []byte
	// Show()
	// SaveOnDisk()
}

type Struct struct {
	uuid              string
	description       string
	creationDate      time.Time
	SSHKeypairED25519 SSHKeypairED25519.Interface
	sshSigner         ssh.Signer
}
