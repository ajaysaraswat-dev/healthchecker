package cmd

import (
	"fmt"
	"vmcli/internals/metrics"
)

func RunCPU() {
	cpuUsage,err := metrics.GetCPUUsage()
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	fmt.Printf("The cpu usage is %v",cpuUsage)
}