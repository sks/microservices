ARG VERSION=unknown

FROM golang:alpine as builder

WORKDIR /go/src/github.com/sks/microservices

COPY . .

RUN go build \
    -ldflags "-w -s" \
    -o /tmp/smsotpverifier cmd/smsotpverifier/main.go

FROM alpine

ARG VERSION

COPY --from=builder /tmp/smsotpverifier /bin/smsotpverifier

ENV VERSION=${VERSION}

ENTRYPOINT [ "/bin/smsotpverifier" ]