package SSHKeypairED25519

import "golang.org/x/crypto/ssh"

/*
 *
 * Generate the pubkey prefix "\0\0\0\nssh-ed25519\0\0\0 "
 *
 */
func (self SSHKeypairED25519Struct) MarshaledPubkeyPrefix() []byte {
	var __char_zero byte = 0x0
	var __char_newline byte = 0x0b
	var __char_space byte = 0x20

	__prefix := []byte{
		__char_zero,
		__char_zero,
		__char_zero,
		__char_newline,
	}
	__prefix = append(
		__prefix,
		[]byte(ssh.KeyAlgoED25519)...,
	)
	__prefix = append(
		__prefix,
		[]byte{
			__char_zero,
			__char_zero,
			__char_zero,
			__char_space,
		}...,
	)

	return __prefix
}
