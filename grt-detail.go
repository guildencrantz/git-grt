package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func detailCmd(id string) {
	details := getChangeDetails(id)

	fmt.Println("Project:")
	fmt.Println(details.Project)

	fmt.Println("Branch:")
	fmt.Println(details.Branch)
}

func getChangeDetails(id string) ChangeDetail {
	var change ChangeInfo
	changeType := []string{"outgoing", "incoming", "closed"}
	for i := 0; change.Id != id && i < len(changeType); i++ {
		_, changeList := getChanges(changeType[i])
		for j := 0; j < len(changeList); j++ {
			if strconv.Itoa(changeList[j].Number) == id {
				change = changeList[j]
			}
		}
	}

	endpoint := fmt.Sprintf(change_endpoint, change.Id)
	cmd := NewGrtCmd("GET", endpoint)
	resp := cmd.Call()

	var details ChangeDetail
	json.Unmarshal([]byte(resp), &details)

	return details
}
