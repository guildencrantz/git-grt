package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Fail if not enough parameters
	if len(os.Args) < 2 {
		log.Fatal("I don't know what you want me to do.")
	}

	var client http.Client

	resp, err := client.Get("http://gerrit.dev.returnpath.net/a/changes/?q=status:open")
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "http://gerrit.dev.returnpath.net/a/changes/?q=status:open", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Do(req)

    user := getConfigValue("gerrit.user")
	fmt.Println(user)
    pass := getConfigValue("gerrit.pass")

	auth := GetAuthorization(user, pass, resp)
	digest := GetAuthString(auth, req.URL, req.Method, 1)
	fmt.Println(digest)
	req.Header.Add("Authorization", digest)

	resp, err = client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

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

func execCommand(command []string) string {
	cmd := exec.Command(command[0], command[0:]...)
    out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + string(out))
	}
	output := strings.TrimSpace(string(out))

	return output
}

func getCurrentBranch() string {
	return execCommand([]string{
		"git",
		"symbolic-ref",
		"--short",
		"HEAD",
	})
}

func getConfigValue(name string) string {
    name = "\"" + name + "\"" 
    val := execCommand([]string{
        "git",
        "config",
        "--get",
        name,
    })

    return val
}

func setConfigValue(name, value string) {
	execCommand([]string{
		"git",
		"config",
		"set",
		name,
		value,
	})
}
