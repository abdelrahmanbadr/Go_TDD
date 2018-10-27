package repository
//
//import (
//	"github.com/magiconair/properties/assert"
//	"gopkg.in/h2non/baloo.v3"
//	"net/http"
//	"net/http/httptest"
//	"tdd/common"
//	"tdd/controllers"
//	"tdd/models"
//	"tdd/routes"
//	"testing"
//)
//var test *baloo.Client
//func init() {
//	Server := httptest.NewServer(common.GetRouter(routes.GetRoutes()))
//	BaseUrl := Server.URL
//	test = baloo.New(BaseUrl)
//}
//
//type ReposClientTest struct {
//	Repos []models.Repo
//	Err   error
//}
//
//func (c ReposClientTest) Get(user string) ([]models.Repo, error) {
//
//	return c.Repos, c.Err
//}
//func TestGetReposHandler(t *testing.T) {
//	client := ReposClientTest{[]models.Repo{
//		models.Repo{Name: "test", Description: "a test"},
//	}, nil}
//	app := controllers.RespoController{client}
//	req, err := http.NewRequest("GET", "/repos", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(app.GetAllReposHandler)
//	handler.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	assert.Equal(t,rr.Body.String(), `[{"Name":"test","Description":"a test"}]`)
//
//}


