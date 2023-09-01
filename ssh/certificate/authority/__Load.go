package SSHCertificateAuthority

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func Load(private_key_file_path string) Interface {
	var __sshSigner ssh.Signer

	// read the CA private key file
	data, err := os.ReadFile(private_key_file_path)
	if err != nil {
		//fmt.Errorf("failed to read ca key: %s\n", err)
		fmt.Printf("failed to read ca key: %s\n", err)
		os.Exit(1)
	}
	//	if *caKeyPassphrasePath == "" {
	__sshSigner, err = ssh.ParsePrivateKey(data)
	if err != nil {
		fmt.Errorf("failed to parse ca key: %s\n", err)
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
		privateKeyFilePath: private_key_file_path,
		sshSigner:          __sshSigner,
	}
	return __self
}
