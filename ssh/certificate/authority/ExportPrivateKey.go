package SSHCertificateAuthority

func (self Struct) ExportPrivateKey() []byte {
	//	return self.getSshKeypairEd25519().exportPrivateKey(map[string]string{
	//		"SSHCertificateAuthorityManager":      fmt.Sprintf("%q", "github.io/tannevaled/ssh-ca"),
	//		"SSHCertificateAuthorityName":         fmt.Sprintf("%q", self.getName()),
	//		"SSHCertificateAuthorityCreationDate": fmt.Sprintf("%q", self.getCreationDate()),
	//	})
	return self.GetSSHKeypairED25519().ExportPrivateKey()
}
