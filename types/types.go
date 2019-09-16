package types

import "fmt"

// Minutes the number of minutes
type Minutes int

// Hours the number of hours
type Hours int

// Weight the weight of an item
type Weight float64

// Title the title of an article
type Title string

// Answer response to a question
type Answer bool

func types() {
	minutes := Minutes(37)
	hours := Hours(2)
	weight := Weight(945.7)
	name := Title("The Matrix")
	answer := Answer(true)
	fmt.Println(minutes, hours, weight, name, answer)
}
