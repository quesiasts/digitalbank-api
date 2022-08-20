package main

import (
	"quesiasts/digitalbank-api.git/database"
	"quesiasts/digitalbank-api.git/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
