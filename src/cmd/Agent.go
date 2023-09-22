package Cmd

import (
	CmdAgent "ssh-ca/src/cmd/agent"

	"github.com/spf13/cobra"
)

/*
 *
 * SSH Agent (sub-)commands
 *
 */
func Agent(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	__cmd := &cobra.Command{
		Use:   "agent",
		Short: "SSH agent sub-commands",
	}

	__cmd.AddCommand(
		CmdAgent.AddCertificate(ssh_ca_config_file_path),
		CmdAgent.ListCertificate(),
	)
	return __cmd
}
