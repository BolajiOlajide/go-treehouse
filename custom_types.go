package main

import (
	"fmt"
	"strings"
)

// Title some random ish
type Title string

// Minutes minutes in a clock
type Minutes int

// FixCase method to fix the case for a title
func (t Title) FixCase() string {
	return strings.Title(string(t))
}

// Increment method for incrementing minutes
func (m *Minutes) Increment() {
	*m = (*m + 1) % 60
}

func main() {
	var name Title
	name = Title("bolaji")
	fmt.Println(name) // bolaji
	fmtName := name.FixCase()
	fmt.Println(fmtName) // Bolaji

	minutes := Minutes(58)
	for i := 0; i < 3; i++ {
		minutes.Increment() // note: no need to indicate the address of minutes (&minutes).Increment() will work
		fmt.Println(minutes)
	}
}
