package sshca

import (
	"crypto/ed25519"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func addCertToAgent(key *ed25519.PrivateKey, cert *ssh.Certificate, validHours int) error {
	fmt.Printf("\n")
	addkey := &agent.AddedKey{
		PrivateKey:   key,
		Certificate:  cert,
		LifetimeSecs: uint32(validHours * 3600),
	}
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		userSshAgent := agent.NewClient(sshAgent)
		fmt.Printf("adding certificate to ssh-agent\n")
		if err := (userSshAgent).Add(*addkey); err == nil {
			fmt.Errorf("added certificate to ssh-agent, please run ssh-add -L to verify\n")
			return nil
		} else {
			fmt.Errorf("Failed to add certificate to ssh-agent: %s\n", err)
			return err
		}
	} else {
		return err
	}
}
