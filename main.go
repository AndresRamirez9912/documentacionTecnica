package main

import (
	"log"
	"techDocumentation/src/models"
)

func main() {
	_, err := models.NewDocumentCreator()
	if err != nil {
		log.Fatal("Error openning the DB connection", err)
	}

}
