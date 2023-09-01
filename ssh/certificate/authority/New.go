package SSHCertificateAuthority

import (
	"fmt"
	"os"
	SSHKeypairED25519 "ssh-ca/ssh/keypair/ed25519"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func New(
	name *string,
	description *string,
) Interface {
	__sshKeypairEd25519 := SSHKeypairED25519.New()

	//	if *caKeyPassphrasePath == "" {
	//__signer, err = ssh.ParsePrivateKey(data)
	__sshSigner, err := ssh.NewSignerFromKey(
		__sshKeypairEd25519.GetSSHKeypairED25519PrivateKey().GetED25519PrivateKey(),
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
