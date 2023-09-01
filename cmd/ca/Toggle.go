package CmdCa

import (
	Config "ssh-ca/config"

	"github.com/spf13/cobra"
)

func Toggle(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	var __ssh_ca_name string
	__Cmd := &cobra.Command{
		Use:   "toggle",
		Short: "Enable/Disable an SSH Certificate Authority",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {
			__Config := Config.New(ssh_ca_config_file_path)
			__Config.ToggleCa(__ssh_ca_name)
		},
	}

	__Cmd.Flags().StringVarP(
		&__ssh_ca_name,
		"name", // long option
		"",     // short option
		"",     // default value
		"The SSH Certificate Authority name",
	)
	__Cmd.MarkFlagRequired("name")

	return __Cmd
}
