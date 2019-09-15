package main

import (
	"fmt"
	"math"
	"os"
	"reflect"

	"github.com/golang/example/stringutil"
	"github.com/treehouse-projects/go-intro/welcome"
)

func main() {
	fmt.Println(welcome.English)
	fmt.Println(welcome.Japanese)
	reversedString := stringutil.Reverse("hello")
	fmt.Println(reversedString)
	fmt.Println()
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.4))

	fmt.Println()
	var age int
	age = 23

	var name = "Bolaji"

	gender := "Male"
	var a, b int
	a, b = 3, 4

	fmt.Println(age, name, gender, a, b)

	var c = "c"
	{
		var d = "d"
		{
			var e = "e"
			fmt.Println(c, d, e)
		}
		fmt.Println(c, d)
	}
	// fmt.Println(c, d) - will return an error saying d is undefined
	// because d is in a scope of its own
	fmt.Println(c)
	fmt.Println(square(9))
	fmt.Println(add(2.3, 5.7))
	fmt.Println(subtract(5, 3))
	sqrt, err := squareRoot(9)
	fmt.Println(sqrt, err)
	fmt.Println()
	// sqrt2, err2 := squareRoot(-9)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// fmt.Println(sqrt2)
	fmt.Println(checkFile("existent.txt"))
	// fmt.Println(checkFile("nonexistent.txt")) // catch the error
}

func square(number int) int {
	return number * number
}

// named return value
func add(a, b float64) (sum float64) {
	return a + b
}

// bare return
func subtract(a, b float64) (difference float64) {
	difference = a - b
	return
}

func checkFile(file string) (int64, error) {
	fileInfo, err := os.Stat(file)
	return fileInfo.Size(), err
}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("Can't calculate square root of a -ve number")
	}
	return math.Sqrt(x), nil
}
