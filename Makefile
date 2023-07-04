build: get
	mkdir -p output
	go build -o output/ssh-ca ./cmd/ssh-ca
	chmod +x output/ssh-ca

get:
	go get golang.org/x/crypto/ed25519
	go get golang.org/x/crypto/ssh
	go get golang.org/x/crypto/ssh/agent
	go get github.com/fatih/color
	go get github.com/itchyny/timefmt-go

gen:
	ssh-keygen -t ed25519 -f output/ca

test: FORCE
	./output/ssh-ca -cakeypath "$$(pwd)/output/ca"

FORCE: ;