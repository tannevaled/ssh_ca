package CmdCa

import (
	Config "ssh-ca/config"

	"github.com/spf13/cobra"
)

func Show(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	var __ssh_ca_name string
	__Show := &cobra.Command{
		Use:   "show",
		Short: "Show the specified SSH Certificate Authority",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {

			__Config := Config.New(ssh_ca_config_file_path)
			__Config.ShowCa(__ssh_ca_name)
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
	//	__Show.PersistentFlags().StringVarP(
	//		ssh_ca_private_key_file_path,
	//		"path",
	//		"p",
	//		"/etc/ssh-ca/ca",
	//		"The SSH Certificate Authority private key file path",
	//	)

	__Show.PersistentFlags().StringVarP(
		&__ssh_ca_name,
		"name",
		"n",
		"sample-ca",
		"The SSH Certificate Authority name",
	)
	__Show.MarkPersistentFlagRequired("name")

	return __Show
}
