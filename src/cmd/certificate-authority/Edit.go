package CmdCertificateAuthority

import (
	Config "ssh-ca/src/config"
	ConfigField "ssh-ca/src/config/field"
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
	var ssh_certificate_authority_state ConfigState.Enum = ConfigState.Keep
	var ssh_certificate_authority_description ConfigField.Enum = ConfigField.Keep
	var ssh_certificate_authority_description_value string

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
				&ssh_certificate_authority_state,
				&ssh_certificate_authority_description,
				&ssh_certificate_authority_description_value,
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

	//__cmdCaEdit.MarkFlagsOneRequired("name")

	cmd.Flags().Var(
		&ssh_certificate_authority_state,
		"state",
		"The SSH Certificate Authority state [ enable, disable, keep ]",
	)

	cmd.Flags().Var(
		&ssh_certificate_authority_description,
		"description",
		"The SSH Certificate Authority description operation [ set, keep ]",
	)
	cmd.Flags().StringVarP(
		&ssh_certificate_authority_description_value,
		"description-value",
		"",
		"",
		"The SSH Certificate Authority description value",
	)
	return cmd
}
