#!/usr/bin/bash

. ./env.sh

CGO_ENABLED=1 go run -race ./cmd/*.go
