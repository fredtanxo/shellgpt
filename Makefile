NAME=shellgpt
BINDIR=bin
GOBUILD=CGO_ENABLED=0 go build

build:
	$(GOBUILD) -o $(BINDIR)/$(NAME)

lint:
	golangci-lint run ./...

test:
	go test

clean:
	rm $(BINDIR)/*

all: linux-amd64 darwin-amd64 darwin-arm64 windows-amd64

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

.PHONY: lint test clean
