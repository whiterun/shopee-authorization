package main

import (
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/whiterun/shopee-authorization/controller"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	
	r.GET("/build", controller.Build)
	
	r.Run(os.Getenv("HOST"))
}
