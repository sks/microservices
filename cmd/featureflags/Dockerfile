ARG VERSION=unknown

FROM golang:alpine as builder

WORKDIR /go/src/github.com/sks/microservices

COPY . .

RUN go build \
    -ldflags "-w -s" \
    -o /tmp/featureflags cmd/featureflags/main.go

FROM alpine

ARG VERSION

COPY --from=builder /tmp/featureflags /bin/featureflags

ENV VERSION=${VERSION}

ENTRYPOINT [ "/bin/featureflags" ]