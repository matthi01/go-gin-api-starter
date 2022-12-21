package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matthi01/go-gin-api-starter/db"
	"github.com/matthi01/go-gin-api-starter/handler"
	"log"
	"os"
)

var port string
var listData *db.List

func main() {
	loadEnvData()
	instantiateData()

	gin.ForceConsoleColor()
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/items", handler.GetItems(listData))
	router.GET("/item/:id", handler.GetItem(listData))
	router.POST("/item", handler.CreateItem(listData))
	router.PUT("/item/:id", handler.UpdateItem(listData))
	router.DELETE("/item/:id", handler.DeleteItem(listData))

	router.Run(port)
}

func loadEnvData() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}
}

func instantiateData() {
	listData = db.New()
}
