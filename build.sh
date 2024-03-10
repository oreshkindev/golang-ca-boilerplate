#!/usr/bin/bash

GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go
