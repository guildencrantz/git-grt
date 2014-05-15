package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type listFormatter struct {
	subjSize int
	projSize int
}

func listCmd(args []string) int {
	if len(args) <= 0 {
		fmt.Println("")
		listCmdExec("outgoing")
		listCmdExec("incoming")
		listCmdExec("closed")
		return 0
	} else {
		fmt.Println("")
		for len(args) > 0 {
			listCmdExec(args[0][2:])
			args = args[1:]
		}
	}
	return 1
}

func listCmdExec(op string) {
	cmd := NewGrtCmd("GET", change_endpoint)
	switch op {
	case "outgoing":
		fmt.Println("Outgoing reviews (--outgoing):")
		cmd.getVars = map[string]string{
			"q": "is:open+owner:self",
		}

	case "incoming":
		fmt.Println("Incoming reviews (--incoming):")
		cmd.getVars = map[string]string{
			"q": "is:open+reviewer:self+-owner:self",
		}

	case "closed":
		fmt.Println("Incoming reviews (--closed):")
		cmd.getVars = map[string]string{
			"q": "is:closed+owner:self+limit:15&o=LABELS",
		}

	default:
		log.Fatal("--" + op + " is an invalid option.")
	}

	resp := cmd.Call()

	var list []ChangeInfo
	json.Unmarshal([]byte(resp), &list)

	lf := analyzeSizes(list)
	printDefaultList(list, lf)
	fmt.Println("")
}

func analyzeSizes(list []ChangeInfo) listFormatter {
	var lf listFormatter
	lf.subjSize = 8
	lf.projSize = 8

	for i := 0; i < len(list); i++ {
		if len(list[i].Subject) > lf.subjSize {
			lf.subjSize = len(list[i].Subject)
		}

		if len(list[i].Project) > lf.projSize {
			lf.projSize = len(list[i].Project)
		}
	}

	return lf
}

func printDefaultList(list []ChangeInfo, formatter listFormatter) {
	subjHeader := "Subject:"
	projHeader := "Project:"
	mergeHeader := "Merge Ready:"

	for i := len(subjHeader); i < formatter.subjSize; i++ {
		subjHeader += " "
	}

	for i := len(projHeader); i < formatter.projSize; i++ {
		projHeader += " "
	}

	fmt.Println(subjHeader + "\t" + projHeader + "\t" + mergeHeader)

	for i := 0; i < len(list); i++ {
		for j := len(list[i].Subject); j < formatter.subjSize; j++ {
			list[i].Subject += " "
		}

		for j := len(list[i].Project); j < formatter.projSize; j++ {
			list[i].Project += " "
		}

		var mergeable string
		if list[i].Mergeable > 0 {
			mergeable = "true"
		} else {
			mergeable = "false"
		}

		fmt.Println(list[i].Subject + "\t" + list[i].Project + "\t" + mergeable)
	}
}
