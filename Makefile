SHELL := $(shell which bash)
ENV = /usr/bin/env

.SHELLFLAGS = -c

.ONESHELL: ;
.NOTPARALLEL: ;
.EXPORT_ALL_VARIABLES:

.PHONY: all
.DEFAULT_GOAL := help

LDFLAGS = -ldflags "-w -s"

help: ## Show Help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@exit 1

build-arm6: ## Build for ARMv6 (pizero)
	GOOS=linux GOARCH=arm GOARM=6 go build ${LDFLAGS}
