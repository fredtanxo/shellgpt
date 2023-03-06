NAME=shellgpt
BINDIR=bin
GOBUILD=CGO_ENABLED=0 go build

all: linux-amd64 darwin-amd64 darwin-arm64 windows-amd64

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe

build:
	$(GOBUILD) -o $(BINDIR)/$(NAME)

lint:
	golangci-lint run ./...

clean:
	rm $(BINDIR)/*

.PHONY: lint clean
