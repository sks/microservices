.DEFAULT_GOAL := all

all: grpc build

COMMIT_SHA := $(shell git rev-parse --short=7 HEAD)
GIT_DIRTY  := $(shell test -n "`git status --porcelain`" && echo "-dirty" || echo "")
DOCKER_TAG = $(COMMIT_SHA)$(GIT_DIRTY)

_proto-%:
	protoc -I ${PWD} \
		-I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:${PWD} ${PWD}/protos/$*-service.proto

proto: _proto-featureflag _proto-smsotpverifier

_reverseproxy-%:
	protoc \
		-I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		${PWD}/protos/$*-service.proto

reverseproxy: _reverseproxy-featureflag _reverseproxy-smsotpverifier

grpc: proto reverseproxy

build-%:
	docker build \
		--no-cache \
		--build-arg VERSION=${DOCKER_TAG}  \
		-f cmd/$*/Dockerfile \
		-t $*:$(DOCKER_TAG) .

build: build-featureflags build-smsotpverifier
