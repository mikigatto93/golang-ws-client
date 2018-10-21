// cmd_parser.go
package main

import (
	"fmt"
	"strings"
)

var commands = make(map[string]func(args []string))

func SetupCommands() {
	commands["pl"] = func(args []string) {
		fmt.Println(args[0])
		fmt.Println(args[1])
		fmt.Println(args[2])
	}
}

func ParseCommand(cmd string) {
	cmd = strings.Trim(cmd, "#")
	flags := strings.Split(cmd, "-")
	cmdParsed := flags[0]
	cmdParsed = strings.TrimSpace(cmdParsed)

	function, ok := commands[cmdParsed]
	if ok {
		if len(flags) != 1 {
			args := flags[1:]
			for i := 0; i < len(args); i++ {
				args[i] = strings.TrimSpace(args[i])
			}
			function(args)
		} else {
			var emptySlice []string
			function(emptySlice)
		}

	} else {
		fmt.Printf("Command %s does not exist\n", cmdParsed)
	}

}
