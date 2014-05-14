package main

import (
	"bytes"
	"fmt"
	"github.com/davemeehan/Neo4j-GO"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getCreds() (string, string) {
	cmd := exec.Command("git", "config", "--get", "gerrit.user")
	var usrout bytes.Buffer
	cmd.Stdout = &usrout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	user := strings.TrimSpace(usrout.String())

	cmd = exec.Command("git", "config", "--get", "gerrit.pass")
	var pwdout bytes.Buffer
	cmd.Stdout = &pwdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	pass := strings.TrimSpace(pwdout.String())

	return user, pass
}

func main() {
	// Fail if not enough parameters
	if len(os.Args) < 2 {
		log.Printf("I don't know what you want me to do.")
		return
	}

	user, pass := getCreds()

	neo, err := neo4j.NewNeo4j("http://gerrit.dev.returnpath.net/a/changes", user, pass)
	if err != nil {
		// log.Printf("%v\n", err)
		return
	}

	node := map[string]string{
		"q": "status:open",
	}

	data, _ := neo.CreateNode(node)
	fmt.Printf("\nNode data: %s\n", data)
	switch os.Args[1] {
	case "branch":
		Branch()
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

func _get_current_branch() string {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")

	var usrout bytes.Buffer
	cmd.Stdout = &usrout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	out := strings.TrimSpace(usrout.String())

	return out
}
