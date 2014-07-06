package main

import (
	"fmt"
	"github.com/bjaskj/learnmethods/handlers"
	"github.com/bjaskj/learnmethods/models"
)

func doNamed() {
	user := &models.User{ FirstName: "Bjarte", LastName: "Skjørestad" }
	location := &models.Location{ City: "Røyneberg", Region: "Rogaland" }
	location2 := &models.Location{ City: "Røyneberg", Region: "Rogaland", Country: "Norway" }

	namedThings := []models.Namer { user, location, location2 }

	for _, v := range namedThings {
		fmt.Println(models.Greet(v))
	}
}

func doHandlers() {
	handler1 := &handlers.CompareHandler{ CompareWith: "bjarte" }
	handler2 := &handlers.CompareHandler{ CompareWith: "Bjarte" }

	handleThings := []handlers.Handler { handler1, handler2 }

	request := "Bjarte"

	for _, v := range handleThings {
		response, next := v.Process(request)

		fmt.Println(response)

		if !next {
			break
		}
	}
}

func main() {
	doNamed()
	doHandlers()
}
