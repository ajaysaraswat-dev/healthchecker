package main

import (
	"fmt"
	"os"
	"vmcli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE : VMCLI[ CPU | MEMORY ]")
		return
	}
	switch os.Args[1] {
	case "cpu":
		cmd.RunCPU()
	case "memory":
		cmd.RunMem()
	default:
		fmt.Println("Unknown Command")
	}

}
