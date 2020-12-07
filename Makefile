MODULE  := github.com/philzon/helm-assert
NAME    := assert
BINDIR  := bin
VERSION := $(shell cat plugin.yaml | grep "version" | cut -d '"' -f 2)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date +"%Y-%m-%d %H:%M:%S")

CFLAGS  := -X "$(MODULE)/internal/app.Name=$(NAME)" \
           -X "$(MODULE)/internal/app.Version=$(VERSION)" \
           -X "$(MODULE)/internal/app.Commit=$(COMMIT)" \
           -X "$(MODULE)/internal/app.Date=$(DATE)" \
           -s \
           -w \

.PHONY: all init lint build test package

all: init build

clean:
	@rm -rf $(BINDIR)

init:
	@mkdir --parent $(BINDIR)

lint:
	@go vet ./...

build:
	@go build -o $(BINDIR)/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-all: build-linux build-windows build-darwin

build-linux: build-linux-amd64 build-linux-arm64

build-linux-amd64:
	@mkdir --parent $(BINDIR)/linux-amd64
	@GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/linux-amd64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-linux-arm64:
	@mkdir --parent $(BINDIR)/linux-arm64
	@GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/linux-arm64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-windows: build-windows-amd64

build-windows-amd64:
	@mkdir --parent $(BINDIR)/windows-amd64
	@GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/windows-amd64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

build-darwin: build-darwin-amd64

build-darwin-amd64:
	@mkdir --parent $(BINDIR)/darwin-amd64
	@GOOS=darwin GOARCH=amd64 go build -o $(BINDIR)/darwin-amd64/$(NAME) -ldflags '$(CFLAGS)' cmd/$(NAME)/*.go

test:
	@go test ./...

package:
	@./scripts/package.sh ${VERSION} $(BINDIR)
