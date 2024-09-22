#!/bin/bash
set -e

PATH=$1

echo "Building $PATH..."

# Build binary
go run $PATH
