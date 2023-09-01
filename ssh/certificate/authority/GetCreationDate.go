package SSHCertificateAuthority

import "time"

func (self Struct) GetCreationDate() time.Time {
	return self.creationDate
}
