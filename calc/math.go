package calc

import "fmt"

// returns sum of two integers
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}

// returns input string
func Echo(testo string, testo2 string) {
	fmt.Println("--->", testo, testo2)
}

