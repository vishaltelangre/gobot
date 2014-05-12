package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// containsString is the function to lookup the slice of strings for an passd
// item; if the item is found in the slice, then it returns true, neither it
// returns false value.
func containsString(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// execCmdAndPrintResult is used to execute system commands and prints the
// result of the command from stdio to the interactive session.
func execCmdAndPrintResult(cmdWithArg string) {
	var cmdFields = strings.Fields(cmdWithArg)
	cmd := exec.Command(cmdFields[0], cmdFields[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Don't mind, something's broken!", "\n", err)
	}
	fmt.Println(strings.TrimSpace(out.String()))
}
