package main

import (
	"encoding/json"
)

const (
	change_endpoint = "/a/changes/"
)

type ChangeInfo struct {
	Kind       string
	Id         string
	Project    string
	Branch     string
	Topic      string
	ChangeId   string `json:"change_id"`
	Subject    string
	Status     string
	Created    string
	Updated    string
	Mergeable  int
	Insertions int
	Deletions  int
	Sortkey    string `json:"_sortkey"`
	Number     int    `json:"_number"`
	Owner      struct {
		Name string
	}
}

func (changeInfo ChangeInfo) String() string {
	ret, _ := json.Marshal(changeInfo)
	return string(ret)
}
