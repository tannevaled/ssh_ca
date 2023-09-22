package SSHCertificateAuthority

import (
	"os"
	"strings"
)

func (self Struct) SaveOnDisk() {

	os.WriteFile(
		self.GetPrivateKeyFilePath(),
		self.ExportPrivateKey(),
		0600,
	)
	os.WriteFile(
		strings.Join([]string{self.GetPrivateKeyFilePath(), "pub"}, "."),
		self.ExportPublicKey(),
		0644,
	)
}
