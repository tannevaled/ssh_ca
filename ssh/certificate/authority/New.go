package SSHCertificateAuthority

import (
	"crypto/ed25519"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"

	SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"
)

func New(
	name *string,
	description *string,
	ed25519_public_key ed25519.PublicKey,
	ed25519_private_key ed25519.PrivateKey,
) Interface {
	__sshKeypairEd25519 := SSHKeypairED25519.Load(
		ed25519_public_key,
		ed25519_private_key,
	)

	//	if *caKeyPassphrasePath == "" {
	//__signer, err = ssh.ParsePrivateKey(data)
	__sshSigner, err := ssh.NewSignerFromKey(
		ed25519_private_key,
	//	__sshKeypairEd25519.GetSSHKeypairED25519PrivateKey().GetED25519PrivateKey(),
	)
	if err != nil {
		//fmt.Errorf("failed to parse ca key: %s\n", err)
		fmt.Println(err)
		os.Exit(1)
	}
	//	} else {
	//		passphrasedata, err := ioutil.ReadFile(*caKeyPassphrasePath)
	//		if err != nil {
	//			fmt.Errorf("failed to read ca key passphrase file: %s\n", err)
	//			os.Exit(1)
	//		}
	//		__signer, err = ssh.ParsePrivateKeyWithPassphrase(data, passphrasedata)
	//		if err != nil {
	//			fmt.Errorf("failed to parse encrypted ca key: %s\n", err)
	//			os.Exit(1)
	//		}
	//	}

	__self := &Struct{
		name:              strings.Clone(*name),
		description:       strings.Clone(*description),
		creationDate:      time.Now(),
		SSHKeypairED25519: __sshKeypairEd25519,
		sshSigner:         __sshSigner,
	}
	return __self
}
