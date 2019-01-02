package controller

import (
	"code_generator/service/model"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_List(t *testing.T) {
	app := newAppList()
	e := httptest.New(t, app)

	//路径正确
	correctPath := model.List{
		Address: "D:/code/go",
	}

	//路径错误
	wrongPath := model.List{
		Address: "D:/src",
	}

	e.GET("/api/v1/lists").WithJSON(correctPath).Expect().Body().Equal("{\"success\":true,\"data\":[\"bin\",\"pkg\",\"src\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\"],\"error\":\" \"}")
	e.GET("/api/v1/lists").WithJSON(wrongPath).Expect().Body().Equal("{\"success\":false,\"data\":null,\"error\":\"002 : Specified path was not found.\"}")
}

//.Status(httptest.StatusOK)
func newAppList() *iris.Application {
	app := iris.New()
	app.Get("/api/v1/lists", List)
	return app
}
