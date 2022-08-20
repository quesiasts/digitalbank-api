package routes

import (
	"quesiasts/digitalbank-api.git/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:name", controllers.Introduction)
	r.Run()
}
