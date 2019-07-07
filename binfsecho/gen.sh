#!/bin/bash

set -e
set -u

go run ../cmd/binfs/main.go --pkg binfsecho_test testdata > binfsecho_gen_test.go
gofmt -s -w binfsecho_gen_test.go
