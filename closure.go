package main

import "fmt"

func demo(name string) func(int) {

	return func(age int) {

		fmt.Println(name, age)
	}

}

func main() {

	a := demo("Vijay")
	a(24)
}

