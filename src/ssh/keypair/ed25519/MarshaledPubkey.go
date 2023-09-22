package SSHKeypairED25519

func (self SSHKeypairED25519Struct) MarshaledPubkey() []byte {
	return append(
		self.MarshaledPubkeyPrefix(),
		[]byte(self.GetSSHKeypairED25519PublicKey().GetED25519PublicKey())...,
	)
}
