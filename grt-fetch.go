package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func Fetch(args []string) {
	// The end point is: project~branch~change_id
	//endpoint := change_endpoint + "auth_policy_publisher~master~Ib0a47bffa7ddee956d980df3f799ebafb9ae1f2a"
	endpoint := change_endpoint + "gerrit-sandbox~master~Ib4d99019ceaf282a9e4b20322d840099611ca1cb"
	cmd := NewGrtCmd("GET", endpoint)
	cmd.rawForm = "o=ALL_REVISIONS&o=DOWNLOAD_COMMANDS"

	changeInfo := ChangeInfo{}
	resp := cmd.Call()

	println(resp)

	json.Unmarshal([]byte(resp), &changeInfo)

	// Currently the FetchInfo isn't being populated in 2.8.1, but we can calculate the info.
	changeInfo.fetchRevisions()
}

func (changeInfo ChangeInfo) fetchRevisions() {
	for hash, revisionInfo := range changeInfo.Revisions {
		_, err := execCommand("log", "-n 0", hash)
		// The revision isn't being tracked if trying to get it's log
		// returns an error.
		if err != nil {
			// The root of the refspec is the last two digits of the change number.
			number := strconv.Itoa(changeInfo.Number)
			root := number[len(number)-2:] // Why no negative indexes golang?
			ref := fmt.Sprintf("%s/%s/%s", root, number, strconv.Itoa(revisionInfo.Number))

			out, err := execGit("fetch", "gerrit", ref); if err != nil {
				log.Fatal(out)
			}
		}
	}
}

