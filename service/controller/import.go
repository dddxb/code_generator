package controller

import (
	"code_generator/service/model"
	"fmt"
	"github.com/kataras/iris"
	"os"
)

//Import {from : D:/code/go/src,to:D:/code/go, name : autoCreate}
func Import(ctx iris.Context) {
	var imp model.Import
	err := ctx.ReadJSON(&imp)

	fmt.Println(imp.From, imp.To, imp.Name)

	if err == nil { //正确读取上下文
		//判断输入路径正确
		if pathExists(imp.From) {
			//输入项目文件名存在
			if pathExists(imp.From + "/" + imp.Name) { //可以移动项目文件夹及其子文件夹
				//移动文件夹
				err := os.Rename(imp.From+"/"+imp.Name, imp.To+"/"+imp.Name)
				if err != nil {
					fmt.Println(err)
					return
				}
				//正常移动，返回移动成功数据
				importTrue := model.Response{
					Success: true,
					Data:    nil,
					Error:   " ",
				}
				ctx.JSON(importTrue)
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