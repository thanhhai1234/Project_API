package controllers

import (
	"example.com/RestAPIgo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// CreateUser creates a new user.
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Hash mật khẩu trước khi lưu vào cơ sở dữ liệu
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
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

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
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
