package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func execCommand(args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func execGit(args ...string) (string, error) {
	command := append([]string{"git"}, args...)
	return execCommand(command...)
}

func execCommandLogFatal(command []string) string {
	cmd := exec.Command(command[0], command[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + string(out))
	}
	output := strings.TrimSpace(string(out))

	return output
}

func GetCurrentBranch() string {
	return execCommandLogFatal([]string{
		"git",
		"symbolic-ref",
		"--short",
		"HEAD",
	})
}

func GetConfigValue(name string) string {
	val, err := execGit("config", name)
	if err != nil {
		log.Fatal("Unable to retrieve git config '", name, "': Please set it and try again.")
	}
	return strings.TrimSpace(val)
}

func SetConfigValue(name, value string) {
	execCommandLogFatal([]string{
		"git",
		"config",
		"--add",
		name,
		value,
	})
}
