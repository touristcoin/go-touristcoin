.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/touristcoin/go-touristcoin/cmd/swarm
	go build -o bin/geth     github.com/touristcoin/go-touristcoin/cmd/geth
	go build -o bin/bootnode github.com/touristcoin/go-touristcoin/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/touristcoin/go-touristcoin/swarm/...
