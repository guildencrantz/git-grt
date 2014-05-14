package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal("I don't know what you want me to do.")
	}

	user, pass := getCreds()

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

func execCommand(command []string) (string, error) {
	cmd := exec.Command(command[0], command[0:]...)

	var usrout bytes.Buffer
	cmd.Stdout = &usrout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	out := strings.TrimSpace(usrout.String())

	return out, nil
}

func getCurrentBranch() (string, error) {
	return execCommand([]string{
		"git",
		"symbolic-ref",
		"--short",
		"HEAD",
	})
}

func setConfigValue(name string, value string) error {
	_, err := execCommand([]string{
		"git",
		"config",
		"set",
		name,
		value,
	})

	return err
}
