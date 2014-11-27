#!/bin/sh

set -e

[[ -z "$GOPATH" ]] && echo "Please set GOPATH" >&2

echo "Install go package manager"
go get github.com/mattn/gom

gom install
