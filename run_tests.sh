#! /bin/sh

export GOPATH=$(cd ../../../..; pwd)

go test github.com/csizsek/prism/tests
