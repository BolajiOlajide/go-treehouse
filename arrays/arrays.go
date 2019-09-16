package arrays

import "fmt"

func main() {
	months := [3]string{"Apr", "May", "Jun"}

	var salesByMonth [3]float64
	salesByMonth[0] = 1712.23
	salesByMonth[1] = 2345.43
	salesByMonth[2] = 543.45

	// fmt.Println(months[0], salesByMonth[0])
	// fmt.Println(months[1], salesByMonth[1])
	// fmt.Println(months[2], salesByMonth[2])

	for i := 0; i < len(months); i++ {
		fmt.Println(months[i], salesByMonth[i])
	}

	for i, month := range months {
		fmt.Println(month, salesByMonth[i])
	}
}
