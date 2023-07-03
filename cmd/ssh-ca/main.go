package sshca

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

var (
	validHours          = flag.Int("hours", 4, "hours the certificate is valid for")
	caKeyPath           = flag.String("cakeypath", "/etc/ssh/ssh_ca.key", "path to the SSH CA private key")
	caKeyPasspharsePath = flag.String("cakeypasspath", "", "path to the SSH CA encrypted private key's passphrase")
)

func main() {
	flag.Parse()
	var caKey ssh.Signer
	// generate client ssh key pair
	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	sshPubKey, err := ssh.NewPublicKey(privKey.Public())
	fmt.Printf("Generated ed25519 key pairs with the public key of:\n[%s]", ssh.MarshalAuthorizedKey(sshPubKey))
	if err != nil {
		fmt.Errorf("failed to create host key pair: %s\n", err)
		os.Exit(1)
	}
	// get current username to be used as certificate principal
	currentUser, err := user.Current()
	if err != nil {
		fmt.Errorf("failed to get current user: %s\n", err)
		os.Exit(1)
	}
	userPrincipal := currentUser.Username
	fmt.Printf("Current username [%s] will be used as certificate principal\n", userPrincipal)
	// read the CA private key file
	dat, err := ioutil.ReadFile(*caKeyPath)
	if err != nil {
		fmt.Errorf("failed to read ca key: %s\n", err)
		os.Exit(1)
	}

	if *caKeyPasspharsePath == "" {
		caKey, err = ssh.ParsePrivateKey(dat)
		if err != nil {
			fmt.Errorf("failed to parse ca key: %s\n", err)
			os.Exit(1)
		}
	} else {
		passphrasedat, err := ioutil.ReadFile(*caKeyPasspharsePath)
		if err != nil {
			fmt.Errorf("failed to read ca key passphrase file: %s\n", err)
			os.Exit(1)
		}
		caKey, err = ssh.ParsePrivateKeyWithPassphrase(dat, passphrasedat)
		if err != nil {
			fmt.Errorf("failed to parse encrypted ca key: %s\n", err)
			os.Exit(1)
		}
	}

	// get the CA public key from the private key
	caPublicKey := caKey.PublicKey()
	// create a signer from the CA private key
	cert, err := signUserCert(userPrincipal, *validHours, sshPubKey, caPublicKey, caKey)
	if err != nil {
		fmt.Errorf("failed to sign certificate: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Signed certificate successfully by CA with the fingerprint [%s]", ssh.FingerprintSHA256(caPublicKey))
	if err := addCertToAgent(&privKey, cert, *validHours); err != nil {
		fmt.Errorf("failed to add ssh certificate to ssh agent: %s\n", err)
		os.Exit(1)
	}
}
