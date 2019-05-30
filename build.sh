#!/bin/bash -eux

VERSION='0.0.1'

export GO111MODULE=on

GOOS=linux go build -ldflags "-X main.version=$VERSION" -o docker-cloner main.go
#go build -ldflags "-X main.version=$VERSION" -o docker-cloner main.go
