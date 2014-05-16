package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
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
	headers, list := getChanges(op)
	lf := analyzeSizes(list)
	printDefaultList(headers, list, lf)
	fmt.Println("")
}

func getChanges(op string) (string, []ChangeInfo) {
	cmd := NewGrtCmd("GET", changes_endpoint)
	var headers string
	switch op {
	case "outgoing":
		headers = "Outgoing reviews (--outgoing):"
		cmd.getVars = map[string]string{
			"q": "is:open+owner:self",
		}

	case "incoming":
		headers = "Incoming reviews (--incoming):"
		cmd.getVars = map[string]string{
			"q": "is:open+reviewer:self+-owner:self",
		}

	case "closed":
		headers = "Incoming reviews (--closed):"
		cmd.getVars = map[string]string{
			"q": "is:closed+owner:self+limit:15&o=LABELS",
		}

	default:
		log.Fatal("--" + op + " is an invalid option.")
	}

	resp := cmd.Call()

	var list []ChangeInfo
	json.Unmarshal([]byte(resp), &list)

	return headers, list
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

func printDefaultList(headers string, list []ChangeInfo, formatter listFormatter) {
	idHeader := "ID:  "
	subjHeader := "Subject:"
	projHeader := "Project:"
	mergeHeader := "Review:  "

	fmt.Println(headers)

	for i := len(subjHeader); i < formatter.subjSize; i++ {
		subjHeader += " "
	}

	for i := len(projHeader); i < formatter.projSize; i++ {
		projHeader += " "
	}

	fmt.Println(idHeader + "\t" + subjHeader + "\t" + projHeader + "\t" + mergeHeader)

	for i := 0; i < len(list); i++ {
		for j := len(list[i].Subject); j < formatter.subjSize; j++ {
			list[i].Subject += " "
		}

		for j := len(list[i].Project); j < formatter.projSize; j++ {
			list[i].Project += " "
		}

		cr := getCodeReview(strconv.Itoa(list[i].Number))

		fmt.Println(strconv.Itoa(list[i].Number) + "\t" + list[i].Subject + "\t" + list[i].Project + "\t" + strconv.Itoa(cr))
	}
}

func getCodeReview(id string) int {
	crh := 0
	crl := 0

	details := getChangeDetails(id)
	for i := 0; i < len(details.Labels.CodeReview.All); i++ {
		if details.Labels.CodeReview.All[i].Value < crl {
			crl = details.Labels.CodeReview.All[i].Value
		}

		if crh < details.Labels.CodeReview.All[i].Value {
			crh = details.Labels.CodeReview.All[i].Value
		}
	}

	var cr int
	if math.Abs(float64(crl)) >= float64(crh) {
		cr = crl
	} else {
		cr = crh
	}

	return cr
}
