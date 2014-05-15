package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Fail if not enough parameters
	if len(os.Args) < 2 {
		log.Fatal("I don't know what you want me to do.")
	}

	switch os.Args[1] {
	case "list":
		listCmd(os.Args[2:])
	case "fetch":
		Fetch(os.Args[2:])
	case "push":
		fallthrough
	case "pull":
		// Check for changeset in params yet
		// If not, do push and get changeset back
		//
		fallthrough
	default:
		fmt.Printf("You have chosen to \"%s\".\n", os.Args[1])
		break
	}
}
