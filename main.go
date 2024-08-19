package main

import (
	"github.com/gin-gonic/gin"
	"golang-test/controllers"
	"golang-test/customers"
	"golang-test/limits"
	"golang-test/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "sunarto:sunarto123@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print("Database connection failed")
		log.Fatal(err)
	}

	err = db.AutoMigrate(customers.Customer{}, limits.Limit{}, model.Transaction{})
	if err != nil {
		log.Print("Database error when migrate")
		log.Fatal(err)
	}
	customerRepo := customers.NewRepository(db)
	customerService := customers.NewService(*customerRepo)
	customerHandler := controllers.NewCustomerHandler(customerService)

	limitRepo := limits.NewRepository(db)
	limitService := limits.NewService(*limitRepo)
	limitHandler := controllers.NewLimitHandler(limitService)

	r := gin.Default()

	r.POST("/customers", customerHandler.CustomerCreate)

	r.POST("/customers/:id/limit", limitHandler.LimitCreate)
	r.POST("transaction/:customerId")

	err = r.Run()

	if err != nil {
		log.Fatal("Server error")
	} // listen and serve on 0.0.0.0:8080
}
