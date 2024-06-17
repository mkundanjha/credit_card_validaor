package main

import (
	"credit_card_validator/controller"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	app.GET("/", controller.ShowWelcomeMsg)
	app.GET("/verify/:cardnum", controller.VerifyCard)
}
