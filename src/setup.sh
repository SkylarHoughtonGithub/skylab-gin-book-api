#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 [--refresh] <module-name>"
    exit 1
}

# Check for arguments
if [ "$#" -lt 1 ]; then
    usage
fi

REFRESH=false
MODULE_NAME=""

# Parse arguments
while [ "$#" -gt 0 ]; do
    case "$1" in
        --refresh)
            REFRESH=true
            shift
            ;;
        *)
            MODULE_NAME="$1"
            shift
            ;;
    esac
done

# Validate module name
if [ -z "$MODULE_NAME" ]; then
    usage
fi

# Refresh existing files if --refresh is specified
if [ "$REFRESH" = true ]; then
    echo "Deleting existing Go module files..."
    if [ -f "go.mod" ]; then
        echo "Removing go.mod"
        rm go.mod
    fi
    if [ -f "go.sum" ]; then
        echo "Removing go.sum"
        rm go.sum
    fi
    # Remove other common Go files if needed
    # Example: rm -rf vendor/
fi

# Initialize the Go module
echo "Initializing Go module with name: $MODULE_NAME"
go mod init "$MODULE_NAME"

# Tidy up the module
echo "Running go mod tidy"
go mod tidy

echo "Go module setup completed."
