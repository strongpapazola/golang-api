package routes

import (
	"log"
	"net/http"
	user "rest_api_gin/controller/users"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		// c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusOK, gin.H{"status": "404", "msg": "Not Found"})
	})
	log.Println("[*] Starting Serve ...")
	log.Println("[+] Served ...")
	router.Run(":8000")
}
