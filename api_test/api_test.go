package api_test

import (
	"fmt"
	"gopkg.in/h2non/baloo.v3"
	"net/http"
	"net/http/httptest"
	"tdd/common"
	"tdd/routes"
	"testing"
)

var test *baloo.Client

func init() {
	Server := httptest.NewServer(common.GetRouter(routes.GetRoutes()))
	BaseUrl := Server.URL
	test = baloo.New(BaseUrl)
}

func TestEndPoint(t *testing.T) {
	resp, _ := test.Get("/repos/test").
		Expect(t).
		Status(http.StatusOK).
		Send()
	fmt.Printf("%v", resp)
}
