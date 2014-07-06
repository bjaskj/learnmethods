package models

import "fmt"


type Location struct {
	City, Country, Region string
}

func (x *Location) Name() string {
	if len(x.Country) > 0 {
		return fmt.Sprintf("%s, %s in %s", x.City, x.Region, x.Country)
	}

	return fmt.Sprintf("%s, %s", x.City, x.Region)
}
