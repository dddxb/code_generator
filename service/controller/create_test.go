package controller

import (
	"code_generator/service/model"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_Create(t *testing.T) {
	app := newAppCreate()
	e := httptest.New(t, app)

	//路径正确,文件名不存在，创建成功
	correctPath := model.Path{
		Address: "D:/code/go/src",
		Name:    "autoCreate",
	}
	//路径正确，文件名已经存在
	projectexist := model.Path{
		Address: "D:/code/go/src",
		Name:    "MKServer",
	}
	//路径错误
	wrongPath := model.Path{
		Address: "D:/src",
		Name:    "autoCreate",
	}

	e.POST("/api/v1/create").WithJSON(correctPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":true,\"data\":null,\"error\":\" \"}")
	e.POST("/api/v1/create").WithJSON(projectexist).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"003 : The project already exists.\"}")
	e.POST("/api/v1/create").WithJSON(wrongPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"002 : Specified path was not found.\"}")
}

func newAppCreate() *iris.Application {
	app := iris.New()
	app.Post("/api/v1/create", Create)
	return app
}
