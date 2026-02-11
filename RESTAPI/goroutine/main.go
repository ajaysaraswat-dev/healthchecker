package main

import (
	"fmt"
	
)

func dosomething(){
	for i:=0;i<10;i++ {
		fmt.Println("process at ",i)
		
	}
}
func main(){
	//go dosomething()
	//dosomething()
	//time.Sleep(1*time.Second)
	
	testch := make(chan string)
	go func(){
		testch <- "value1"
		testch <- "value2"
		testch <- "value3"
		testch <- "value4"
		close(testch)
	}()
	for res:= range testch {
		fmt.Println(res)
	}
}