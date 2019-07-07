package binfsecho_test

import (
	"github.com/labstack/echo/v4"
	"go.guoyk.net/binfs/binfsecho"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatic(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/file1.txt", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	m := binfsecho.StaticWithConfig(binfsecho.StaticConfig{
		Root:  "testdata",
		Index: []string{"index.txt"},
	})
	h := m(handler)
	h(c)

	if rec.Code != http.StatusOK {
		t.Fatal("not 200")
	}
	if rec.Body.String() != "hello1\n" {
		t.Fatal("not hello1")
	}
	if !strings.Contains(rec.Header().Get("Content-Type"), "text/plain") {
		t.Fatal("not text/plain: ", rec.Header().Get("Content-Type"))
	}

	req = httptest.NewRequest(http.MethodGet, "/dir1/file2.txt", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	handler = func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	m = binfsecho.StaticWithConfig(binfsecho.StaticConfig{
		Root:  "testdata",
		Index: []string{"index.txt"},
	})
	h = m(handler)
	h(c)

	if rec.Code != http.StatusOK {
		t.Fatal("not 200")
	}
	if rec.Body.String() != "hello2\n" {
		t.Fatal("not hello2")
	}
	if !strings.Contains(rec.Header().Get("Content-Type"), "text/plain") {
		t.Fatal("not text/plain: ", rec.Header().Get("Content-Type"))
	}

	req = httptest.NewRequest(http.MethodGet, "/dir3/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	handler = func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}
	m = binfsecho.StaticWithConfig(binfsecho.StaticConfig{
		Root:  "testdata",
		Index: []string{"index.txt"},
	})
	h = m(handler)
	h(c)

	if rec.Code != http.StatusOK {
		t.Fatal("not 200")
	}
	if rec.Body.String() != "hello3\n" {
		t.Fatal("not hello3")
	}
	if !strings.Contains(rec.Header().Get("Content-Type"), "text/plain") {
		t.Fatal("not text/plain: ", rec.Header().Get("Content-Type"))
	}
}
