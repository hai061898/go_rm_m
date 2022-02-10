package main

import {
	"os"
	"Restaurant Management/database"
	"github.com/gin-gonic/gin"
}

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	
}