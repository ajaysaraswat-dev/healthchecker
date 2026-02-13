package cmd

import (
	"fmt"
	"vmcli/internals/metrics"
)

func RunMem() {
	mem, err := metrics.GetMemory()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Total Memory : %v\n", mem.Total)
	fmt.Printf("Free Memory : %v\n", mem.Total)
	fmt.Printf("idle Memory : %v\n", mem.Total)
}
