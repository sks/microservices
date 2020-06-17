FROM golang:alpine as builder

WORKDIR /go/src/github.com/sks/microservices

COPY . .

RUN go build \
    -ldflags "-w -s" \
    -o /tmp/featureflags cmd/featureflags/main.go

FROM alpine

COPY --from=builder /tmp/featureflags /bin/featureflags

ENTRYPOINT [ "/bin/featureflags" ]