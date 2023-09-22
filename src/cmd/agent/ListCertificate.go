package CmdAgent

import (
	SSHAgent "ssh-ca/src/ssh/agent"

	"github.com/spf13/cobra"
)

func ListCertificate() *cobra.Command {
	__cmd := &cobra.Command{
		Use:     "list-certificate",
		Aliases: []string{"list"},
		Short:   "List the certificates available in the user's SSH agent",
		Run: func(cmd *cobra.Command, args []string) {
			__ssh_agent := SSHAgent.New()
			__ssh_agent.PrintList()
		},
	}

	return __cmd
}
