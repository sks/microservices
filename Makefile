proto:
	@mkdir -p ${PWD}/vendor
	protoc -I ${PWD} --go_out=plugins=grpc:${PWD} ${PWD}/protos/*.proto

build-%:
	docker build \
		--no-cache \
		-f cmd/$*/Dockerfile \
		-t $*:latest .

build: proto build-featureflags
