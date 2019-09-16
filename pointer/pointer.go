package pointer

import "fmt"

func main() {
	myNumber := 2.6
	halve(&myNumber)
	fmt.Println("myNumber in 'main':", myNumber)
}

func halve(number *float64) {
	*number = *number / 2
	fmt.Println("*number in 'halve':", *number)
}
