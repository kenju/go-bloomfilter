# https://www.gnu.org/software/make/manual/make.html#index-_002eEXPORT_005fALL_005fVARIABLES
.EXPORT_ALL_VARIABLES:
GO111MODULE=on

NAME := app
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Run tests
test:
	gotest ./... -v

## Run Benchmark tests
bench:
	gotest ./... -v -bench . -benchmem

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup: update-submodule
	# get development tools
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/rakyll/gotest

## Update go modules
update:
	go get -u
	go mod tidy

## Format source codes
fmt:
	for d in $$(go list ./...); do \
		goimports -w ${GOPATH}/src/$$d; \
	done


