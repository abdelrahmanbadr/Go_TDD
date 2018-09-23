package controllers_test
//
//import (
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"tdd/controllers"
//	"tdd/models"
//	"testing"
//)
//
//type ReposTestClient struct {
//	Repos []models.Repo
//	Err   error
//}
//func (c ReposTestClient) Get(user string) ([]models.Repo, error) {
//
//	return c.Repos, c.Err
//}
//
//func TestGetReposHandler(t *testing.T) {
//	assert := assert.New(t)
//
//	tests := []struct {
//		description        string
//		reposClient        *ReposTestClient
//		url                string
//		expectedStatusCode int
//		expectedBody       string
//	}{
//		//{
//		//	description:        "missing argument user",
//		//	reposClient:        &ReposTestClient{},
//		//	url:                "/repos",
//		//	expectedStatusCode: 400,
//		//	expectedBody:       "MISSING_ARG_USER\n",
//		//},
//
//		//{
//		//	description: "no repos found",
//		//	reposClient: &ReposTestClient{
//		//		Repos: []models.Repo{},
//		//		Err:   nil,
//		//	},
//		//	url:                "/repos/fakeuser",
//		//	expectedStatusCode: 200,
//		//	expectedBody:       `[]`,
//		//},
//		{
//			description: "succesfull query",
//			reposClient: &ReposTestClient{
//				Repos: []models.Repo{
//					models.Repo{Name: "test", Description: "a test"},
//				},
//				Err: nil,
//			},
//			url:                "/repos/fakeuser",
//			expectedStatusCode: 200,
//			expectedBody:       `[{"name":"test","description":"a test"}]`,
//		},
//		// TODO not all cases are covered
//	}
//
//	for _, tc := range tests {
//		app := &controllers.App{Repos: tc.reposClient}
//
//		req, err := http.NewRequest("GET", tc.url, nil)
//		assert.NoError(err)
//
//		w := httptest.NewRecorder()
//		app.GetReposHandler(w, req)
//
//		assert.Equal(tc.expectedStatusCode, w.Code, tc.description)
//		assert.Equal(tc.expectedBody, w.Body.String(), tc.description)
//	}
//}
////func TestGetReposHandler(t *testing.T) {
////	assert := assert.New(t)
////	app := controllers.App{ReposClientTest{}}
////	req, err := http.NewRequest("GET", "/repos/11", nil)
////	assert.NoError(err)
////	w := httptest.NewRecorder()
////	app.GetReposHandler(w, req)
////	assert.Equal(400, w.Code, "MISSING_ARG_USER")
////
////}
