package SSHCertificate

import (
	"golang.org/x/crypto/ssh"
)

/*
 *
 * pointer to struct so that we can modify it
 *
 */
func (self *Struct) SetCertificate(certificate *ssh.Certificate) {
	//utils.Trace()
	//log.Printf("certificate: %#v\n", certificate)
	self.certificate = certificate
	//utils.Trace()
	//log.Printf("self.certificate: %#v\n", self.certificate)
}
