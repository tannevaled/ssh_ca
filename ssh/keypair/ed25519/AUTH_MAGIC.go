package SSHKeypairED25519

/*
 * https://cvsweb.openbsd.org/src/usr.bin/ssh/PROTOCOL.key?annotate=HEAD
 *
 * AUTH_MAGIC is a hard-coded, null-terminated string,
 * #define AUTH_MAGIC      "openssh-key-v1"
 * byte[]  AUTH_MAGIC
 */
func (self SSHKeypairED25519Struct) AUTH_MAGIC() []byte {
	return append([]byte("openssh-key-v1"), 0)
}
