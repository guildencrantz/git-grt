package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	endpoint = "changes"
)

type ChangeInfo struct {
	project string
	subject string
	branch  string
	topic   string
	status  string
}

func Branch() {
	var changeInfo ChangeInfo
	reader := bufio.NewReader(os.Stdin)

	proposed_project := "CALCULATED PROJECT NAME"
	fmt.Printf("project [%s]: ", proposed_project)
	line, err := reader.ReadString('\n')
	changeInfo.project = strings.TrimSpace(line)
	if err != nil {
		panic("There was an issue reading the project name: " + err.Error())
	}
	if len(changeInfo.project) == 0 {
		changeInfo.project = proposed_project
	}

	fmt.Printf(changeInfo.project)
}
