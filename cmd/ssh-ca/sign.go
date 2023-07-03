package sshca

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"
	"time"

	"golang.org/x/crypto/ssh"
)

func signUserCert(userPrincipal string, validHours int, pubKey ssh.PublicKey, caPublicKey ssh.PublicKey, signer ssh.Signer) (*ssh.Certificate, error) {
	exts := map[string]string{
		"permit-X11-forwarding":   "",
		"permit-agent-forwarding": "",
		"permit-port-forwarding":  "",
		"permit-pty":              "",
		"permit-user-rc":          "",
	}
	cert := &ssh.Certificate{
		ValidPrincipals: []string{userPrincipal},
		ValidAfter:      uint64(time.Now().Unix()),
		ValidBefore:     uint64(time.Now().Unix() + int64(validHours)*3600),
		Key:             pubKey,
		Serial:          mathrand.Uint64(),
		KeyId:           fmt.Sprintf("ssh cert for %s", userPrincipal),
		CertType:        ssh.UserCert,
		SignatureKey:    caPublicKey,
		Permissions: ssh.Permissions{
			CriticalOptions: map[string]string{},
			Extensions:      exts,
		},
	}
	err := cert.SignCert(rand.Reader, signer)
	return cert, err
}
