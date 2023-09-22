package Cmd

import (
	CmdCertificateAuthority "ssh-ca/src/cmd/certificate-authority"

	"github.com/spf13/cobra"
	// Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	//_ "github.com/glebarez/sqlite"
)

/*
 *
 * SSH CA (sub-)commands
 *
 */
func CertificateAuthority(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	__cmd := &cobra.Command{
		Use:     "certificate-authority",
		Aliases: []string{"ca"},
		Short:   "certificate-authority sub-commands",
	}

	__cmd.AddCommand(CmdCertificateAuthority.New(ssh_ca_config_file_path))
	__cmd.AddCommand(CmdCertificateAuthority.Show(ssh_ca_config_file_path))
	__cmd.AddCommand(CmdCertificateAuthority.List(ssh_ca_config_file_path))
	__cmd.AddCommand(CmdCertificateAuthority.Edit(ssh_ca_config_file_path))
	//__Ca.AddCommand(Toggle(ssh_ca_config_file_path))

	__cmd.PersistentFlags().StringVarP(
		ssh_ca_config_file_path,
		"config-file-path", //long-flag
		"",                 // short flag
		"",                 // default value
		"The SSH Certificate Authority config file path",
	)
	__cmd.MarkPersistentFlagRequired("config-file-path")

	return __cmd
}
