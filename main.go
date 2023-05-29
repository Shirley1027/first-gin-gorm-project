package main

import (
	_ "first_gin-gorm_proj/api/database"
	"first_gin-gorm_proj/api/router"
)

func main() {
	//defer orm.Db.Close()
	router := router.InitRouter()
	router.Run()
}
