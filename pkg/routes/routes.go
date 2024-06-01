package routes

import (
	"log"
	"time"

	"hex/cms/pkg/config"
	"hex/cms/pkg/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter () *gin.Engine {
	r := gin.Default()

	domain := config.GetEnv("DOMAIN")
	log.Printf("domain: %v", domain)

	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{domain},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")

	userApi := api.Group("/users")

	{
		userApi.POST("/signup", controllers.CreateUser)
	}

	return r
}