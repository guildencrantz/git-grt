package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func execCommand(command []string) string {
	cmd := exec.Command(command[0], command[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + string(out))
	}
	output := strings.TrimSpace(string(out))

	return output
}

func GetCurrentBranch() string {
	return execCommand([]string{
		"git",
		"symbolic-ref",
		"--short",
		"HEAD",
	})
}

func GetConfigValue(name string) string {
	val := execCommand([]string{
		"git",
		"config",
		name,
	})
	return val
}

func SetConfigValue(name, value string) {
	execCommand([]string{
		"git",
		"config",
		"--add",
		name,
		value,
	})
}
