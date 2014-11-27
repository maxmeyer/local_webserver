#!/bin/sh

set -e

[[ -z "$GOPATH" ]] && echo "Please set GOPATH" >&2

echo "Install go package manager"

go get github.com/mattn/gom

echo "Installation packages locally"
gom install

echo "Installation packages globally"
go get 'github.com/gorilla/handlers'
go get 'github.com/skratchdot/open-golang/open'
go get 'github.com/smartystreets/goconvey'
go get 'gopkg.in/alecthomas/kingpin.v1'
go get 'github.com/pwaller/goupx'
