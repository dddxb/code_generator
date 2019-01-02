package biz

import (
	"code_generator/service/model"
	"fmt"
	//_连库
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Create 插库，存储创建项目日志
func Create(create model.Create) error {
	fmt.Println(create)
	//创建orm引擎
	engine, err := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//日志打印SQL
	engine.ShowSQL(true)

	//插库
	a, e := engine.Insert(create)
	fmt.Println(a, e)
	return e
}
