package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// create empty slice
	slice := make([]int, 0, 3)

	for {
		// asking user to put integer or x to quit
		fmt.Print("Enter an integer or 'X' to quit: ")
		var strInput string
		fmt.Scanln(&strInput)
		strInput = strings.TrimSpace(strInput)

		// check if the user wants to quit
		if strings.ToLower(strInput) == "x" {
			fmt.Println("Exiting program.")
			break
		}

		// convert string input to be integer
		intNumber, err := strconv.Atoi(strInput)
		if err != nil {
			// if fails, ask user to put number again
			fmt.Println("Invalid input.")
			continue
		}

		// add number to slice
		slice = append(slice, intNumber)

		// sort the slice
		sort.Ints(slice)

		// print the slice
		fmt.Println("Sorted slice:", slice)
	}
}
