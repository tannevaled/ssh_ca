package SSHCertificateAuthority

import (
	"crypto/rand"
	"fmt"
	mathRand "math/rand"
	"os"
	SSHCertificate "ssh-ca/ssh/certificate"
	"time"

	"golang.org/x/crypto/ssh"
)

func (self Struct) SignUserCertificate(
	certificate SSHCertificate.SSHCertificate,
	userPrincipal string,
	validHours int,
) error {
	__critical_options := map[string]string{}
	__extensions := map[string]string{
		"permit-X11-forwarding":   "",
		"permit-agent-forwarding": "",
		"permit-port-forwarding":  "",
		"permit-pty":              "",
		"permit-user-rc":          "",
	}
	__certificate := &ssh.Certificate{
		//ValidPrincipals: []string{userPrincipal, "ubuntu"},
		ValidPrincipals: []string{"ubuntu"},
		ValidAfter:      uint64(time.Now().Unix()),
		ValidBefore:     uint64(time.Now().Unix() + int64(validHours)*3600),
		Key:             certificate.GetSSHKeypairED25519().GetSSHKeypairED25519PublicKey().GetSSHPublicKey(),
		Serial:          mathRand.Uint64(),
		KeyId:           fmt.Sprintf("ssh cert for %s", userPrincipal),
		CertType:        ssh.UserCert,
		SignatureKey:    self.sshSigner.PublicKey(),
		Permissions: ssh.Permissions{
			CriticalOptions: __critical_options,
			Extensions:      __extensions,
		},
	}
	err := __certificate.SignCert(rand.Reader, self.sshSigner)
	if err != nil {
		fmt.Println([]byte(err.Error()))
		os.Exit(1)
	}
	certificate.SetCertificate(__certificate)
	return err
}
