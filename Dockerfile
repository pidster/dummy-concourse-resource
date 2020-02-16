FROM golang:1.13.8-alpine3.11 AS builder

WORKDIR /go/src/github.com/pidster/dummy-concourse-resource

COPY . .

ENV CGO_ENABLED=0

RUN GOARCH=amd64 GOOS=linux \
 && go test -v models_test.go \
 && go test -v check/check_test.go \
 && go test -v in/in_test.go \
 && go test -v out/out_test.go \
 && echo "TESTS COMPLETE"

RUN GOARCH=amd64 GOOS=linux \
 && go build -v -o outputs/check check/check.go \
 && go build -v -o outputs/in in/in.go \
 && go build -v -o outputs/out out/out.go \
 && chmod 755 outputs/*

FROM alpine:3.11

RUN apk -U update \
 && apk upgrade --purge \
 && apk add --update bash ca-certificates \
 && rm -rf /var/cache/apk/*

COPY --from=builder /go/src/github.com/pidster/dummy-concourse-resource/outputs/* /opt/resource/
