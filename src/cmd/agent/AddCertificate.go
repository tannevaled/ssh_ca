package CmdAgent

import (
	"log"
	Config "ssh-ca/src/config"
	SSHAgent "ssh-ca/src/ssh/agent"
	SSHCertificate "ssh-ca/src/ssh/certificate"
	"ssh-ca/src/utils"

	"github.com/spf13/cobra"
)

func AddCertificate(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	var __ssh_ca_name string
	__cmd := &cobra.Command{
		Use:     "add-certificate",
		Aliases: []string{"add"},
		Short:   "Add a new certificate to the user's SSH agent",
		Run: func(cmd *cobra.Command, args []string) {
			__Config := Config.New(ssh_ca_config_file_path)
			__ssh_agent := SSHAgent.New()
			__ssh_certificate := SSHCertificate.New()
			__ssh_certificate_authority := __Config.GetCertificateAuthority(&__ssh_ca_name)
			log.Printf("%#v\n", __ssh_certificate_authority)
			utils.Trace()
			__ssh_certificate_signed := __ssh_certificate_authority.SignUserCertificate(
				&__ssh_certificate,
				"",
				4,
			)
			utils.Trace()
			log.Printf("__ssh_certificate: %#v\n", __ssh_certificate)
			log.Printf("__ssh_certificate_signed: %#v\n", *__ssh_certificate_signed)
			__ssh_agent.Add(&__ssh_certificate)
		},
	}

	__cmd.PersistentFlags().StringVarP(
		ssh_ca_config_file_path,
		"config-file-path", // long-flag
		"",                 // short flag
		"",                 // default value
		"The SSH Certificate Authority config file path",
	)
	__cmd.MarkPersistentFlagRequired("config-file-path")

	__cmd.PersistentFlags().StringVarP(
		&__ssh_ca_name,
		"signed-by",
		"",
		"",
		"The SSH Certificate Authority name",
	)
	__cmd.MarkPersistentFlagRequired("signed-by")

	return __cmd
}
