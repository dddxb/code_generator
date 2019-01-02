package main

import (
    "code_generator/controller"
	stdContext "context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"time"
)


func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations,iris.WithoutInterruptHandler)
}


func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())

	//
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})


	//创建项目文件夹
	//POST {address，name}
	// response
	app.Post("/api/v1/create", controller.Create)

	//导入项目文件夹
	//POST {from, to, name}
	//response
	app.Post("/api/v1/import", controller.Import)

	//删除项目文件夹
	//POST {address, name}
	//response
	app.Post("/api/v1/delete", controller.Delete)

	//获取项目文件夹列表
	//GET {address}
	//response
	app.Get("/api/v1/lists",controller.List )

	return app
}