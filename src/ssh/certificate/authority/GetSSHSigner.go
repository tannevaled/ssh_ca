package SSHCertificateAuthority

import "golang.org/x/crypto/ssh"

func (self Struct) GetSSHSigner() ssh.Signer {
	return self.sshSigner
}
