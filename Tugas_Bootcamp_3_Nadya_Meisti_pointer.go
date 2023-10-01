package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	var rect = Rectangle{
		Width:  15,
		Height: 2,
	}

	var area = rect.Area()

	rect.Width = 20

	fmt.Printf("Area: %d\n", area)
	fmt.Printf("Updated Width: %d\n", rect.Width)

}
