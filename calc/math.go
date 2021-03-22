// my test package
package calc

import (
	"errors"
	"fmt"
)

// returns sum of two integers
func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("provide more than 2 members"), sum
	}

	for _, num := range numbers {
		sum = sum + num
	}

	return nil, sum
}

// returns sum of two integers
func echo(testo string) {
	fmt.Println("--->", testo)
}
