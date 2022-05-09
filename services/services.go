package services

import (
	"encoding/json"
	"fmt"
	"strings"

	//"io/ioutil"
	//"strconv"

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
		c.JSON(http.StatusNotFound, gin.H{"error": "URI not found"})
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

	var resp map[string]interface{}
	err = json.Unmarshal(response.Body(), &resp)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	res := strings.ReplaceAll(string(jsonResp), "\"", "")

	request := models.Requests{Uri: newRequest.Uri, Response: res}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&request); err != nil {
		log.Println(err)
	}

	fmt.Println(request)

	c.JSON(http.StatusOK, request)
}
