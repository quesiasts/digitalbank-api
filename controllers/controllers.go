package controllers

import "github.com/gin-gonic/gin"

func Introduction(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hi " + name + ", are you okay?",
	})
}
