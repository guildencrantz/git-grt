package main

import (
    "fmt"
    "encoding/json"
)

func listCmd(args []string) int {
	if len(args) <= 0 {
		listCmdDefault()
		return 0
	}
	return 1
}

func listCmdDefault() {
	cmd := NewGrtCmd("GET", change_endpoint)
	resp := cmd.Call()

    var list []ChangeInfo
    json.Unmarshal([]byte(resp), &list)

    for i := 0; i < len(list); i++ {
        fmt.Println(list[i].Sortkey)
    }
}
