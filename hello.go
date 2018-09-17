package main

import (
	"fmt"
)

func PrintNumbers(m int) {
	for i := 0; i < m; i++ {
		fmt.Printf("Number: %d\n", i)
	}
}
func main() {
	fmt.Println("Hello World, again")
	PrintNumbers(10)
}
