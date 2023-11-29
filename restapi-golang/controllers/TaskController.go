package controllers

import (
	"example.com/RestAPIgo/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateTaskInput struct {
	Title     string `json:"Title" binding:"required"`
	Completed string `json:"Completed" binding:"required"`
	CreatedAt string `json:"CreatedAt" binding:"required"`
}

type UpdateTaskInput struct {
	Title     string `json:"Title"`
	Completed string `json:"Completed"`
}

func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// DetermineCompletedStatus determines the completion status of a task based on its creation time and the current time.
func DetermineCompletedStatus(createdAt time.Time) string {
	var nowLocal = time.Now()

	nowLocalString := nowLocal.Format("2006-01-02")
	createdAtString := createdAt.Format("2006-01-02")

	if createdAtString < nowLocalString {
		return "OverDue"
	} else if createdAtString > nowLocalString {
		return "Open"
	} else {
		return "Due"
	}
}

// CreateTask creates a new task using the provided input, sets its completion status based on the creation time,
// associates it with the user specified in the token, and returns the task's details.
func CreateTask(c *gin.Context) {
	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdAt, err := time.Parse("2006-01-02", input.CreatedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	// Get UserID from token
	userID, exists := getUserIDFromToken(c)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	var user models.User

	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	task := models.Task{
		Title:     input.Title,
		Completed: DetermineCompletedStatus(createdAt),
		CreatedAt: createdAt,
		UserID:    user.ID,
	}
	models.DB.Create(&task)

	responseData := gin.H{
		"id":        task.ID,
		"title":     task.Title,
		"completed": task.Completed,
		"createdAt": task.CreatedAt.Format("2006-01-02"),
		"userID":    user.ID,
	}
	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// getUserIDFromToken extracts the user's ID from the authentication token.
func getUserIDFromToken(c *gin.Context) (uint, bool) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		return 0, false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})

	if err != nil || !token.Valid {
		return 0, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		return 0, false
	}

	return uint(userID), true
}

// FindTask returns details of a task based on ID
func FindTask(c *gin.Context) {
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask Update tasks based on id
func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask delete task based on id
func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
