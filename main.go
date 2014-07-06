package main

import (
	"fmt"
	"github.com/bjaskj/learnmethods/models"
)

func main() {
	user := &models.User{ FirstName: "Bjarte", LastName: "Skjørestad" }
	location := &models.Location{ City: "Røyneberg", Region: "Rogaland" }
	location2 := &models.Location{ City: "Røyneberg", Region: "Rogaland", Country: "Norway" }

	namedThings := []models.Namer { user, location, location2 }

	for _, v := range namedThings {
		fmt.Println(models.Greet(v))
	}
}
