package CmdAgentCertificate

import (
	SSHAgent "ssh-ca/ssh/agent"

	"github.com/spf13/cobra"
)

func AgentCertificateList() *cobra.Command {
	__cmdAgentCertificateList := &cobra.Command{
		Use:   "list",
		Short: "list the certificates available in the user agent",
		Run: func(cmd *cobra.Command, args []string) {
			__ssh_agent := SSHAgent.New()
			__ssh_agent.PrintList()
		},
	}

	return __cmdAgentCertificateList
}
