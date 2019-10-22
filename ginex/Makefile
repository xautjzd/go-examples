GitBranch=$(shell git rev-parse --abbrev-ref HEAD)
GitCommit=$(shell git rev-parse --short HEAD)
Built=$(shell date +"%Y-%m-%d %T%z")
GOBUILD=CGO_ENABLED=0 go build -ldflags "-X 'github.com/xautjzd/ginex/version.gitBranch=$(GitBranch)' -X 'github.com/xautjzd/ginex/version.gitCommit=$(GitCommit)' -X 'github.com/xautjzd/ginex/version.built=$(Built)' -X 'github.com/xautjzd/ginex/version.goVersion=$(shell go version)'"

.PHONY: all
all:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/main main.go

.PHONY: all-macos
all-macos:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/main main.go
