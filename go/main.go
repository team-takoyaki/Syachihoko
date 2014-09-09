package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const WEBAPP_PATH = "/home/isucon/webapp"

var executableCommands map[string]string = map[string]string{
	"info":  "golang",
	"pull":  "git pull origin master",
	"push":  "git pull origin master",
	"bench": "isucon3 benchmark",
}

func getCommand(a string) string {
	value, ok := executableCommands[a]
	if ok {
		return value
	} else {
		return ""
	}
}

func main() {

	var subCommand = ""
	if len(os.Args) == 2 {
		subCommand = os.Args[1]
	} else {
		subCommand = "info"
	}

	execute(subCommand)
}

func runCommand(s string) {
	fmt.Println(s)
	var cmdStr string = "/usr/bin/which ls"
	parts := strings.Fields(cmdStr)
	head := parts[0]
	parts = parts[1:len(parts)]
	cmd, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cmd)
}

func execute(subCommand string) {
	// executable_commands
	var command string = getCommand(subCommand)
	if command == "" {
		fmt.Println("Command not found !!")
		return
	}
	runCommand(command)
}
