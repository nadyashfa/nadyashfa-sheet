package main

import "fmt"

func main() {
	var slicegue1 = []int{1, 2, 4, 7, 6}
	fmt.Println(slicegue1)

	var slicegue2 = append(slicegue1, 9, 1, 5)
	fmt.Println(slicegue2)
}
