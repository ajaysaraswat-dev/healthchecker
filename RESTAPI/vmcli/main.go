package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"strconv"
	"strings"
	"time"
)
func ReadCputime()(total float64 ,idle float64){
file, err := os.Open("/proc/stat")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstline := scanner.Text()
		fmt.Println(firstline)
		fields := strings.Fields(firstline)
		values := [] float64{}
		for _,val := range fields[1:] {
			// fmt.Printf("%v index value is %v", i, val)
			intval,err := strconv.ParseFloat(val,64)
			if err !=nil{
				panic(err)
			}
			values = append(values,intval)
			

		}
		
		for _,v := range values[:8] {
			total += v
		}
		idle = values[3]
	}
	

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return total, idle
}

func readMemuse(){
	file,err := os.Open("/proc/meminfo")
	if err!=nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
}

func main(){
	prevtotal,previdle := ReadCputime()
	

	time.Sleep(1*time.Second)

	newtotal,newidle := ReadCputime()

	delta_idle := newidle - previdle
	delta_total := newtotal - prevtotal

	
	cpuUsage := (1-(delta_idle/delta_total)) * 100;
	fmt.Println("cpu uage is ",cpuUsage)
}
