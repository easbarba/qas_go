.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64
NAME := qas
TO := ${HOME}/.local/bin

deps:
	go mod download

build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${NAME} ./cmd/qas/main.go

install: build
	mv -v ./${NAME} ${TO}/${NAME}

lint:
	golangci-lint run --enable-all internal cmd/pak

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

coverage:
	go test --cover ./... -coverprofile=coverage.out

.PHONY: imports grab vet test lint install deps coverage
