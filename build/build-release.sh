#!/usr/bin/env bash

NAME="cargoboat"
cd ../
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./releases/${NAME}-win32.exe ./cmd
echo "build win32 succeed!"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./releases/${NAME}-win64.exe ./cmd
echo "build win64 succeed!"
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ./releases/${NAME}-darwin32 ./cmd
echo "build darwin32 succeed!"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./releases/${NAME}-darwin64 ./cmd
echo "build darwin64 succeed!"
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./releases/${NAME}-linux32 ./cmd
echo "build linux32 succeed!"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./releases/${NAME}-linux64 ./cmd
echo "build linux64 succeed!"