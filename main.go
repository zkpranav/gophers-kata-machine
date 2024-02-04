package main

import "fmt"

func main() {
	arr := [5]int {
		0, 1, 2, 3, 4,
	}

	ptr := &arr
	fmt.Printf("%T\n", ptr)
}
