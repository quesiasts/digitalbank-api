package routes

import (
	"quesiasts/digitalbank-api.git/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:name", controllers.Introduction)
	r.GET("/accounts", controllers.ListAllAccounts)
	r.GET("/accounts/:account_id/balance", controllers.SearchForBalance)
	r.POST("/accounts", controllers.CreateAccount)
	r.Run()
}
