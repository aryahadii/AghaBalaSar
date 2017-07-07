ROOT := github.com/aryahadii/AghaBalaSar
GO_VARS ?= ENABLE_CGO=1 GOOS=darwin GOARCH=amd64
GO ?= go
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME) -X $(ROOT).Title=aghabalasar

.PHONY: help clean update-dependencies dependencies prepare_test_cassandra test docker push deploy test-docker


help:
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  update-dependencies  to update glide.lock (refs to dependencies)"
	@echo "  dependencies         to install the dependencies"
	@echo "  aghabalasar          to build the binary"
	@echo "  clean                to remove generated files"

clean:
	rm -rf aghabalasar

update-dependencies:
	glide up

dependencies:
	glide install

aghabalasar: *.go */*.go */*/*.go glide.lock
	$(GO_VARS) $(GO) build -o="aghabalasar" -ldflags="$(LD_FLAGS)" $(ROOT)/cmd/aghabalasar

push:
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker push $(DOCKER_IMAGE):latest

