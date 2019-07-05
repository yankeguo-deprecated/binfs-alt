#!/bin/bash

set -e
set -u

PKG=binfs_test go run cmd/binfs/main.go testdata1 testdata2 > binfs_gen_test.go

