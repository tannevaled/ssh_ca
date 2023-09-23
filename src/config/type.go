package Config

import (
	"database/sql"
	ConfigField "ssh-ca/src/config/field"
	ConfigState "ssh-ca/src/config/state"
	SSHCertificateAuthority "ssh-ca/src/ssh/certificate/authority"
)

type CertificateAuthority struct {
	ID           int64
	Enabled      bool
	Uuid         string
	Description  string
	PrivateKey   string
	PublicKey    string
	CreationDate string
}

type Struct struct {
	config_file_path *string
	db               *sql.DB
}

type Interface interface {
	SetCertificateAuthority(string, string, string)
	GetCertificateAuthority(*string) SSHCertificateAuthority.Interface
	ShowCertificateAuthority(string)
	ListCertificateAuthorities()
	EditCertificateAuthority(
		*string,
		*ConfigState.Enum,
		*ConfigField.Enum,
		*string,
	)
	// ToggleCa(string)
}
