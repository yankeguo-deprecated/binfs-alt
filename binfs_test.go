package binfs_test

import (
	"bytes"
	"go.guoyk.net/binfs"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFind(t *testing.T) {
	n := binfs.Find("testdata2", "testdata3", "hello3.txt")
	if !bytes.Equal(n.Chunk.Data, []byte("hello\n")) {
		t.Fatal("not equal")
	}
}

func TestFileSystem(t *testing.T) {
	fs := binfs.FileSystem()
	var f http.File
	var err error
	if f, err = fs.Open("testdata2/hello2.txt"); err != nil {
		t.Fatal(err)
	}
	f.Close()

	var buf []byte

	if buf, err = ioutil.ReadAll(f); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf, []byte("hello\n")) {
		t.Fatal("not equal")
	}
}

func TestOpen(t *testing.T) {
	var f http.File
	var err error
	if f, err = binfs.Open("/testdata2/hello2.txt"); err != nil {
		t.Fatal(err)
	}
	f.Close()

	var buf []byte

	if buf, err = ioutil.ReadAll(f); err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf, []byte("hello\n")) {
		t.Fatal("not equal")
	}
}

func TestWalk(t *testing.T) {
	binfs.Walk(func(n *binfs.Node) {
		t.Logf(n.Name, n.Path)
	})
}

func TestFile_Readdir(t *testing.T) {
	f, _ := binfs.Open("/testdata2")
	files, _ := f.Readdir(-1)
	t.Log(files)
}
