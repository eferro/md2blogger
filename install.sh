#!/bin/bash

set -e

echo "Installing md2blogger..."

if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

make build
make install

echo ""
echo "md2blogger installed successfully!"
echo ""
echo "Usage:"
echo "  md2blogger mypost.md > output.html"
echo "  cat mypost.md | md2blogger > output.html"
echo ""
echo "Run 'md2blogger -h' for more options."
