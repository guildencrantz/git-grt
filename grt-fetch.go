package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func Fetch(args []string) {
	// The end point is: project~branch~change_id
	cmd := NewGrtCmd("GET", changes_endpoint)

	// I'm using the query instead of the direct changeset endpoint because then
	// we don't have to store the gerrit id (project~branch~change_id)
	cmd.getVars = make(map[string]string)
	cmd.getVars["q"] = getTrackingNumber()
	cmd.getVars["o"] = "ALL_REVISIONS&o=DOWNLOAD_COMMANDS"

	resp := cmd.Call()

	changeInfo := []ChangeInfo{}
	json.Unmarshal([]byte(resp), &changeInfo)

	// the info. We specify the gerrit changeset number so we should only get one
	// response, but we're using the query interface so we have to pull that
	// response out of a list.
	changeInfo[0].fetchRevisions()
}

func (changeInfo ChangeInfo) fetchRevisions() {
	for hash, revisionInfo := range changeInfo.Revisions {
		_, err := execCommand("log", "-n 0", hash)
		// The revision isn't being tracked if trying to get it's log
		// returns an error. If we already have the hash we assume it's tagged.
		if err != nil {
			// Currently the FetchInfo isn't being populated in 2.8.1, so we calculate it
			// The root of the refspec is the last two digits of the change number.
			number := strconv.Itoa(changeInfo.Number)
			root := number[len(number)-2:] // Why no negative indexes golang?
			rev := strconv.Itoa(revisionInfo.Number)
			ref := fmt.Sprintf("refs/changes/%s/%s/%s", root, number, rev)

			out, err := execGit("fetch", "gerrit", ref); if err != nil {
				log.Fatal(out)
			}

			tag := fmt.Sprintf("%s_%s", number, rev)
			out, err = execGit("tag", tag, hash); if err != nil {
				log.Fatal(out)
			}
		}
	}
}
