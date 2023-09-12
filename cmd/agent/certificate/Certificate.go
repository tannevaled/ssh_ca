package CmdAgentCertificate

import (
	"github.com/spf13/cobra"
)

func AgentCertificate(
	ssh_ca_config_file_path *string,

) *cobra.Command {
	__cmd := &cobra.Command{
		Use:     "certificate",
		Aliases: []string{"cert"},
		Short:   "certificate sub-commands",
	}

	__cmd.AddCommand(AgentCertificateCreate(ssh_ca_config_file_path))
	__cmd.AddCommand(AgentCertificateList())

	return __cmd
}
