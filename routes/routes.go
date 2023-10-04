package routes

import "github.com/gin-gonic/gin"

func HandleRequests() {
	r := gin.Default()
	r.GET("/users", controllers.UsersController)
	r.Run()
}
