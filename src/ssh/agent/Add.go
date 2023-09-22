package SSHAgent

import (
	"fmt"
	"os"
	"time"

	strftime "github.com/itchyny/timefmt-go"
	"golang.org/x/crypto/ssh/agent"

	SSHCertificate "ssh-ca/src/ssh/certificate"
	"ssh-ca/src/utils"
)

/**
 *
 *
 *
 */
func (self SSHAgentStruct) Add(
	certificate *SSHCertificate.Interface,
) {
	__validBefore := time.Unix(int64((*certificate).GetCertificate().ValidBefore), 0)
	utils.Trace()
	__lifetimeSecs := time.Now().Sub(__validBefore)
	utils.Trace()
	fmt.Printf("%#v\n", __lifetimeSecs)
	__comment := fmt.Sprintf(
		"SSH-CA:%s:VALID_BEFORE:%s",
		(*certificate).GetCertificateAuthorityPublicKeyFingerprintSHA256(),
		strftime.Format(__validBefore, "%Y:%m:%d:%H:%M:%S"),
	)
	utils.Trace()
	__addkey := &agent.AddedKey{
		PrivateKey:  (*certificate).GetSSHKeypairED25519().GetSSHKeypairED25519PrivateKey().GetED25519PrivateKey(),
		Certificate: (*certificate).GetCertificate(),
		//LifetimeSecs: uint32(validHours * 3600),
		LifetimeSecs: uint32(__lifetimeSecs),
		Comment:      __comment,
	}

	if err := (self.agent).Add(*__addkey); err == nil {
		fmt.Printf("Added certificate to ssh-agent, please run ssh-add -l to verify\n")
	} else {
		fmt.Println(fmt.Errorf("Failed to add SSH certificate to ssh-agent: %s\n", err))
		os.Exit(1)
	}
}
