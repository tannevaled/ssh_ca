package SSHCertificateAuthority

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"golang.org/x/crypto/ssh"
)

func (self Struct) Show() {
	//, _ := ssh.ParsePublicKey([]byte(self.signer.PublicKey().Marshal()))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"KEY", "VALUE"})
	t.AppendRow(table.Row{
		"CA_PRIVATE_KEY_FILE_PATH",
		fmt.Sprintf("%q", self.GetPrivateKeyFilePath()),
	})
	t.AppendRow(table.Row{
		"CA_CREATION_DATE",
		fmt.Sprintf("%q", self.GetCreationDate()),
	})
	t.AppendRow(table.Row{
		"CA_NAME",
		fmt.Sprintf("%q", self.GetName()),
	})
	t.AppendRow(table.Row{
		"CA_PUBLIC_KEY",
		string(ssh.MarshalAuthorizedKey(self.sshSigner.PublicKey())),
	})
	//fmt.Println("%#v")
	t.Render()

	//fmt.Printf("%s\n\n", self.exportPrivateKey())
	//self.saveOnDisk()
	//fmt.Printf("%s\n\n", self.exportPublicKey())
}
