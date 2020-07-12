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
func square(x int) error {
	for i:=0; i<x: i++ {
		fmt.Printf("%v -> %v\n", i, i*i)
	}
	return nil
}

func main() {
	fmt.Println("Hello World, again")
	fmt.Println(PrintNumbers(10))
	err := square(10)
	if err != nil {
		fmt.Printf("Error occured %v\n", err)
	}
}
