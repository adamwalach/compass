APP_NAME = compass-tenant-fetcher
APP_PATH = components/tenant-fetcher
BUILDPACK = eu.gcr.io/kyma-project/test-infra/buildpack-golang-toolbox:v20210125-6234473e
SCRIPTS_DIR = $(realpath $(shell pwd)/../..)/scripts
export GO111MODULE = on
export SKIP_STEP_MESSAGE = "Do nothing for Go modules project"

include $(SCRIPTS_DIR)/generic_make_go.mk

errcheck-local:
	errcheck -blank -asserts -ignoregenerated ./...

test-local:
	@echo ${SKIP_STEP_MESSAGE}

resolve-local:
	@echo ${SKIP_STEP_MESSAGE}

ensure-local:
	@echo ${SKIP_STEP_MESSAGE}

dep-status-local:
	@echo ${SKIP_STEP_MESSAGE}

build-local:
	env go build -o bin/tenant-fetcher ./cmd/main.go