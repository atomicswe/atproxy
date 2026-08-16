#!/bin/bash
# Build script for atproxy
# Author: @atomicswe (github.com/atomicswe)
# Url: github.com/atomicswe/atproxy

OPT_DIR="/opt/atproxy"
EXE_NAME="atproxy"

arch=""
if [[ $(uname -m) == *"arm"* ]]; then
	arch="arm64"
else
	arch="amd64"
fi

os=""
if [[ $(uname -s) == "Darwin" ]]; then
	os="darwin"
else
	os="linux"
fi

echo "arch: $arch"
echo "sys: $os"

echo "Building the executable with: os=$os and arch=$arch..."
GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $EXE_NAME ./cmd/atproxy
if ! [[ -e $EXE_NAME ]]; then
	echo "ERROR: The executable build failed."
	exit 1
fi
echo "Finished building the go executable."

mkdir -p $OPT_DIR
if ! [[ -d "$OPT_DIR" ]]; then
	echo "ERROR: Failed to create the opt directory for atproxy at: '$OPT_DIR'"
	exit 1
fi

mv $EXE_NAME $OPT_DIR
if ! [[ -e "$OPT_DIR/$EXE_NAME" ]]; then
	echo "ERROR: Failed to move the executable to '$OPT_DIR'."
	exit 1
fi
