package main

import (
	"log"

	"github.com/carlosm27/rester/models"
	//"github.com/go-resty/resty/v2"
)

func FetchResponse() *models.Requests {

	var response models.Requests

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.Find(&response)

	return &response
}
