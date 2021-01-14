package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go-go/controllers"
)

func main()  {
	fmt.Println("START")
	controllers.StartWebServer()
}
