package CmdCa

import (
	"github.com/spf13/cobra"
	// Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	//_ "github.com/glebarez/sqlite"
)

/*
 *
 * SSH CA (sub-)commands
 *
 */
func Ca(
	ssh_ca_config_file_path *string,
	ssh_ca_name *string,
	ssh_ca_private_key_file_path *string,
) *cobra.Command {
	__Ca := &cobra.Command{
		Use:     "ca",
		Aliases: []string{"certificate-authority"},
		Short:   "certificate-authority sub-commands",
	}

	__Ca.AddCommand(New(ssh_ca_config_file_path))
	__Ca.AddCommand(Show(ssh_ca_config_file_path))
	__Ca.AddCommand(List(ssh_ca_config_file_path))
	__Ca.AddCommand(Edit(ssh_ca_config_file_path))
	__Ca.AddCommand(Toggle(ssh_ca_config_file_path))

	__Ca.PersistentFlags().StringVarP(
		ssh_ca_config_file_path,
		"config-file-path",
		"c",
		"",
		"The SSH Certificate Authority config file path",
	)
	__Ca.MarkPersistentFlagRequired("config-file-path")

	return __Ca
}
