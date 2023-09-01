package SSHCertificateAuthority

import "golang.org/x/crypto/ssh"

func (self Struct) PublicKeyFingerprintSHA256() string {
	return ssh.FingerprintSHA256(self.GetSSHSigner().PublicKey())
}
