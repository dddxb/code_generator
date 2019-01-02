package controller

import (
	"code_generator/service/biz"
	"code_generator/service/model"
	"fmt"
	"os"

	"github.com/kataras/iris"
)

//Create {address : D:/code/go/src, name : autoCreate}
func Create(ctx iris.Context) {
	var path model.Create
	err := ctx.ReadJSON(&path)

	fmt.Println(path.Address, path.Name)

	//http.body正确解析
	if err == nil {
		//输入路径正确
		if pathExists(path.Address) {
			//输入项目文件名不存在
			if !pathExists(path.Address + "/" + path.Name) { //可以创建项目文件夹及其子文件夹
				//创建根目录和权限
				os.Mkdir(path.Address+"/"+path.Name, 0777)
				//创建dist文件夹及其子文件夹
				os.Mkdir(path.Address+"/"+path.Name+"/dist", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/dist"+"/linux", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/dist"+"/mac", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/dist"+"/win", 0777)
				//创建service文件夹及其子文件夹
				os.Mkdir(path.Address+"/"+path.Name+"/service", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/bin", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/biz", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/controller", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/dist", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/log", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/model", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/service"+"/utils", 0777)
				fileYaml, _ := os.Create(path.Address + "/" + path.Name + "/service" + "/config.yaml")
				fmt.Println(fileYaml)
				fileMain, _ := os.Create(path.Address + "/" + path.Name + "/service" + "/main.go")
				fmt.Println(fileMain)
				//创建webui文件夹及其子文件夹
				os.Mkdir(path.Address+"/"+path.Name+"/webui", 0777)
				os.Mkdir(path.Address+"/"+path.Name+"/webui"+"/dist", 0777)

				createTrue := model.Response{
					Success: true,
					Data:    nil,
					Error:   " ",
				}
				ctx.JSON(createTrue)
				//创建成功日志存入数据库
				e := biz.Create(path)
				fmt.Println(e)
			} else { //输入项目文件名已经存在
				fileExists := model.Response{
					Success: false,
					Data:    nil,
					Error:   "003 : The project already exists.",
				}
				ctx.JSON(fileExists)
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

//判断路径是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
