package main

import (
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/api"
	"lanshan_homework/go1.19.2/go_homework/class_5_work/dao"
)

func main() {
	dao.InitRedis()
	dao.Start()
	api.InitRouter()
}
