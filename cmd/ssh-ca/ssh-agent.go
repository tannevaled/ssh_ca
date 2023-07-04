package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	strftime "github.com/itchyny/timefmt-go"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SSHAgent struct {
	socket     string
	connection net.Conn
	agent      agent.ExtendedAgent
}

type SSHAgentInterface interface {
	printList()
	add(SSHCertificateInterface, int) error
}

/*
 * cf https://pkg.go.dev/golang.org/x/crypto/ssh/agent#example-NewClient
 */
func NewSSHAgent() SSHAgentInterface {
	__socket := os.Getenv("SSH_AUTH_SOCK")
	__connection, err := net.Dial("unix", __socket)
	if err != nil {
		log.Fatalf("Failed to open SSH_AUTH_SOCK: %v", err)
	}
	__agent := agent.NewClient(__connection)
	__self := &SSHAgent{
		socket:     __socket,
		connection: __connection,
		agent:      __agent,
	}
	return __self
}

/**
 *
 *
 *
 */
func (self *SSHAgent) add(certificate SSHCertificateInterface, validHours int) error {
	//fmt.Printf("\n")
	__validBefore := time.Unix(int64(certificate.getCertificate().ValidBefore), 0)
	__comment := fmt.Sprintf(
		"SSH-CA:%s:VALID_BEFORE:%s",
		certificate.caPublicKeyFingerprintSHA256(),
		strftime.Format(__validBefore, "%Y:%m:%d:%H:%M:%S"),
	)
	addkey := &agent.AddedKey{
		PrivateKey:   certificate.getPrivateKey(),
		Certificate:  certificate.getCertificate(),
		LifetimeSecs: uint32(validHours * 3600),
		Comment:      __comment,
	}

	fmt.Printf("adding certificate to ssh-agent\n")
	if err := (self.agent).Add(*addkey); err == nil {
		fmt.Errorf("added certificate to ssh-agent, please run ssh-add -L to verify\n")
		return nil
	} else {
		fmt.Errorf("Failed to add certificate to ssh-agent: %s\n", err)
		return err
	}
}

/**
 *
 * Lists fingerprints of all identities currently represented by the agent.
 *
 */
func (self *SSHAgent) printList() {
	key_array, err := self.agent.List()
	if err != nil {
		fmt.Println([]byte(err.Error()))
		os.Exit(1)
	}
	for _, ssh_key := range key_array {
		/**
		 *
		 * x/crypto: document how to unmarshal ssh certs
		 *
		 * https://github.com/golang/go/issues/22046
		 *
		 * https://cs.opensource.google/go/x/crypto/+/refs/tags/v0.10.0:ssh/keys.go;l=271
		 *
		 */
		public_key, err := ssh.ParsePublicKey(ssh_key.Blob)
		if err != nil {
			fmt.Println([]byte(err.Error()))
			os.Exit(1)
		}
		ssh_certificate := public_key.(*ssh.Certificate)
		str_valid_after := fmt.Sprintf(
			"<ValidAfter>%s</ValidAfter>",
			strftime.Format(time.Unix(int64(ssh_certificate.ValidAfter), 0), "%Y/%m/%d %H:%M:%S"),
		)
		str_valid_before := fmt.Sprintf(
			"<ValidBefore>%s</ValidBefore>",
			strftime.Format(time.Unix(int64(ssh_certificate.ValidBefore), 0), "%Y/%m/%d %H:%M:%S"),
		)
		str_principals := fmt.Sprintf("<Principals>%#v<Principals>", ssh_certificate.ValidPrincipals)
		fmt.Printf(
			"<SSHCertificate>\n\t%s\n\t%s\n\t%s\n</SSHCertificate>\n",
			str_valid_after,
			str_valid_before,
			str_principals,
		)
	}
}
