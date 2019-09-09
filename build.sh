#!/bin/sh
GOOS=linux   GOARCH=amd64 go build -o goHome0-linux-amd64       -ldflags="-w -s" main.go
GOOS=windows GOARCH=386   go build -o goHome0-windows-386.exe   -ldflags="-w -s" main.go
GOOS=windows GOARCH=amd64 go build -o goHome0-windows-amd64.exe -ldflags="-w -s" main.go
# upx --best goHome0 goHome0.exe
upx *.exe goHome0-linux-amd64
