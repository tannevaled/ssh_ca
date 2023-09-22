package SSHAgent

import (
	"net"

	SSHCertificate "ssh-ca/src/ssh/certificate"

	"golang.org/x/crypto/ssh/agent"
)

type SSHAgent interface {
	PrintList()
	Add(*SSHCertificate.Interface)
}

type SSHAgentStruct struct {
	socket     string
	connection net.Conn
	agent      agent.ExtendedAgent
}
