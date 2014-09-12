package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 構造体で実行するコマンドとかディレクトリを定義しようかな
// type CommandInfo struct {
// 	execCommand   []string
// 	commandResult string
//  dir string
// }

var executableCommands map[string][]string = map[string][]string{
	"help":  {"cat " + getBaseDirPath() + "/aa.txt"},
	"pull":  {"cd " + getWebAppPath(), "git pull origin master"},
	"push":  {"cd " + getWebAppPath(), "git pull origin master"},
	"bench": {"cd " + getWebAppPath(), "isucon3 benchmark"},
}

func getBaseDirPath() string {
	var pwdResult string = runCommand("pwd")
	return strings.Trim(pwdResult, "\n")
}

func getWebAppPath() string {
	return "/home/isucon/webapp"
}

func main() {
	var subCommand string = ""
	if len(os.Args) == 2 {
		subCommand = os.Args[1]
	} else {
		subCommand = "help"
	}
	execute(subCommand)
}

func execute(subCommand string) {
	var (
		command []string = getCommand(subCommand)
		result  string
		i       int
	)
	// if subCommand NOT have key in executableCommands
	if command == nil {
		fmt.Println("Command not found !!")
		return
	}

	for i = 0; i < len(command); i++ {
		fmt.Printf(command[i])
		result = runCommand(command[i])
		fmt.Printf("%s", result)
	}
}

func getCommand(a string) []string {
	var (
		value []string
		ok    bool
	)

	value, ok = executableCommands[a]
	if ok {
		return value
	} else {
		return nil
	}
}

func runCommand(cmdStr string) string {
	var (
		parts   []string = strings.Fields(cmdStr)
		head    string   = parts[0]
		outPut  []byte
		err     error
		command *exec.Cmd
	)
	parts = parts[1:len(parts)]
	command = exec.Command(head, parts...)
	outPut, err = command.Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(outPut[:])
}
