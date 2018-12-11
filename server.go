package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thanhson1085/api-starter/config"
	"github.com/thanhson1085/api-starter/controllers"
)

func main() {
	env := flag.String("env", "default", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		userGroup := api.Group("/user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/", user.Get)
		}
	}

	// Start and run the server

	config := config.GetConfig()
	router.Run(":" + config.GetString("port"))
}
