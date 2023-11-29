package controllers

import (
	"example.com/RestAPIgo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserResponse struct {
	ID       uint   `json:"ID"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

// CreateUser creates a new user based on the JSON data sent from the client.
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
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
		"id":       user.ID,
		"name":     user.Name,
		"password": user.Password,
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

// FindUser returns details of a user based on ID
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// FindUsers returns a list of all users in the database
func FindUsers(c *gin.Context) {
	var users []models.User
	var userResponses []UserResponse

	models.DB.Find(&users)

	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Password: user.Password,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponses})
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
