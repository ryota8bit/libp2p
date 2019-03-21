deps:
	go get -t -d ./...

bld:
	go build -v -o ${GOPATH}/bin/p2p-app ./cmd/main.go
