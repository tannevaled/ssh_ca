package CmdAgent

import (
	CmdAgentCertificate "ssh-ca/cmd/agent/certificate"

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
	__cmdAgent := &cobra.Command{
		Use:   "agent",
		Short: "SSH agent sub-commands",
	}

	__cmdAgent.AddCommand(
		CmdAgentCertificate.AgentCertificate(
			ssh_ca_config_file_path,
		)
	)
	return __cmdAgent
}
