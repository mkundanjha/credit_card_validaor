package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	Routes(app)
	app.Run(":8080")

}
