package routes

import (
	"log"
	"net/http"
	user "users_update/controller/users"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
//		api.GET("/users", user.GetAllUser)
//		api.POST("/users", user.CreateUser)
//		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
//		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	log.Println("[*] Starting Serve ...")
	router.Run(":8000")
	log.Println("[+] Served ...")
}
