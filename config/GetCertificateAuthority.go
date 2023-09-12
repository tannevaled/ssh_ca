package Config

import (
	"crypto/ed25519"
	"fmt"
	"log"
	"os"
	SSHCertificateAuthority "ssh-ca/ssh/certificate/authority"

	"golang.org/x/crypto/ssh"
)

func (self Struct) GetCertificateAuthority(
	certificate_authority_name *string,
) SSHCertificateAuthority.Interface {
	var certificateAuthority CertificateAuthority
	row := self.db.QueryRow(
		`SELECT * from certificate_authorities WHERE name=?`,
		certificate_authority_name,
	)
	if err := row.Scan(
		&certificateAuthority.ID,
		&certificateAuthority.Enabled,
		&certificateAuthority.Name,
		&certificateAuthority.Description,
		&certificateAuthority.PrivateKey,
		&certificateAuthority.PublicKey,
		&certificateAuthority.CreationDate,
	); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	__raw_private_key, err := ssh.ParseRawPrivateKey([]byte(certificateAuthority.PrivateKey))
	if err != nil {
		log.Fatal("Decrypt failed: %v", err)
	}
	ed25519_private_key, ok := __raw_private_key.(ed25519.PrivateKey)
	if !ok {
		log.Fatal("Type assertion failed\n")
	}
	ed25519_public_key := ed25519_private_key.Public().(ed25519.PublicKey)

	return SSHCertificateAuthority.New(
		&certificateAuthority.Name,
		&certificateAuthority.Description,
		ed25519_public_key,
		ed25519_private_key,
	)
}
