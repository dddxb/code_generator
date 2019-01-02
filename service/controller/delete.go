package controller

import (
	"code_generator/service/biz"
	"code_generator/service/model"
	"fmt"
	"os"

	"github.com/kataras/iris"
)

//Delete {address : D:/code/go/src, name : autoCreate}
func Delete(ctx iris.Context) {
	var path model.Delete
	err := ctx.ReadJSON(&path)

	fmt.Println(path.Address, path.Name)

	//http.body正确解析
	if err == nil {
		//输入路径正确
		if pathExists(path.Address) {
			//输入项目文件名存在
			if pathExists(path.Address + "/" + path.Name) { //可以删除项目文件夹及其子文件夹
				//删除多级目录
				os.RemoveAll(path.Address + "/" + path.Name)
				deleteTrue := model.Response{
					Success: true,
					Data:    nil,
					Error:   " ",
				}
				ctx.JSON(deleteTrue)
				//删除成功日志存入数据库
				a := biz.Delete(path)
				fmt.Println(a)
			} else { //输入项目文件名不存在
				noProject := model.Response{
					Success: false,
					Data:    nil,
					Error:   "004 : The project doesn't exists.",
				}
				ctx.JSON(noProject)
			}
		} else { //输入路径不正确
			pathWrong := model.Response{
				Success: false,
				Data:    nil,
				Error:   "002 : Specified path was not found.",
			}
			ctx.JSON(pathWrong)
		}
	} else { //http.body 没有正确解析
		praseWrong := model.Response{
			Success: false,
			Data:    nil,
			Error:   "001 : Service is down.",
		}
		ctx.JSON(praseWrong)
	}
}
