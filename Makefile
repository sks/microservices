.DEFAULT_GOAL := all

all: proto build

COMMIT_SHA := $(shell git rev-parse --short=7 HEAD)
GIT_DIRTY  := $(shell test -n "`git status --porcelain`" && echo "-dirty" || echo "")
DOCKER_TAG = $(COMMIT_SHA)$(GIT_DIRTY)

proto:
	@mkdir -p ${PWD}/vendor
	protoc -I ${PWD} --go_out=plugins=grpc:${PWD} ${PWD}/protos/*.proto

build-%:
	docker build \
		--no-cache \
		-f cmd/$*/Dockerfile \
		-t $*:$(DOCKER_TAG) .

build: proto build-featureflags
