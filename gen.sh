#!/bin/bash

set -e
set -u

go run cmd/binfs/main.go --pkg binfs_test testdata1 testdata2 > binfs_gen_test.go
gofmt -s -w binfs_gen_test.go
