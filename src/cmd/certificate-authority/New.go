package CmdCertificateAuthority

import (
	"log"
	Config "ssh-ca/src/config"
	SSHKeypairED25519 "ssh-ca/src/ssh/keypair/ed25519"

	"github.com/spf13/cobra"
)

func New(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	//var __ssh_ca_name string
	var __ssh_ca_description string
	__cmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new SSH Certificate Authority",
		Example: "  ssh-ca --config-file-path ~/.ssh-ca/config.db \\\n" +
			"         ca new \\\n" +
			"         --name sample\\\n" +
			"         --description 'sample SSH CA'",
		Run: func(cmd *cobra.Command, args []string) {
			//__ssh_ca_private_key_file_path_abs, err := filepath.Abs(*ssh_ca_private_key_file_path)
			//if err != nil {
			//	fmt.Println([]byte(err.Error()))
			//	os.Exit(1)
			//}
			//__private_key_path := filepath.Join(__ssh_ca_path, __ssh_ca_name)
			//if utils.FileExists(__ssh_ca_private_key_file_path_abs) {
			//	fmt.Printf("The SSH Certificate Authority private key file path %q already exist.\n", __ssh_ca_private_key_file_path_abs)
			//	os.Exit(1)
			//} else {
			//__SSHCertificateAuthority := SSHCertificateAuthority.New(ssh_ca_name, __ssh_ca_private_key_file_path_abs)
			__config := Config.New(ssh_ca_config_file_path)
			__sshKeypairEd25519 := SSHKeypairED25519.New()
			log.Printf("%#v\n", __sshKeypairEd25519)
			__config.SetCertificateAuthority(
				__ssh_ca_description,
				string(__sshKeypairEd25519.ExportPrivateKey()),
				string(__sshKeypairEd25519.ExportPublicKey()),
			)
			//__SSHCertificateAuthority.Show()
			//__SSHCertificateAuthority.SaveOnDisk()
			//}
		},
	}

	//__New.PersistentFlags().StringVarP(ssh_ca_private_key_file_path, "path", "p", "/etc/ssh-ca/ca", "The SSH Certificate Authority private key file path")
	/*
		__cmd.PersistentFlags().StringVarP(
			&__ssh_ca_name,
			"name",
			"",
			"",
			"The SSH Certificate Authority name",
		)
		__cmd.MarkPersistentFlagRequired("name")
	*/
	__cmd.PersistentFlags().StringVarP(
		&__ssh_ca_description,
		"description",
		"",
		"",
		"The SSH Certificate Authority description",
	)
	return __cmd
}
