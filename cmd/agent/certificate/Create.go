package CmdAgentCertificate

import (
	Config "ssh-ca/config"
	SSHAgent "ssh-ca/ssh/agent"
	SSHCertificate "ssh-ca/ssh/certificate"
	SSHCertificateAuthority "ssh-ca/ssh/certificate/authority"

	"github.com/spf13/cobra"
)

func AgentCertificateCreate(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	__cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new certificate in the user's SSH agent",
		Run: func(cmd *cobra.Command, args []string) {
			__Config := Config.New(ssh_ca_config_file_path)
			__ssh_agent := SSHAgent.New()
			__ssh_certificate := SSHCertificate.New()
			__ssh_certificate_authority := SSHCertificateAuthority.Load()
			__ssh_certificate_authority.SignUserCertificate(__ssh_certificate)
			__ssh_agent.Add(
				__ssh_certificate, // the user certificate to sign and add to the agent
				4,
			)
		},
	}

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
