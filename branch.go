package main

import (
	"encoding/json"
	"fmt"
)

const (
	endpoint = "changes"
)

type ChangeInfo struct {
	kind string
	id string
	project string
	branch string
	topic string
	change_id string
	subject string
	status string
	created string
	updated string
	mergeable int
	insertions int
	deletions int
	_sortkey string
	_number int
	owner struct {
		name string
	}
}

func Branch() {
	id := createGerritBranch()
	addChangeIdToGitconfig(id)
}

func addChangeIdToGitconfig(id string) {
	branch := GetCurrentBranch()
	name := fmt.Sprintf("branch.%s.change-id", branch)
	SetConfigValue(name, id)
}


func createGerritBranch() *ChangeInfo {

	/*
	defaultRequestInfo = ChangeInfo{
		status: "DRAFT",
	}

	// Pass defaultRequestInfo to rest client
	*/

	resp := `{
		"kind": "gerritcodereview#change",
		"id": "myProject~master~I8473b95934b5732ac55d26311a706c9c2bde9941",
		"project": "myProject",
		"branch": "master",
		"topic": "create-change-in-browser",
		"change_id": "I8473b95934b5732ac55d26311a706c9c2bde9941",
		"subject": "Let's support 100% Gerrit workflow direct in browser",
		"status": "DRAFT",
		"created": "2014-05-05 07:15:44.639000000",
		"updated": "2014-05-05 07:15:44.639000000",
		"mergeable": true,
		"insertions": 0,
		"deletions": 0,
		"_sortkey": "002cbc25000004e5",
		"_number": 4711,
		"owner": {
			"name": "John Doe"
		}
	}`

	changeInfo := ChangeInfo{}
	json.Unmarshal([]byte(resp), &changeInfo)

	return &changeInfo
}

