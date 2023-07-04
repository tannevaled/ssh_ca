package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	mathrand "math/rand"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHCA struct {
	signer ssh.Signer
}

type SSHCAInterface interface {
	signUserCertificate(SSHCertificateInterface, string, int) error
	FingerprintSHA256() string
}

func NewSSHCA(caKeyPath *string, caKeyPassphrasePath *string) SSHCAInterface {
	var __signer ssh.Signer

	// read the CA private key file
	data, err := ioutil.ReadFile(*caKeyPath)
	if err != nil {
		//fmt.Errorf("failed to read ca key: %s\n", err)
		fmt.Printf("failed to read ca key: %s\n", err)
		os.Exit(1)
	}
	if *caKeyPassphrasePath == "" {
		__signer, err = ssh.ParsePrivateKey(data)
		if err != nil {
			fmt.Errorf("failed to parse ca key: %s\n", err)
			os.Exit(1)
		}
	} else {
		passphrasedata, err := ioutil.ReadFile(*caKeyPassphrasePath)
		if err != nil {
			fmt.Errorf("failed to read ca key passphrase file: %s\n", err)
			os.Exit(1)
		}
		__signer, err = ssh.ParsePrivateKeyWithPassphrase(data, passphrasedata)
		if err != nil {
			fmt.Errorf("failed to parse encrypted ca key: %s\n", err)
			os.Exit(1)
		}
	}

	__self := &SSHCA{
		signer: __signer,
	}
	return __self
}
func (self *SSHCA) signUserCertificate(certificate SSHCertificateInterface, userPrincipal string, validHours int) error {
	__critical_options := map[string]string{}
	__extensions := map[string]string{
		"permit-X11-forwarding":   "",
		"permit-agent-forwarding": "",
		"permit-port-forwarding":  "",
		"permit-pty":              "",
		"permit-user-rc":          "",
	}
	__certificate := &ssh.Certificate{
		ValidPrincipals: []string{userPrincipal, "ubuntu"},
		ValidAfter:      uint64(time.Now().Unix()),
		ValidBefore:     uint64(time.Now().Unix() + int64(validHours)*3600),
		Key:             certificate.getPublicKey(),
		Serial:          mathrand.Uint64(),
		KeyId:           fmt.Sprintf("ssh cert for %s", userPrincipal),
		CertType:        ssh.UserCert,
		SignatureKey:    self.signer.PublicKey(),
		Permissions: ssh.Permissions{
			CriticalOptions: __critical_options,
			Extensions:      __extensions,
		},
	}
	err := __certificate.SignCert(rand.Reader, self.signer)
	if err != nil {
		fmt.Println([]byte(err.Error()))
		os.Exit(1)
	}
	certificate.setCertificate(__certificate)
	return err
}

func (self *SSHCA) FingerprintSHA256() string {
	return ssh.FingerprintSHA256(self.signer.PublicKey())
}
