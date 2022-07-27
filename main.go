package main

import (
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/whiterun/shopee-authorization/controller"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	
	r.GET("/auth", controller.Auth)
	r.GET("/token", controller.GetToken)
	
	r.Run(os.Getenv("HOST"))
}
