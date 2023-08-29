package main

import "fmt"

type Data struct {
	val1 int
	val2 int
	val3 int
}

func fii() {
	data := []Data{
		{2, 5, 7},
		{2, 3, 5},
		{2, -1, 8},
	}

	for _, val := range data {
		var sum int
		var expected int

		sum = val.val1 + val.val2
		expected = val.val3

		if sum != expected {

			fmt.Printf(" %d + %d = %d, expected %d", val.val1, val.val2, sum, expected)
		}

	}

}

func main() {
	fii()
}
