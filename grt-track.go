package main

import (
	"fmt"
	"log"
	"strings"
)

func Track(args []string) {
	switch {
	case len(args) == 0:
		number := getTrackingNumber()
		if number != "" {
			fmt.Println("You are currently tracking gerrit change: ", number)
		} else {
			fmt.Println("You aren't currently tracking a gerrit change.")
		}
	case len(args) == 1:
		setTrackingNumber(args[0])
		fmt.Println("You are now tracking gerrit change: ", args[0])
	default:
		// Need to write help for this.
		log.Fatal("Too many arguments for Track")
	}
}

func getGerritNumberKey() string {
	return fmt.Sprintf("branch.%s.gerritnumber", GetCurrentBranch())
}

func getTrackingNumber() string {
	val, err := execGit("config", getGerritNumberKey())
	if err != nil {
		return ""
	}
	return strings.TrimSpace(val)
}

func setTrackingNumber(number string) {
	out, err := execGit("config", "--add", getGerritNumberKey(), number)
	if err != nil {
		log.Fatal(out)
	}
}
