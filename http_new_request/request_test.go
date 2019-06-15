package http_new_request_test

import (
	"github.com/flxxyz/hello-go/http_new_request"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCore_Request(t *testing.T) {
	router := gin.Default()

	_ = http_new_request.New(router)
}
