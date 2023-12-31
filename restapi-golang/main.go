package main

import (
	"example.com/RestAPIgo/Routes"
	"example.com/RestAPIgo/models"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create a default Gin object to handle HTTP requests
	r := gin.Default()

	models.ConnectDatabase()

	Routes.ConfigTaskRouter(r)
	Routes.ConfigUserRouter(r)
	Routes.ConfigAuthRouter(r)

	err := r.Run()
	if err != nil {
		return
	}
}
