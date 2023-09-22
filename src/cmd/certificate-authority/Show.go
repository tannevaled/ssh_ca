package CmdCertificateAuthority

import (
	Config "ssh-ca/src/config"

	"github.com/spf13/cobra"
)

func Show(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	var __ssh_ca_uuid string
	__cmd := &cobra.Command{
		Use:   "show",
		Short: "Show the specified SSH Certificate Authority",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {

			__config := Config.New(ssh_ca_config_file_path)
			__config.ShowCertificateAuthority(__ssh_ca_uuid)
			//fmt.Println(text.Bold.Sprint(("todo")))
			//__sshCaKeyFilePath, _ := cmd.Flags().GetString("sshCaKeyFilePath")
			//__sshCaKeyPassphraseFilePath, _ := cmd.Flags().GetString("sshCaKeyPassphraseFilePath")
			//__sshCA := NewSSHCA(&__ssh_ca_dir, &__sshCaKeyPassphraseFilePath)
			//__sshCA.printList()
			//		__ssh_ca_private_key_file_path_abs, err := filepath.Abs(*ssh_ca_private_key_file_path)
			//			if err != nil {
			//				fmt.Println([]byte(err.Error()))
			//				os.Exit(1)
			//			}
			//			if utils.FileExists(__ssh_ca_private_key_file_path_abs) {
			//				__sshCa := SSHCertificateAuthority.Load(__ssh_ca_private_key_file_path_abs)
			//				__sshCa.Show()
			//			} else {
			//				fmt.Printf("The SSH Certificate Authority private key file path %q does not exist.\n", __ssh_ca_private_key_file_path_abs)
			//				os.Exit(1)
			//			}
		},
	}

	__cmd.PersistentFlags().StringVarP(
		&__ssh_ca_uuid,
		"uuid",
		"",
		"",
		"The SSH Certificate Authority uuid",
	)
	__cmd.MarkPersistentFlagRequired("uuid")

	return __cmd
}
