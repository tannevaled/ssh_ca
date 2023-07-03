build: get
	mkdir -p output
	go build -v -o output/ssh-ca ./cmd/ssh-ca
	chmod +x output/ssh-ca

get:
	go get golang.org/x/crypto/ed25519
	go get golang.org/x/crypto/ssh
	go get golang.org/x/crypto/ssh/agent
