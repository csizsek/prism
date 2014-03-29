#! /bin/sh

find $(dirname $0) -name "*.go" | xargs gofmt -w
