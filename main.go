package main

import (
	"MINIPROJECT/module/admin"
	"MINIPROJECT/module/users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/crm?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	usersHandler := users.DefaultRequestHandler(db)
	adminHandler := admin.DefaultRequestHandler(db)
	//create user
	r.POST("/users", usersHandler.Create)
	//register admin
	r.POST("/admin/register",adminHandler.Register)
	//get all users
	r.GET("/all-users", usersHandler.GetAll)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
