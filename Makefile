MODULE  := github.com/philzon/helm-assert
NAME    := assert
BINDIR  := bin
INSDIR  := /usr/local/bin
VERSION := $(shell cat plugin.yaml | grep "version" | cut -d '"' -f 2)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date +"%Y-%m-%d %H:%M:%S")

CFLAGS  := -X "$(MODULE)/internal/app.Name=$(NAME)" \
           -X "$(MODULE)/internal/app.Version=$(VERSION)" \
           -X "$(MODULE)/internal/app.Commit=$(COMMIT)" \
           -X "$(MODULE)/internal/app.Date=$(DATE)" \
           -s \
           -w \

.PHONY: all init lint build build-linux-amd64 build-linux-arm64 build-windows-amd64 build-darwin-amd64 test functional-test package install uninstall

all: init build

clean:
	@rm -rf $(BINDIR)

init:
	@mkdir --parent $(BINDIR)

dependencies:
	@go mod download

lint:
	@go vet ./...

build:
	@cp plugin.yaml $(BINDIR)
	@cp LICENSE.txt $(BINDIR)
	@go build -o $(BINDIR)/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-all: init build-linux-amd64 build-linux-arm64 build-windows-amd64 build-darwin-amd64

build-linux-amd64:
	@mkdir --parent $(BINDIR)/linux-amd64
	@cp plugin.yaml $(BINDIR)/linux-amd64
	@cp LICENSE.txt $(BINDIR)/linux-amd64
	@GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/linux-amd64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-linux-arm64:
	@mkdir --parent $(BINDIR)/linux-arm64
	@cp plugin.yaml $(BINDIR)/linux-arm64
	@cp LICENSE.txt $(BINDIR)/linux-arm64
	@GOOS=linux GOARCH=arm64 go build -o $(BINDIR)/linux-arm64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-windows-amd64:
	@mkdir --parent $(BINDIR)/windows-amd64
	@cp plugin.yaml $(BINDIR)/windows-amd64
	@cp LICENSE.txt $(BINDIR)/windows-amd64
	@GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/windows-amd64/$(NAME).exe -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-darwin-amd64:
	@mkdir --parent $(BINDIR)/darwin-amd64
	@cp plugin.yaml $(BINDIR)/darwin-amd64
	@cp LICENSE.txt $(BINDIR)/darwin-amd64
	@GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/darwin-amd64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

test: unit-test functional-test

unit-test:
	@go test ./...

functional-test:
	@$(BINDIR)/assert tests/assert.yaml tests/chart-example/
	@$(BINDIR)/assert tests/test.yaml tests/chart-example/

package:
	@./scripts/package.sh $(BINDIR)

install:
	@mkdir -p $(INSDIR)
	@cp ./$(BINDIR)/$(NAME) $(INSDIR)/$(NAME)
	@echo "Installed $(INSDIR)/$(NAME)"

uninstall:
	@rm $(INSDIR)/$(NAME)
	@echo "Uninstalled $(INSDIR)/$(NAME)"
