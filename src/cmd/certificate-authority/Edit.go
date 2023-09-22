package CmdCertificateAuthority

import (
	Config "ssh-ca/src/config"
	ConfigState "ssh-ca/src/config/state"

	"github.com/spf13/cobra"
)

/*
 *
 * https://stackoverflow.com/questions/50824554/permitted-flag-values-for-cobra
 *
 */
func Edit(
	ssh_ca_config_file_path *string,
) *cobra.Command {
	var ssh_certificate_authority_uuid string
	var ssh_certificate_authority_description string = "UNSET"
	var ssh_certificate_authority_state = ConfigState.Keep

	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit an SSH Certificate Authority",
		Run: func(
			cmd *cobra.Command,
			args []string,
		) {
			config := Config.New(ssh_ca_config_file_path)

			config.EditCertificateAuthority(
				&ssh_certificate_authority_uuid,
				&ssh_certificate_authority_description,
				&ssh_certificate_authority_state,
			)
			config.ListCertificateAuthorities()
		},
	}

	cmd.Flags().StringVarP(
		&ssh_certificate_authority_uuid,
		"uuid",
		"",
		"", // default value
		"The SSH Certificate Authority uuid",
	)
	cmd.MarkFlagRequired("uuid")

	cmd.Flags().StringVarP(
		&ssh_certificate_authority_description,
		"description",
		"",
		"",
		"The SSH Certificate Authority description",
	)
	//__cmdCaEdit.MarkFlagsOneRequired("name")

	cmd.Flags().Var(
		&ssh_certificate_authority_state,
		"state",
		"The SSH Certificate Authority state [ enable, disable, keep ]",
	)

	cmd.Flags().
	return cmd
}
