package config

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	params gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method

	c.Request.Header.Set("Content-Type", "application/json")

	c.Params = params

	c.Request.Body = body

	c.Request.URL.RawQuery = u.Encode()

}
