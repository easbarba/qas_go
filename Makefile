.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64

BINARY_NAME := qas
MAIN := ./main.go
TO := ${HOME}/.local/bin

all: build test

deps:
	go mod download

build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${BINARY_NAME} ${MAIN}

install: build
	mv -v ./${BINARY_NAME} ${TO}/${BINARY_NAME}

lint:
	golangci-lint run --enable-all internal cmd/pak

test:
	go test -v ./...

clean:
	go clean
	rm ${BINARY_NAME}

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

image:
	podman build --file ./Dockerfile --tag $USER/${BINARY_NAME}:$(shell cat .env)

.PHONY: imports grab vet test lint install deps coverage
