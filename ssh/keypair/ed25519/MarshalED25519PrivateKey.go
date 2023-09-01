package SSHKeypairED25519

import (
	mathRand "math/rand"

	"golang.org/x/crypto/ssh"
)

/*
 * cf https://cvsweb.openbsd.org/src/usr.bin/ssh/PROTOCOL.key?annotate=HEAD
 * cf https://dnaeon.github.io/openssh-private-key-binary-format/
 *
 * Writes ed25519 private keys into the new OpenSSH private key format.
 */
func (self SSHKeypairED25519Struct) MarshalED25519PrivateKey() []byte {
	// Add our key header (followed by a null byte)
	//const AUTH_MAGIC []byte = append([]byte("openssh-key-v1"), 0)

	var w struct {
		CipherName   string
		KdfName      string
		KdfOpts      string
		NumKeys      uint32 // 1
		PubKey       []byte
		PrivKeyBlock []byte
	}

	// Fill out the private key fields
	pk1 := struct {
		Check1  uint32
		Check2  uint32
		Keytype string
		Pub     []byte
		Priv    []byte
		Comment string
		Pad     []byte `ssh:"rest"`
	}{}

	// Set our check ints
	ci := mathRand.Uint32()
	pk1.Check1 = ci
	pk1.Check2 = ci

	// Set our key type
	pk1.Keytype = ssh.KeyAlgoED25519

	// Add the pubkey to the optionally-encrypted block
	pk := self.GetSSHKeypairED25519PublicKey().GetED25519PublicKey()
	pubKey := []byte(pk)
	pk1.Pub = pubKey

	// Add our private key
	pk1.Priv = []byte(self.GetSSHKeypairED25519PrivateKey().GetED25519PrivateKey())

	// Might be useful to put something in here at some point
	pk1.Comment = ""

	// Add some padding to match the encryption block size within PrivKeyBlock (without Pad field)
	// 8 doesn't match the documentation, but that's what ssh-keygen uses for unencrypted keys. *shrug*
	bs := 8
	blockLen := len(ssh.Marshal(pk1))
	padLen := (bs - (blockLen % bs)) % bs
	pk1.Pad = make([]byte, padLen)

	// Padding is a sequence of bytes like: 1, 2, 3...
	for i := 0; i < padLen; i++ {
		pk1.Pad[i] = byte(i + 1)
	}

	// Generate the pubkey prefix "\0\0\0\nssh-ed25519\0\0\0 "
	//prefix := []byte{0x0, 0x0, 0x0, 0x0b}
	//prefix = append(prefix, []byte(ssh.KeyAlgoED25519)...)
	//prefix = append(prefix, []byte{0x0, 0x0, 0x0, 0x20}...)

	// Only going to support unencrypted keys for now
	w.CipherName = "none"
	w.KdfName = "none"
	w.KdfOpts = ""
	w.NumKeys = 1
	//w.PubKey = append(self.marshaledPubkeyPrefix(), pubKey...)
	w.PubKey = self.MarshaledPubkey()
	w.PrivKeyBlock = ssh.Marshal(pk1)

	return append(self.AUTH_MAGIC(), ssh.Marshal(w)...)
}
