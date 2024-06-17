package controller

import (
	"credit_card_validator/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowWelcomeMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Card Validator. You provide us you card number, we validate.",
	})

}

func VerifyCard(c *gin.Context) {
	//get input cardnumber
	cardNumberStr := c.Param("cardnum")

	cardNum, _ := strconv.ParseInt(cardNumberStr, 10, 64)
	//run card number checker algo
	isValid := models.CheckSumVerifier(cardNum)

	//if true return success
	if isValid {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hurray! Your credit card is Valid!",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Opps! Your entred a wrong credit card number.",
		})
	}

}
