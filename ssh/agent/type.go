package SSHAgent

import (
	"net"

	SSHCertificate "ssh-ca/ssh/certificate"

	"golang.org/x/crypto/ssh/agent"
)

type SSHAgent interface {
	PrintList()
	Add(SSHCertificate.SSHCertificate, int)
}

type SSHAgentStruct struct {
	socket     string
	connection net.Conn
	agent      agent.ExtendedAgent
}
