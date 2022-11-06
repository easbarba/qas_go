deps:
	go

install:
	go build -o qas

lint:
	golint ./...

vet:
	go vet ./...

imports:
	goimports -l -w .
