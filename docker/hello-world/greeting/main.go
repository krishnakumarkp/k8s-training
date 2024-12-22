package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("Usage: greeting <message>")
		return
	}

	// Join the arguments into a single message
	message := strings.Join(os.Args[1:], " ")

	// Append "World" to the message
	messageWithWorld := fmt.Sprintf("%s World!", message)

	// Print the updated message
	fmt.Println(messageWithWorld)
}
