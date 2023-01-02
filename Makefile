deps:
	go mod download

install:
	go build -o ${HOME}/.local/bin/qas ./cmd/qas/main.go

lint:
	golint ./...

test:
	go test -v ./...

vet:
	go vet ./...

grab:
	go run cmd/qas/main.go --grab

imports:
	goimports -l -w .

.PHONY: imports grab vet test lint install deps
