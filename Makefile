SHELL       = /bin/bash -euo pipefail
BINPATH     = $(PWD)/bin
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPROXY     = https://proxy.golang.org,https://gocenter.io,direct
MODULE      = $(shell go list -m)
PKGS        = $(shell go list ./... | grep -v vendor)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
SHELL       = /bin/bash -euo pipefail
TIMEOUT     = 1s

export PATH := $(BINPATH):$(PATH)

.DEFAULT_GOAL = test-with-coverage

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(shell go env GOFLAGS)"
	@echo "GOPRIVATE:   $(shell go env GOPRIVATE)"
	@echo "GOPROXY:     $(shell go env GOPROXY)"
	@echo "GONOPROXY:   $(shell go env GONOPROXY)"
	@echo "GOSUMDB:     $(shell go env GOSUMDB)"
	@echo "GONOSUMDB:   $(shell go env GONOSUMDB)"
	@echo "MODULE:      $(MODULE)"
	@echo "PATH:        $(PATH)"
	@echo "PATHS:       $(PATHS)"
	@echo "SHELL:       $(SHELL)"
	@echo "TIMEOUT:     $(TIMEOUT)"


.PHONY: deps
deps: deps-main deps-tools

.PHONY: deps-main
deps-main:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: deps-tools
deps-tools:
	@cd tools && make

.PHONY: deps-update
deps-update:
	@go get -mod= -u all


.PHONY: format
format:
	@goimports -local $(dir $(shell go list -m)) -ungroup -w $(PATHS)

.PHONY: generate
generate:
	@go generate $(PKGS)

.PHONY: refresh
refresh: deps-update deps generate format test-with-coverage


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PKGS)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PKGS) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PKGS)
