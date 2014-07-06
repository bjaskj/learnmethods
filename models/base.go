package models

import "fmt"

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprint(n.Name())
}
