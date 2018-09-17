package main

import (
	"fmt"
)

// PrintNumbers prints numbers from 0 to given integer.
func PrintNumbers(m int) int {
	c := 0
	for i := 0; i < m; i++ {
		fmt.Printf("Number: %d\n", i)
		c++
	}
	return c
}
func main() {
	fmt.Println("Hello World, again")
	fmt.Println(PrintNumbers(10))
}
