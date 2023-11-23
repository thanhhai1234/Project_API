package Routes

import (
	"example.com/RestAPIgo/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigTaskRouter(r *gin.Engine) {
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}

func ConfigUserRouter(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/users", controllers.FindUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)
}

func ConfigAuthRouter(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
