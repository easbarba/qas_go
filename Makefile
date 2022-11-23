deps:
	go mod download

install:
	go build -o ${HOME}/.local/bin/qas ./cmd/qas/main.go

lint:
	golint ./...

vet:
	go vet ./...

grab:
	go run cmd/qas/main.go --grab

imports:
	goimports -l -w .
