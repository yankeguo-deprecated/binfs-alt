# binfs

[![Build Status](https://travis-ci.org/go-guoyk/binfs.svg?branch=master)](https://travis-ci.org/go-guoyk/binfs)
![GitHub](https://img.shields.io/github/license/go-guoyk/binfs.svg)

embedded filesystem for go binary

## Usage

### Get binfs

```bash
go get go.guoyk.net/binfs/cmd/binfs # the cli tool
go get go.guoyk.net/binfs           # the runtime package
```

### Generate File

```bash
binfs -pkg binfs_test public view > binfs.gen.go
```

This command read the content of directory `public` and `view`, output a `binfs.gen.go` file

The argument `pkg` specifies package name in `binfs.gen.go` file, default is `main`

## Use File

As long as binfs.gen.go is compiled with your source code, you can extract file with

```go
binfs.Open("/public/robots.txt")
```

You can also use `binfs.FileSystem()` or `binfs.Find("subdir").FileSystem()` to get a implementation of `http.FileSystem`

## Credits

Guo Y.K., MIT License
