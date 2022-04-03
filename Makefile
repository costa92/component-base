GO :=go
ROOT_PACKAGE := github.com/costa92/component-base
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(shell pwd)
endif

lint:
	golangci-lint run  -v