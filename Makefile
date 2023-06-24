all:
	mkdir -p Escrow Card History
	abigen --abi=abis/Escrow.abi --pkg=Escrow --out=Escrow/Escrow.go
	abigen --abi=abis/Card.abi --pkg=Card --out=Card/Card.go
	abigen --abi=abis/History.abi --pkg=History --out=History/History.go

	go build -v -ldflags "-X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(date)'" && go test -v && go vet
