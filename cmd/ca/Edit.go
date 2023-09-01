package CmdCa

import (
	config "ssh-ca/config"

	"github.com/spf13/cobra"
)

func Edit(ssh_ca_config_file_path *string) *cobra.Command {
	var __ssh_ca_description string
	var __ssh_ca_name string

	__Edit := &cobra.Command{
		Use:   "edit",
		Short: "Edit an SSH Certificate Authority",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {
			__Config := config.New(ssh_ca_config_file_path)
			__Config.ListCa()
		},
	}

	__Edit.Flags().StringVarP(
		&__ssh_ca_name,
		"name",
		"n",
		"", // default value
		"The SSH Certificate Authority name",
	)
	__Edit.MarkFlagRequired("name")

	__Edit.Flags().StringVarP(
		&__ssh_ca_description,
		"desc",
		"d",
		"",
		"The SSH Certificate Authority description",
	)
	//__cmdCaEdit.MarkFlagsOneRequired("name")

	return __Edit
}
