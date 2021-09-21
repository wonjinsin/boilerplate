PACKAGE = pikachu
CUSTOM_OS = ${GOOS}
BASE_PATH = $(shell pwd)
BIN = $(BASE_PATH)/bin
BINARY_NAME = bin/$(PACKAGE)
MAIN = $(BASE_PATH)/main.go
MOCK = $(BIN)/mockgen

ifneq (, $(CUSTOM_OS))
	OS ?= $(CUSTOM_OS)
else
	OS ?= $(shell uname | awk '{print tolower($0)}')
endif
build:
	GOOS=$(OS) go build -o $(BINARY_NAME) $(MAIN)

.PHONY: test
test:
	go test -v ./...

build-mocks:
	GOBIN=$(BIN) go get github.com/golang/mock/mockgen
	$(MOCK) -source=service/service.go -destination=mock/mock_service.go -package=mock
	$(MOCK) -source=repository/repository.go -destination=mock/mock_repository.go -package=mock

.PHONY: vendor
vendor: build-mocks
	go mod vendor

