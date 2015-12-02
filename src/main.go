package main

import (
	"cli"
	"fmt"
	"os"
)

var supportedCmds = map[string]func(string){
	"hub":    cli.Hub,
	"stream": cli.Stream,
}

func main() {
	args := os.Args
	argc := len(args)

	if argc > 2 {
		cmd := args[1]
		subCmd := args[2]
		if cliFunc, ok := supportedCmds[cmd]; ok {
			cliFunc(subCmd)
		} else {
			//parse flags, show help or version
			switch cmd {
			case "-v":
				cli.Version()
			case "-h":
				cli.Help()
			default:
				fmt.Println("Unknown cmd ", cmd)
			}
		}
	} else {
		fmt.Println("Use help or help [cmd1 [cmd2 [cmd3 ...]]] to see supported commands.")
	}
}
