#!/bin/bash -eux

GOOS=linux go build -o docker-cloner main.go
