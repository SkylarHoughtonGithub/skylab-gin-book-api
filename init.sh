#!/bin/bash
set -e

# Use the first argument as APP_NAME, or fall back to "myapp" if not provided
NAME=$1

echo "init module $NAME..."

# Download dependencies
go mod init $NAME
go mod tidy
