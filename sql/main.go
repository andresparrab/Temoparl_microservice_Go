package main

import (
	"andresparrab/atest/app"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	app.ConnectToMySQL()
	app.GetInt()
	r := gin.Default()
	r.PUT("/", app.UpdateInt2)
	r.GET("/", app.GetInt2)
	fmt.Println(r)
	app.UpdateInt()

	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
