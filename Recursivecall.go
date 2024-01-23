package main

import "fmt"

func Recursivecall(arr []int, b int) []int {

	if b == 0 {

		return arr
	}

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}

	Recursivecall(arr, b-1)
	return arr
}

func main() {

	nums := []int{3, 2, 1, 0}
	k := Recursivecall(nums, len(nums))

	fmt.Println(k)
}
