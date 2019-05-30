#!/bin/bash -eux

VERSION='0.0.1'

GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -ldflags "-X main.version=$VERSION" -o docker-cloner main.go
#go build -ldflags "-X main.version=$VERSION" -o docker-cloner main.go
