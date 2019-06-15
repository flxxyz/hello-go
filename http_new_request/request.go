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

func (c *core) Request(method string, url string, body []byte) (rw http.ResponseWriter) {
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
