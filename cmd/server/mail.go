package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("starting")
	StartGin()
}

func StartGin() *gin.Engine {
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	port := "8090"
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
	return router

}
