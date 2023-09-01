package SSHAgent

import (
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
)

/*
 * cf https://pkg.go.dev/golang.org/x/crypto/ssh/agent#example-NewClient
 */
func New() SSHAgent {
	__socket := os.Getenv("SSH_AUTH_SOCK")
	__connection, err := net.Dial("unix", __socket)
	if err != nil {
		log.Fatalf("Failed to open SSH_AUTH_SOCK: %v", err)
	}
	__agent := agent.NewClient(__connection)
	__self := &SSHAgentStruct{
		socket:     __socket,
		connection: __connection,
		agent:      __agent,
	}
	return __self
}
