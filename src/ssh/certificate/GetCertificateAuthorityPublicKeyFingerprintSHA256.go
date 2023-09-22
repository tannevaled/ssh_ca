package SSHCertificate

import (
	"log"
	"ssh-ca/src/utils"

	"golang.org/x/crypto/ssh"
)

func (self Struct) GetCertificateAuthorityPublicKeyFingerprintSHA256() string {
	utils.Trace()
	certificate := self.GetCertificate()
	log.Printf("certificate: %#v\n", certificate)

	utils.Trace()
	signatureKey := certificate.SignatureKey
	log.Printf("signatureKey: %#v\n", signatureKey)

	utils.Trace()
	fingerprint := ssh.FingerprintSHA256(signatureKey)
	log.Printf("fingerprint: %#v\n", fingerprint)

	return fingerprint
}
