package main

import ()

func listCmd(args []string) int {
	if len(args) <= 0 {
		listCmdDefault()
		return 0
	}
	return 1
}

func listCmdDefault() {
	cmd := NewGrtCmd("GET", "/a/changes/")
	cmd.Call()
}
