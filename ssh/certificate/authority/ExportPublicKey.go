package SSHCertificateAuthority

func (self Struct) ExportPublicKey() []byte {
	return self.GetSSHKeypairED25519().ExportPublicKey()
}
