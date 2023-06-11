package main

import (
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
	//create user
	r.POST("/user/customer/add-customer", usersHandler.AddCustomer)
	//register admin
	r.POST("/user/admin/register", usersHandler.Register)
	//get all Customer
	r.GET("/user/customer/all-customer", usersHandler.GetAllCustomer)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
