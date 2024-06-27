package main

import (
	"jwt-go/controllers"
	"jwt-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.Run()
}

//https://www.youtube.com/watch?v=ma7rUS_vW9M
