FROM golang:1.13-alpine3.11 AS builder

WORKDIR /go/src/github.com/pidster/dummy-concourse-resource

COPY . .

RUN GOARCH=amd64 GOOS=linux && \
    go build -o outputs/in in/in.go && \
    go build -o outputs/out out/out.go && \
    go build -o outputs/check check/check.go

FROM alpine:3.11

RUN apk -U update \
 && apk upgrade --purge \
 && rm -rf /var/cache/apk/* 

COPY --from=builder /go/src/github.com/pidster/dummy-concourse-resource/outputs/* /opt/resource/
