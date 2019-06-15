package http_new_request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
)

type core struct {
	router http.Handler
	onUa   bool
}

func New(router http.Handler) *core {
	return &core{router, true}
}

func (c *core) Request(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	reader := bytes.NewReader(body)
	req, _ := http.NewRequest(method, url, reader)
	if c.onUa {
		req.Header.Set("User-Agent", "HttpRequestTest/1.0 (+https://github.com/flxxyz/hello-go/HttpRequestTest)")
	}

	rw = httptest.NewRecorder()
	c.router.ServeHTTP(rw, req)

	return
}

func (c *core) GET(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodGet, url, body)
}

func (c *core) POST(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodPost, url, body)
}

func (c *core) PUT(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodPut, url, body)
}

func (c *core) OPTIONS(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodOptions, url, body)
}

func (c *core) DELETE(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodDelete, url, body)
}

func (c *core) HAED(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodHead, url, body)
}

func (c *core) PATCH(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodPatch, url, body)
}

func (c *core) CONNECT(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodConnect, url, body)
}

func (c *core) TRACE(url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(http.MethodTrace, url, body)
}
