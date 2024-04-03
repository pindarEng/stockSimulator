package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pindarEng/stockSimulator/controllers"
	"github.com/pindarEng/stockSimulator/initializers"
	"github.com/pindarEng/stockSimulator/middleware"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/validate", func(c *gin.Context) {
		c.HTML(http.StatusOK, "validate.html", nil)
	})
	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", nil)
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)

	fmt.Println("server listening on port:", os.Getenv("PORT"))
	r.Run()
}
