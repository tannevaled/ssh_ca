package SSHKeypairED25519

import "golang.org/x/crypto/ssh"

func (self SSHKeypairED25519Struct) ExportPublicKey() []byte {
	return ssh.MarshalAuthorizedKey(self.GetSSHKeypairED25519PublicKey().GetSSHPublicKey())
}
