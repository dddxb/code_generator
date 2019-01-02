package controller

import (
	"code_generator/service/model"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_Delete(t *testing.T) {
	app := newAppDelete()
	e := httptest.New(t, app)

	//路径正确,文件名存在，删除成功
	correctPath := model.Path{
		Address: "D:/code/go/src",
		Name:    "autoCreate",
	}
	//路径正确，文件名不存在，删除失败
	projectexist := model.Path{
		Address: "D:/code/go/src",
		Name:    "哈哈",
	}
	//路径错误，删除失败
	wrongPath := model.Path{
		Address: "D:/src",
		Name:    "autoCreate",
	}

	e.POST("/api/v1/delete").WithJSON(correctPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":true,\"data\":null,\"error\":\" \"}")
	e.POST("/api/v1/delete").WithJSON(projectexist).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"004 : The project doesn't exists.\"}")
	e.POST("/api/v1/delete").WithJSON(wrongPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"002 : Specified path was not found.\"}")
}

func newAppDelete() *iris.Application {
	app := iris.New()
	app.Post("/api/v1/delete", Delete)
	return app
}
