package cmd

import (
	SSHAgent "ssh-ca/ssh/agent"

	"github.com/spf13/cobra"
)

/*
 *
 * SSH Agent (sub-)commands
 *
 */
func Agent() *cobra.Command {
	__cmdAgent := &cobra.Command{
		Use:   "agent",
		Short: "SSH agent sub-commands",
	}

	__cmdAgent.AddCommand(AgentCert())
	return __cmdAgent
}

func AgentCert() *cobra.Command {
	__cmdAgentCert := &cobra.Command{
		Use:   "cert",
		Short: "certificate sub-commands",
	}

	__cmdAgentCert.AddCommand(AgentCertList())

	return __cmdAgentCert
}

func AgentCertList() *cobra.Command {
	__cmdAgentCertList := &cobra.Command{
		Use:   "list",
		Short: "list the certificates available in the user agent",
		Run: func(cmd *cobra.Command, args []string) {
			__ssh_agent := SSHAgent.New()
			__ssh_agent.PrintList()
		},
	}

	return __cmdAgentCertList
}
