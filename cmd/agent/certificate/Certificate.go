package CmdAgentCertificate

import (
	"github.com/spf13/cobra"
)

func AgentCertificate() *cobra.Command {
	__cmdAgentCertificate := &cobra.Command{
		Use:     "certificate",
		Aliases: []string{"cert"},
		Short:   "certificate sub-commands",
	}

	__cmdAgentCertificate.AddCommand(AgentCertificateList())

	return __cmdAgentCertificate
}
