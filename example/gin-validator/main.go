package main

import (
	"gin-validator/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	exampleController := &controllers.ExampleController{}
	router.GET("/ping", exampleController.Ping)
	_ = router.Run(":9999")
}
