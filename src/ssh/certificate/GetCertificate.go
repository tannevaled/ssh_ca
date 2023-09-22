package SSHCertificate

import "golang.org/x/crypto/ssh"

func (self Struct) GetCertificate() *ssh.Certificate {
	return self.certificate
}
