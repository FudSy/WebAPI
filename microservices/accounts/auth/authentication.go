package auth

import (
	"net/http"

	"github.com/FudSy/WebApi/initializers"
	"github.com/FudSy/WebApi/models"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var userdata models.UserData
	if err := c.ShouldBindJSON(&userdata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if err := initializers.DB.Create(&userdata).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "user created"})
}