#!/bin/sh
GOOS=linux   GOARCH=amd64 go build -o goHome0     -ldflags="-w -s" main.go
GOOS=windows GOARCH=386   go build -o goHome0.exe -ldflags="-w -s" main.go
# upx --best goHome0 goHome0.exe
upx goHome0 goHome0.exe
