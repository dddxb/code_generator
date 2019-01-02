package controller

import (
	"code_generator/service/model"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

func Test_Import(t *testing.T) {
	app := newAppImport()
	e := httptest.New(t, app)

	//路径正确,文件名存在，移动成功
	correctPath := model.Import{
		From: "D:/code/go/src",
		To:   "D:/code",
		Name: "test",
	}
	//路径正确，文件不存在，移动失败
	noexist := model.Import{
		From: "D:/code/go/src",
		To:   "D:/code/go",
		Name: "哈哈",
	}
	//路径错误
	wrongPath := model.Import{
		From: "D:/src",
		To:   "D:/code/go",
		Name: "autoCreate",
	}

	e.POST("/api/v1/import").WithJSON(correctPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":true,\"data\":null,\"error\":\" \"}")
	e.POST("/api/v1/import").WithJSON(noexist).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"004 : The project doesn't exists.\"}")
	e.POST("/api/v1/import").WithJSON(wrongPath).Expect().Status(httptest.StatusOK).Body().Equal("{\"success\":false,\"data\":null,\"error\":\"002 : Specified path was not found.\"}")
}

func newAppImport() *iris.Application {
	app := iris.New()
	app.Post("/api/v1/import", Import)
	return app
}
