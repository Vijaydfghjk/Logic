package main

import (
	"fmt"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func processValue(value int) (int, error) {
	if value < 0 {
		panic(&CustomError{Message: "Negative value encountered"})
	}
	return value * 2, nil
}

func main() {
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered:", r) // r will storing the value of what panic return.
		}

	}()

	fmt.Println("Processing values...")
	value1, err1 := processValue(5)
	if err1 != nil {
		fmt.Println("Error 1:", err1)
	} else {
		fmt.Println("Value 1:", value1)
	}

	value2, err2 := processValue(-3)
	if err2 != nil {
		fmt.Println("Error 2:", err2)
	} else {
		fmt.Println("Value 2:", value2)
	}

	fmt.Println("Continuing program execution...")
}
