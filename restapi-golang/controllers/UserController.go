package controllers

import (
	"example.com/RestAPIgo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserResponse struct {
	ID    uint          `json:"ID"`
	Name  string        `json:"Name"`
	Tasks []models.Task `json:"Tasks"`
}

// CreateUser creates a new user based on the JSON data sent from the client.
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Hash the password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	models.DB.Create(&user)

	responseData := gin.H{
		"id":   user.ID,
		"name": user.Name,
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// FindUser returns details of a user based on ID
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Preload("Tasks").Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	responseData := gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"tasks": user.Tasks,
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// FindUsers returns a list of all users in the database
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Preload("Tasks").Find(&users)

	// Create a new slice of User to store the data without passwords
	var usersWithoutPassword []UserResponse

	// store the new slice with the necessary data
	for _, user := range users {
		userWithoutPassword := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Tasks: user.Tasks,
		}
		usersWithoutPassword = append(usersWithoutPassword, userWithoutPassword)
	}

	c.JSON(http.StatusOK, gin.H{"data": usersWithoutPassword})
}

// DeleteUser deletes a user based on ID.
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
