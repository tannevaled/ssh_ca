package SSHCertificateAuthority

import SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"

func (self Struct) GetSSHKeypairED25519() SSHKeypairED25519.SSHKeypairED25519 {
	return self.SSHKeypairED25519
}
