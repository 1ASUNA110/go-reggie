package test

import (
	"fmt"
	"go-reggie/config"
	"go-reggie/internal/db"
	"go-reggie/internal/model/pojo"
	"os"
	"testing"
)

func TestDb(t *testing.T) {

	os.Chdir("..")

	config.InitConfig()

	// 测试数据库连接
	// 连接数据库
	db, err := db.InitDB()

	if err != nil {
		panic(err)
	}

	// 查询数据库
	var employee pojo.Employee
	db.First(&employee, 1)

	fmt.Println(employee)
}
