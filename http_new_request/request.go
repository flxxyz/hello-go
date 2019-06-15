package http_new_request

import (
	"bytes"
	"fmt"
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
	fmt.Println("请求路径: ", url)

	reader := bytes.NewReader(body)
	req, _ := http.NewRequest(method, url, reader)
	if c.onUa {
		req.Header.Set("User-Agent", "HttpRequestTest/1.0 (+https://github.com/flxxyz/hello-go/HttpRequestTest)")
	}

	rw = httptest.NewRecorder()
	c.router.ServeHTTP(rw, req)

	return
}

func (c *core) GET(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) POST(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) PUT(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) OPTION(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) DELETE(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) HAED(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) PATCH(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) CONNECT(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}

func (c *core) TRACE(method string, url string, body []byte) (rw *httptest.ResponseRecorder) {
	return c.Request(method, url, body)
}
