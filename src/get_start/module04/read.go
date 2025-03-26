package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// define name struct
type name struct {
	fname string
	lname string
}

func main() {

	// ask user to put file name
	var filename string
	fmt.Print("Please enter file name: ")
	fmt.Scan(&filename)

	// open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while open the file:", err)
		return
	}

	// close file
	defer file.Close()

	// create names slice to store first name and last name in the file
	var names []name

	// read file using scanner
	scanner := bufio.NewScanner(file)

	// loop through all line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// split first name and last name. assuming there will separate by 1 space
		texts := strings.SplitN(line, " ", 2)
		// check if the line doesn't have 2 name that is first name and last name
		if len(texts) != 2 {
			continue
		}

		// create a struct for the names, then store in the array
		n := name{
			fname: texts[0],
			lname: texts[1],
		}
		names = append(names, n)
	}

	// if have an error when reading file, display error then exit.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}

	// loop through the slice and print first name and last name
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}

}
