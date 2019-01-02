package controller

import (
	"code_generator/service/model"
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
)

//List {address : D:/code/go/src}
func List(ctx iris.Context) {
	var path model.List
	err := ctx.ReadJSON(&path)

	fmt.Println(path.Address)

	//http.body正确解析
	if err == nil {
		list, e := ioutil.ReadDir(path.Address)
		if e == nil { //输入路径正确
			var a [20]string
			for i, v := range list {
				a[i] = v.Name()
			}
			listTrue := model.ResponseList{
				Success: true,
				Data:    a,
				Error:   " ",
			}
			ctx.JSON(listTrue)
		} else { //输入路径不正确
			pathWrong := model.Response{
				Success: false,
				Data:    nil,
				Error:   "002 : Specified path was not found.",
			}
			ctx.JSON(pathWrong)
		}
	} else { //http.body没有正确解析
		praseWrong := model.Response{
			Success: false,
			Data:    nil,
			Error:   "001 : Service is down.",
		}
		ctx.JSON(praseWrong)
	}
}
