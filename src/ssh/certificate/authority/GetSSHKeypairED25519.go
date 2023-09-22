package SSHCertificateAuthority

import SSHKeypairED25519 "ssh-ca/src/ssh/keypair/ed25519"

func (self Struct) GetSSHKeypairED25519() SSHKeypairED25519.Interface {
	return self.SSHKeypairED25519
}
