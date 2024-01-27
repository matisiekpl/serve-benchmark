package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Record struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:root@tcp(mysql:3306)/benchmark?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Record{})
	if err != nil {
		return
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		record := Record{Name: "Hello World"}
		db.Create(&record)
		return c.JSON(http.StatusOK, record)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
