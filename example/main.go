package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	exampleController := &ExampleController{}
	router.GET("/ping", exampleController.Ping)
	_ = router.Run(":9999")
}
