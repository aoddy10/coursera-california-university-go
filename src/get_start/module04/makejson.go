package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// create reader to get user input
	reader := bufio.NewReader(os.Stdin)

	// ask for user name
	fmt.Print("Please enter your name: ")
	userName, _ := reader.ReadString('\n')

	// ask for user address
	fmt.Print("Please enter your address: ")
	useAddress, _ := reader.ReadString('\n')

	// create map as assign user name and address
	user := map[string]string{
		"name":    userName[:len(userName)-1],
		"address": useAddress[:len(useAddress)-1],
	}

	// marshal user map to json
	jsonUser, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// display user detail in json format
	fmt.Println(string(jsonUser))
}
