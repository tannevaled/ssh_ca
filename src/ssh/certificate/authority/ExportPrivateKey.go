package SSHCertificateAuthority

func (self Struct) ExportPrivateKey() []byte {
	return self.GetSSHKeypairED25519().ExportPrivateKey()
}
