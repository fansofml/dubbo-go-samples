# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, softwarek
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PROJECT_NAME = $(shell basename "$(PWD)")
PID = /tmp/.$(PROJECT_NAME).pid
PROJECT_DIR=$(shell pwd)
BASE_DIR := $(PROJECT_DIR)/go-server/dist
DOCKER_DIR := ./integrate_test/dockercompose

SOURCES = $(wildcard $(PROJECT_DIR)/go-server/cmd/*.go)
GO = go
GO_PATH = $(shell $(GO) env GOPATH)
GO_OS = $(shell $(GO) env GOOS)
ifeq ($(GO_OS), darwin)
    GO_OS = mac
endif

# shell
SHELL = /bin/bash

# license
GO_LICENSE_CHECKER_DIR = license-header-checker-$(GO_OS)
GO_LICENSE_CHECKER = $(GO_PATH)/bin/license-header-checker
LICENSE_DIR = /tmp/tools/license

export GO111MODULE ?= on
export GOSUMDB ?= sum.golang.org
export GOARCH ?= amd64
export GONOPROXY ?= **.gitee.com**

OS := $(shell uname)
ifeq ($(OS), Linux)
	export GOOS ?= linux
else ifeq ($(OS), Darwin)
	export GOOS ?= darwin
else
	export GOOS ?= windows
endif

ifeq ($(GOOS), windows)
	export EXT_NAME ?= .exe
else
	export EXT_NAME ?=
endif

CGO ?= 0
ifeq ($(DEBUG), true)
	BUILD_TYPE := debug
	GCFLAGS := -gcflags="all=-N -l"
	LCFLAGS :=
else
	BUILD_TYPE := release
	LDFLAGS := "-s -w"
endif

OUT_DIR := $(BASE_DIR)/$(GOOS)_$(GOARCH)/$(BUILD_TYPE)
LOG_FILE := $(OUT_DIR)/$(PROJECT_NAME).log

export APP_LOG_CONF_FILE ?= $(OUT_DIR)/conf/log.yml

.PHONY: license
license:
	rm -rf coverage.txt
	rm -rf license-header-checker*
	$(GO_LICENSE_CHECKER) -version || (wget https://github.com/lsm-dev/license-header-checker/releases/download/v1.2.0/$(GO_LICENSE_CHECKER_DIR).zip -O $(GO_LICENSE_CHECKER_DIR).zip && unzip -o $(GO_LICENSE_CHECKER_DIR).zip && mkdir -p $(GO_PATH)/bin/ && cp $(GO_LICENSE_CHECKER_DIR)/64bit/license-header-checker $(GO_PATH)/bin/)
	ls /tmp/tools/license/license.txt || wget -P $(LICENSE_DIR) https://github.com/dubbogo/resources/raw/master/tools/license/license.txt
	$(GO_LICENSE_CHECKER) -v -a -r -i vendor $(LICENSE_DIR)/license.txt . go && [[ -z `git status -s` ]]

.PHONY: all
all: help
help: $(realpath $(firstword $(MAKEFILE_LIST)))
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## build: Build application's binaries
.PHONY: build
build: $(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME)

.PHONY: $(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME)
$(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME):
	$(info   >  Building application binary: $(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME))
	@CGO_ENABLED=$(CGO) GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GCFLAGS) -ldflags=$(LDFLAGS) -o $(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME) $(SOURCES)

## docker-health-check: check services health on docker
.PHONY: docker-health-check
docker-health-check:
	$(info   >  run docker health check with $(INTEGRATE_DIR)/docker/docker-health-check.sh)
	@-test -f $(INTEGRATE_DIR)/docker/docker-health-check.sh && bash -f $(INTEGRATE_DIR)/docker/docker-health-check.sh

## docker-up: Shutdown dependency services on docker
.PHONY: docker-up
docker-up:
	$(info   >  Starting dependency services with $(DOCKER_DIR)/docker-compose.yml)
	@docker-compose -f $(DOCKER_DIR)/docker-compose.yml up -d

## docker-down: Shutdown dependency services on docker
.PHONY: docker-down
docker-down:
	$(info   >  Stopping dependency services with $(DOCKER_DIR)/docker-compose.ym)
	@docker-compose -f $(DOCKER_DIR)/docker-compose.yml down

## clean: Clean up the output and the binary of the application
.PHONY: clean
clean: stop
	$(info   >  Cleanning up $(OUT_DIR))
	@-rm -rf $(OUT_DIR)
	@-rm $(PID)

## start: Start the application (for server)
.PHONY: start
start: export DUBBO_GO_CONFIG_PATH ?= $(PROJECT_DIR)/go-server/conf/dubbogo.yml
start: build
	$(info   >  Starting application $(PROJECT_NAME), output is redirected to $(LOG_FILE))
	@ls $(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME)
	@-$(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME) > $(LOG_FILE) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

## start: print application log (for server)
.PHONY: print-server-log
print-server-log:
	$(info   >  print server log with $(LOG_FILE))
	@-cat $(LOG_FILE)

## run: Run the application (for client)
.PHONY: run
run: build
	$(info   >  Running application $(PROJECT_NAME), output is redirected to $(LOG_FILE))
	@-$(OUT_DIR)/$(PROJECT_NAME)$(EXT_NAME) 2>&1 | tee $(LOG_FILE)

## stop: Stop running the application (for server)
.PHONY: stop
stop:
	$(info   >  Stopping the application $(PROJECT_NAME))
	@cat $(PID) | sed "/^/s/^/  \>  Killing PID: /"
	@-kill `cat $(PID)` 2>/dev/null || true

## integration: Run integration test for this application
.PHONY: integration
integration: export DUBBO_GO_CONFIG_PATH ?= $(PROJECT_DIR)/go-client/conf/dubbogo.yml
integration:
	$(info   >  Running integration test for application $(INTEGRATE_DIR))
	@go clean -testcache
	@go test -tags integration -v $(INTEGRATE_DIR)/tests/integration/...

## integration-java: Integration java test (for client)
.PHONY: integration-java
integration-java:
	@$(info   >  Running java test $(INTEGRATE_DIR), project name $(PROJECT_NAME))
	@chmod +x $(INTEGRATE_DIR)/tests/java/run.sh
	@$(shell )$(INTEGRATE_DIR)/tests/java/run.sh $(PROJECT_DIR)/java-client
	@exit $(.SHELLSTATUS)
