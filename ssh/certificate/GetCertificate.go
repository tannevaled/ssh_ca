package SSHCertificate

import "golang.org/x/crypto/ssh"

func (self SSHCertificateStruct) GetCertificate() *ssh.Certificate {
	return self.certificate
}
