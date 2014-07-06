package main

import (
	"fmt"
	"github.com/bjaskj/learnmethods/models"
)

func main() {
	user := &models.User{ FirstName: "Bjarte", LastName: "Skjørestad" }
	fmt.Println(models.Greet(user))

	location := &models.Location{ City: "Røyneberg", Region: "Rogaland" }
	fmt.Println(models.Greet(location))

	location2 := &models.Location{ City: "Røyneberg", Region: "Rogaland", Country: "Norway" }
	fmt.Println(models.Greet(location2))
}
