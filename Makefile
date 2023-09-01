build: get
	mkdir -p output
	go build -o output/ssh-ca
	#go build -o output/ssh-ca .
	chmod +x output/ssh-ca

get:
	go get golang.org/x/crypto/ed25519
	go get golang.org/x/crypto/ssh
	go get golang.org/x/crypto/ssh/agent
	#go get github.com/fatih/color
	go get github.com/itchyny/timefmt-go
	go get github.com/jedib0t/go-pretty/v6/table
	go get github.com/jedib0t/go-pretty/v6/text
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	#go get -u gorm.io/gorm
	#go get -u gorm.io/driver/sqlite

gen:
	ssh-keygen -t ed25519 -f output/ca

test: FORCE
	./output/ssh-ca -cakeypath "$$(pwd)/output/ca"

FORCE: ;