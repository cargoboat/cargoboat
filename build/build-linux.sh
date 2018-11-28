#!/bin/bash
cp ../cargoboat.toml ../bin/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/cargoboat ../