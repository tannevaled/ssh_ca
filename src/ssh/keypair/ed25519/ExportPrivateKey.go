package SSHKeypairED25519

import "encoding/pem"

func (self SSHKeypairED25519Struct) ExportPrivateKey() []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:    "OPENSSH PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   self.MarshalED25519PrivateKey(),
	})
}
