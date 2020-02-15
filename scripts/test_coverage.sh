#!/bin/bash
set -e

go test -coverprofile=coverage.out ./...

# show html coverage if specified, otherwise print function coverage
if [ "$1" == "html" ]; then
    go tool cover -html=coverage.out
else
    go tool cover -func=coverage.out
fi