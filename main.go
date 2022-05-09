package main

import (
	"log"
	"net/http"

	"github.com/carlosm27/rester/models"
	"github.com/carlosm27/rester/services"

	"github.com/gin-gonic/gin"
)

func main() {

	models.Database()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/response", func(c *gin.Context) {
		response := FetchResponse()

		c.HTML(http.StatusOK, "response.tmpl", gin.H{
			"title": response,
		})
	})
	router.GET("/form", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.tmpl", gin.H{
			"title": "Form",
		})
	})

	router.GET("/requests", services.GetRequests)
	router.GET("/request/:id", services.GetRequest)

	router.POST("/request", services.PostRequest)

	log.Fatal(router.Run(":10000"))
}
