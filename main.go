package main

import (

	"os"

    "Restaurant Management/database"
	
	routes "Restaurant Management/routes"
	middleware "Restaurant Management/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + port)

}