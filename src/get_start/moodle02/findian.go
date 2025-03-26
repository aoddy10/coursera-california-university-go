package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// ask user to enter string
	fmt.Print("Please enter some word: ")
	input_string, _ := reader.ReadString('\n')
	input_string = strings.TrimSpace(input_string)

	// convert to lower case
	input_string = strings.ToLower(input_string)

	// print "Fount!" if the string starts with 'i' ,ends with the character ‘n’, and contains the character ‘a’
	if strings.HasPrefix(input_string, "i") && strings.HasSuffix(input_string, "n") && strings.Contains(input_string, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
