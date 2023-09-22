package Config

import (
	"crypto/ed25519"
	"fmt"
	"log"
	"os"
	SSHCertificateAuthority "ssh-ca/src/ssh/certificate/authority"
	DateTime "ssh-ca/src/utils/datetime"

	"golang.org/x/crypto/ssh"
)

/**
 *
 * Given the SSH Certificate Authority Name,
 * create a new SSHCertificateAuthority Interface
 *
 */
func (self Struct) GetCertificateAuthority(
	certificate_authority_name *string,
) SSHCertificateAuthority.Interface {
	var certificateAuthority CertificateAuthority
	var ed25519_private_key ed25519.PrivateKey
	row := self.db.QueryRow(
		`SELECT * from certificate_authorities WHERE name=?`,
		certificate_authority_name,
	)
	if err := row.Scan(
		&certificateAuthority.ID,
		&certificateAuthority.Enabled,
		&certificateAuthority.Uuid,
		&certificateAuthority.Description,
		&certificateAuthority.PrivateKey,
		&certificateAuthority.PublicKey,
		&certificateAuthority.CreationDate,
	); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	raw_private_key, err := ssh.ParseRawPrivateKey([]byte(certificateAuthority.PrivateKey))
	if err == nil {
		//fmt.Printf(" __raw_private_key: %#v\n", __raw_private_key)
	} else {
		log.Fatal("Decrypt failed: %v", err)
	}

	p_ed25519_private_key, ok := (raw_private_key).(*ed25519.PrivateKey)
	if ok {
		ed25519_private_key = *p_ed25519_private_key
	}
	//fmt.Printf("__ed25519_private_key: %#v\n", __ed25519_private_key)
	//fmt.Printf("%#v\n", __raw_private_key)
	//if !ok {
	//	log.Fatal("Type assertion failed\n")
	//}
	//fmt.Printf("__ed25519_private_key.Public(): %#v\n", __ed25519_private_key.Public().(ed25519.PublicKey))
	ed25519_public_key := ed25519_private_key.Public().(ed25519.PublicKey)

	return SSHCertificateAuthority.New(
		&certificateAuthority.Uuid,
		&certificateAuthority.Description,
		DateTime.FromString(certificateAuthority.CreationDate),
		ed25519_public_key,
		ed25519_private_key,
	)
}
