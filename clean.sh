#!/bin/sh
find . -name \*~ -type f | xargs -r rm
go fmt ./...
