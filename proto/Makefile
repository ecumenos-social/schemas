GOPRIVATE=github.com/ecumenos-social

SHELL = /bin/bash
GO_GEN_PATH = ./gen

.PHONY: all
all: check lint clean generate go_tidy go_vet go_fmt

.PHONY: clean
clean:
	/bin/rm -rf $(GO_GEN_PATH)/*

.PHONY: check
check:
	buf build

.PHONY: generate
generate:
	buf generate .

.PHONY: lint
lint:
	buf lint .

.PHONY: go_tidy
go_tidy:
	go mod tidy

.PHONY: go_vet
go_vet:
	go vet ./...

.PHONY: go_fmt
go_fmt:
	go fmt ./...
