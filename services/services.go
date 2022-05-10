package services

import (
	"fmt"

	"log"
	"net/http"

	"github.com/carlosm27/rester/httpMethods"
	"github.com/carlosm27/rester/models"
	"github.com/gin-gonic/gin"
)

type NewRequest struct {
	Uri string `json:"uri" binding:"required"`
}

func GetRequests(c *gin.Context) {

	var requests []models.Requests

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&requests).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requests)

}

func GetRequest(c *gin.Context) {

	var request models.Requests

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&request).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(request)
	c.JSON(http.StatusOK, request)

}

func PostRequest(c *gin.Context) {

	var newRequest NewRequest

	if err := c.ShouldBindJSON(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := httpMethods.GetUri(newRequest.Uri)

	if err != nil {
		log.Println(err)
	}

	request := models.Requests{Uri: newRequest.Uri, Response: response.Body()}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&request).Error; err != nil {
		log.Println(err.Error())
	}

	fmt.Println(request)

	c.JSON(http.StatusOK, request)
}
