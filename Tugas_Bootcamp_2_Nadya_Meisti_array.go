package main

import "fmt"

func main() {
	numbers := []int{8, 2, 4, 7, 6}

	fmt.Println(numbers)

	bubbleSort(numbers)

	fmt.Println(numbers)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
