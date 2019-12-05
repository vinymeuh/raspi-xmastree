SHELL := $(shell which bash)
ENV = /usr/bin/env

.SHELLFLAGS = -c

.SILENT: ;
.ONESHELL: ;
.NOTPARALLEL: ;
.EXPORT_ALL_VARIABLES:

.PHONY: all
.DEFAULT: help

help: ## Show Help
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-pizerow: ## Build the app
	GOOS=linux GOARCH=arm GOARM=6 go build

pizerow: build-pizerow	## Copy the app on pizerow.local
	scp xmastree pizerow.local:/tmp
	scp xmastree.openrc pizerow.local:/tmp
