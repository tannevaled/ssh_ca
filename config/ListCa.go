package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (self Struct) ListCa() {
	//var __sshSigner ssh.Signer
	var certificateAuthorities []CertificateAuthority
	rows, err := self.db.Query("SELECT * from ca ORDER BY creation_date")
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var certificateAuthority CertificateAuthority
		if err := rows.Scan(
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
		certificateAuthorities = append(certificateAuthorities, certificateAuthority)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"CREATION DATE", "ENABLED", "NAME", "PRIVATE KEY", "PUBLIC KEY", "DESCRIPTION"})
	//for _, row := range config.(map[string]interface{}) {
	for _, certificateAuthority := range certificateAuthorities {
		log.Printf("%#v", certificateAuthority)
		//__sshSigner, err = ssh.ParsePrivateKey([]byte(certificateAuthority.PrivateKey))
		if err != nil {
			fmt.Errorf("failed to parse ca key: %s\n", err)
			os.Exit(1)
		}
		t.AppendRow(table.Row{
			certificateAuthority.CreationDate,
			certificateAuthority.Enabled,
			certificateAuthority.Name,
			certificateAuthority.PrivateKey,
			certificateAuthority.PublicKey,
			certificateAuthority.Description,
			//fmt.Sprintf("%q", self.getPrivateKeyFilePath()),
		})
		/*
			t.AppendRow(table.Row{
				"CA_CREATION_DATE",
				fmt.Sprintf("%q", self.getCreationDate()),
			})
			t.AppendRow(table.Row{
				"CA_NAME",
				fmt.Sprintf("%q", self.getName()),
			})
			t.AppendRow(table.Row{
				"CA_PUBLIC_KEY",
				string(ssh.MarshalAuthorizedKey(self.sshSigner.PublicKey())),
			})
			//fmt.Println("%#v")
		*/
	}
	t.Render()
}
