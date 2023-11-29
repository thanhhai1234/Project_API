package Routes

import (
	"example.com/RestAPIgo/controllers"
	"github.com/gin-gonic/gin"
)

// ConfigTaskRouter configures routes related to TASKS in the application
func ConfigTaskRouter(r *gin.Engine) {
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}

// ConfigUserRouter configures routes related to USERS in the application
func ConfigUserRouter(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/users", controllers.FindUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)
}

// ConfigAuthRouter configures routes related to AUTHENTICATION in the application.
func ConfigAuthRouter(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
