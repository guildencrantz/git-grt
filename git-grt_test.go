package main

import (
	"testing"
)

func TestExecCommand(t *testing.T) {
	command := "certainlyThisCommandDoesNotExistOnAnySystemEver12343251aoeusntnhaoeulhg"
	_, err := execCommand([]string{command})
	if err != nil {
		t.Error("Error expected when trying to execute '%s'", command)
	}
}
