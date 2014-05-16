package main

import (
	"encoding/json"
)

const (
	change_endpoint = "/a/changes/"
)

type ChangeInfo struct {
	Branch     string
	ChangeId   string `json:"change_id"`
	Created    string
	CurrentRevision string `json:"current_revision"`
	Deletions  int
	Id         string
	Insertions int
	Kind       string
	Mergeable  int
	Number     int `json:"_number"`
	Owner      struct {
		Name string
	}
	Project    string
	Reviewed   string
	Revisions  map[string]RevisionInfo
	Sortkey    string `json:"_sortkey"`
	Status     string
	Subject    string
	Topic      string
	Updated    string
}

func (changeInfo ChangeInfo) String() string {
	ret, _ := json.Marshal(changeInfo)
	return string(ret)
}

type FetchInfo struct {
	Url string
	Ref string
}

func (fetchInfo FetchInfo) String() string {
	ret, _ := json.Marshal(fetchInfo)
	return string(ret)
}

type RevisionInfo struct {
	Draft            string
	HasDraftComments bool `json:"has_draft_comments"`
	Number           int `json:"_number"`
	Fetch            map[string]FetchInfo
}

func (revisionInfo RevisionInfo) String() string {
	ret, _ := json.Marshal(revisionInfo)
	return string(ret)
}


