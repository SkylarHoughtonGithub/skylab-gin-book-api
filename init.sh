#!/bin/bash
set -e

NAME=$1

echo "init module $NAME..."

# Download dependencies
go mod init $NAME
go mod tidy
