package Config

import (
	"database/sql"
)

type CertificateAuthority struct {
	ID           int64
	Enabled      bool
	Name         string
	Description  string
	PrivateKey   string
	PublicKey    string
	CreationDate string
}

type Interface interface {
	AddCa(string, string, string, string)
	ShowCa(string)
	ListCa()
	EditCa(string)
	ToggleCa(string)
}

type Struct struct {
	config_file_path *string
	db               *sql.DB
}
