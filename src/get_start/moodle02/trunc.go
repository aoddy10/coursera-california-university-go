package main

import "fmt"

func main() {
	var input float64

	// asking user to enter number
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&input)

	// print the result
	fmt.Println("Truncated integer number:", int(input))
}
