package main

import (
	"fmt"

	"github.com/treehouse-projects/go-intro/parts"
)

// all the structs in Parts have the Specs and Price method defined on them

type Parts interface {
	Specs() string
	Price() string
}

func Summary(part Part) string {
	return part.Specs() + "\n" + part.Price()
}

func main() {
	catalog := []Part{}
	catalog = append(catalog, parts.Monitor) // forgot to pass in the data to this struct
	catalog = append(catalog, parts.HardDrive)
	catalog = append(catalog, parts.Keyboard)

	for _, part := range catalog {
		Summary(part)
		fmt.Println("------------------------")
	}
}
