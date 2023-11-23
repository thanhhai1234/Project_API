package main

import (
	"example.com/RestAPIgo/Routes"
	"example.com/RestAPIgo/models"
	"github.com/gin-gonic/gin"
)

func main() {
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
