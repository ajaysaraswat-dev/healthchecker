package metrics

import (
	"bufio"
	"fmt"
	"os"
)

type CPUStat struct {
User   float64
System float64
Idle   float64
Total  float64
Usage  float64
}

func Readcpu(){
	file,err := os.Open("/proc/stat")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err();err != nil {
		panic(err)
	}

}