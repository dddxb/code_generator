package biz

import (
	"code_generator/service/model"
	"fmt"
	//_连库
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Delete 插库，存储删除日志
func Delete(delete model.Delete) error {
	fmt.Println(delete)
	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//插库
	a, e := engine.Insert(delete)
	fmt.Println(a, e)
	return e
}
