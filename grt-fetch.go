package main

import (
	"encoding/json"
)

func Fetch(args []string) {
	// The end point is: project~branch~change_id
	endpoint := change_endpoint + "auth_policy_publisher~master~Ib0a47bffa7ddee956d980df3f799ebafb9ae1f2a"
	cmd := NewGrtCmd("GET", endpoint)
	cmd.rawForm = "o=ALL_REVISIONS"

	changeInfo := ChangeInfo{}
	resp := cmd.Call()
	println(resp)

	json.Unmarshal([]byte(resp), &changeInfo)
	println(changeInfo.String())

}

