NAME := qas
TO := ${HOME}/.local/bin

deps:
	go mod download

install:
	go build -o ${TO}/${NAME} ./cmd/qas/main.go

lint:
	golint ./...

test:
	go test -v ./...

vet:
	go vet ./...

grab:
	go run cmd/qas/main.go --grab

archive:
	go run cmd/qas/main.go --archive meh,forevis,tar

imports:
	goimports -l -w .

.PHONY: imports grab vet test lint install deps
