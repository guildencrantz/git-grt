package main

import (
	"fmt"
)

const (
	endpoint = "changes"
)

func Branch() {
	id := createGerritBranch()
	addChangeIdToGitconfig(id)
}

func addChangeIdToGitconfig(id string) {
	branch := GetCurrentBranch()
	name := fmt.Sprintf("branch.%s.change-id", branch)
	SetConfigValue(name, id)
}

func createGerritBranch() string {
	return "bob"
}

/*
	changeInfo := struct {
		project string
		subject string
		branch  string
		topic   string
		status  string
	}{
		status: "DRAFT",
	}

	resp := struct {
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
	}{}

	// Here we need to post the payload to a url. Need to get the napping session from git-grt to do this.

	return resp.id
}
*/
