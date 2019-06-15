package http_new_request_test

import (
	"github.com/flxxyz/hello-go/http_new_request"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestCore_Request(t *testing.T) {
	router := gin.Default()

	type route struct {
		Method   string
		Path     string
		Body     []byte
		Response string
	}

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodHead,
		http.MethodOptions,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodTrace,
	}

	routes := make([]route, 0)

	for _, method := range methods {
		routes = append(routes, route{
			Method:   method,
			Path:     "/test",
			Body:     nil,
			Response: "response test string!",
		})
	}

	for _, r := range routes {
		//这里需要闭包封装一下，不然r无法使用
		func(r route) {
			router.Handle(r.Method, r.Path, func(c *gin.Context) {
				data, _ := c.GetRawData()
				t.Logf("method: %s, data: %s", r.Method, data)

				c.String(http.StatusOK, r.Response)
			})
		}(r)
	}

	hnr := http_new_request.New(router)
	for _, r := range routes {
		response := hnr.Request(r.Method, r.Path, r.Body)

		t.Logf("source response: %s, request() response: %s", r.Response, response.Body.String())
	}
}
