package auth

import (
	"net/http"

	"github.com/FudSy/WebApi/initializers"
	"github.com/FudSy/WebApi/models"
	"github.com/FudSy/WebApi/models/roles"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var userdata models.User
	if err := c.ShouldBindJSON(&userdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	userdata.Role = roles.User
	
	var count int64
	initializers.DB.Where("user_name = ?", userdata.UserName).Find(&models.User{}).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid username"})
		return
	}
	var hash []byte
	hash, _ = bcrypt.GenerateFromPassword([]byte(userdata.Password), bcrypt.DefaultCost)
	userdata.Password = string(hash)
	if err := initializers.DB.Create(&userdata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "user created"})
}