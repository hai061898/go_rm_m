package main

import (

	"os"
	// "Restaurant Management/database"
	"github.com/gin-gonic/gin"

	routes "Restaurant Management/routes"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use()

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)
	routes.OrderItemRoutes(router)

	router.Run(":" + port)

}