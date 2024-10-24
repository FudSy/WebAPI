package main

import (
	"github.com/FudSy/WebApi/initializers"
	"github.com/FudSy/WebApi/microservices/accounts/auth"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.InitializeEnvFile()
	initializers.InitializeDB()
}

func main() {
	r := gin.Default()
	r.POST("/api/Authentication/SignUp", auth.SignUp)
	r.Run()
}