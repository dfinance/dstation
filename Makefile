#!/usr/bin/make -f

# Options
LEDGER_ENABLED := false

# Common vars
DOCKER := $(shell which docker)

# Build version

git_tag := $(shell git describe --tags --abbrev=0)
git_commit := $(shell git rev-list -1 HEAD)
app_version := $(git_tag)-$(git_commit)
cosmos_version := $(shell awk '/github.com\/cosmos\/cosmos-sdk/ {print $$NF}' go.mod)
tm_version := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')

export GO111MODULE = on

# Process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for Ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling Ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for Ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# Process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=dnode \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=dstation \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(version) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(git_commit) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(tm_version)

ldflags := $(strip $(ldflags))

# Makefile rules

export GO111MODULE=on

all: install

install: go.sum install-dstation

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

install-dstation:
	@echo "--> dstation app build / install"
	@echo "  App version: $(version)"
	@echo "  SDK version: $(cosmos_version)"
	@echo "  TM version:  $(tm_version)"

	@go install -ldflags "$(ldflags)" -tags "$(build_tags)" ./cmd/dstation

lint:
	@echo "--> Running Golang linter (unused variable / function warning are skipped)"
	@golangci-lint run --exclude 'unused'

proto-gen:
	@echo "--> Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh
