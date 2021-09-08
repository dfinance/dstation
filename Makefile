#!/usr/bin/make -f

# Options
LEDGER_ENABLED ?= true

# Common vars
DOCKER := $(shell which docker)
PROTOC := $(shell which protoc)

# Build version

git_tag := $(shell git describe --tags --abbrev=0)
ifeq ($(git_tag),)
  git_tag := v0.0.0
endif
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
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(git_tag) \
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
	@echo "  App version: $(app_version)"
	@echo "  SDK version: $(cosmos_version)"
	@echo "  TM version:  $(tm_version)"
	@echo "  Ledger: $(LEDGER_ENABLED)"
	@go install -ldflags "$(ldflags)" -tags "$(build_tags)" ./cmd/dstation

lint:
	@echo "--> Running Golang linter (unused variable / function warning are skipped)"
	@golangci-lint run --exclude 'unused'

tests:
	@echo "--> Running tests"
	go clean -testcache
	go test ./... -v

proto-gen:
	@echo "--> Generating DVM Protobuf files"
	@#mkdir -p $(CURDIR)/pkg/types/dvm_proto/types_grpc
	@#mkdir -p $(CURDIR)/pkg/types/dvm_proto/compiler_grpc
	@#mkdir -p $(CURDIR)/pkg/types/dvm_proto/ds_grpc
	@#mkdir -p $(CURDIR)/pkg/types/dvm_proto/metadata_grpc
	@#mkdir -p $(CURDIR)/pkg/types/dvm_proto/vm_grpc
	@#$(PROTOC) --proto_path=$(CURDIR)/third_party/proto/dvm_proto --go_out=plugins=grpc:$(CURDIR)/pkg/types/dvm_proto/types_grpc --go_opt=paths=source_relative $(CURDIR)/third_party/proto/dvm_proto/common-types.proto
	@#$(PROTOC) --proto_path=$(CURDIR)/third_party/proto/dvm_proto --go_out=plugins=grpc:$(CURDIR)/pkg/types/dvm_proto/compiler_grpc --go_opt=paths=source_relative $(CURDIR)/third_party/proto/dvm_proto/compiler.proto
	@#$(PROTOC) --proto_path=$(CURDIR)/third_party/proto/dvm_proto --go_out=plugins=grpc:$(CURDIR)/pkg/types/dvm_proto/ds_grpc --go_opt=paths=source_relative $(CURDIR)/third_party/proto/dvm_proto/data-source.proto
	@#$(PROTOC) --proto_path=$(CURDIR)/third_party/proto/dvm_proto --go_out=plugins=grpc:$(CURDIR)/pkg/types/dvm_proto/metadata_grpc --go_opt=paths=source_relative $(CURDIR)/third_party/proto/dvm_proto/metadata.proto
	@#$(PROTOC) --proto_path=$(CURDIR)/third_party/proto/dvm_proto --go_out=plugins=grpc:$(CURDIR)/pkg/types/dvm_proto/vm_grpc --go_opt=paths=source_relative $(CURDIR)/third_party/proto/dvm_proto/vm.proto

	@echo "--> Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh

## Binaries build using xgo (https://github.com/karalabe/xgo) fork as xgo seems abandoned
build-binaries: go.sum
	@echo "--> dstation app cross-build (windows/amd64, linux/amd64, darwin/amd64)"
	@echo "  App version: $(app_version)"
	@echo "  SDK version: $(cosmos_version)"
	@echo "  TM version:  $(tm_version)"
	@mkdir -p ./builds
	@go get github.com/crazy-max/xgo
	xgo -go='1.15.13' --ldflags='$(ldflags)' --tags='$(build_tags)' --out='./builds/dstation-${app_version}' -targets='windows/amd64,linux/amd64,darwin/amd64' ./cmd/dstation