package main

import (
    "fmt"
    "encoding/json"
)

type listChangeInfo struct {
	Kind       string
	Id         string
	Project    string
	Branch     string
	Change_id  string
	Subject    string
	Status     string
	Created    string
	Updated    string
	Mergeable  int
	Sortkey   string `json:"_sortkey"`
	Number    int    `json:"_number"`
	Owner      struct {
		Name string
	}
}

func listCmd(args []string) int {
	if len(args) <= 0 {
		listCmdDefault()
		return 0
	}
	return 1
}

func listCmdDefault() {
	cmd := NewGrtCmd("GET", "/a/changes/")
	resp := cmd.Call()

    var list []listChangeInfo
    json.Unmarshal([]byte(resp), &list)

    for i := 0; i < len(list); i++ {
        fmt.Println(list[i].Sortkey)
    }
}
