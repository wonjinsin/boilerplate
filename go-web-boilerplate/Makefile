PACKAGE = pikachu
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BINARY_NAME = bin/$(PACKAGE)
MAIN = $(BASE_PATH)/main.go

ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif

.PHONY: vendor
vendor:
	go mod vendor

build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

test-all:
	go test -v ./...

