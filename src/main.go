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
			fmt.Println("Unknown cmd ", cmd)
		}
	} else {
		if argc > 1 {
			//parse flags, show help or version
			option := args[1]
			switch option {
			case "-v":
				cli.Version()
			case "-h":
				cli.Help()
			default:
				fmt.Println("Unknow option", option)
			}
		} else {
			fmt.Println("Use -h to see supported commands.")
		}
	}
}
