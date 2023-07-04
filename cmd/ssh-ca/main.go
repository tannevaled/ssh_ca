package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/fatih/color"
)

var (
	validHours          = flag.Int("hours", 24, "hours the certificate is valid for")
	caKeyPath           = flag.String("cakeypath", "/etc/ssh/ssh_ca.key", "path to the SSH CA private key")
	caKeyPassphrasePath = flag.String("cakeypasspath", "", "path to the SSH CA encrypted private key's passphrase")
)

func main() {
	flag.Parse()
	ssh_ca := NewSSHCA(caKeyPath, caKeyPassphrasePath)
	ssh_agent := NewSSHAgent()
	ssh_certificate := NewSSHCertificate()

	// get current username to be used as certificate principal
	currentUser, err := user.Current()
	if err != nil {
		fmt.Errorf("failed to get current user: %s\n", err)
		os.Exit(1)
	}
	userPrincipal := currentUser.Username
	color.Green("Current username [%s] will be used as certificate principal\n", userPrincipal)
	err = ssh_ca.signUserCertificate(ssh_certificate, userPrincipal, *validHours)
	if err != nil {
		fmt.Errorf("failed to sign SSH certificate: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Signed certificate successfully by SSH CA with the fingerprint [%s]\n", ssh_ca.FingerprintSHA256())

	if err := ssh_agent.add(ssh_certificate, *validHours); err != nil {
		fmt.Errorf("failed to add SSH certificate to ssh agent: %s\n", err)
		os.Exit(1)
	}

	//ssh_agent.list()
	fmt.Println(ssh_certificate.publicKeyFingerprintSHA256())

	ssh_agent.printList()
}
