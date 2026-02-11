package main

import (
	"fmt"
	"sync"
	
)


func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("worker", id ,"started")
	fmt.Println("worker", id ,"finished")
}

func Worker2(t int,wg *sync.WaitGroup){
	defer wg.Done()
	for k:=0;k<100;k++ {
		//run this loop
	}
	fmt.Println("run the func.......worker2 ",t)
}

func main(){
	var wg sync.WaitGroup
	for i:=1;i<3;i++ {
		wg.Add(1)
		go worker(i , &wg)
	}
	
	//wait for all workers for finish
	wg.Add(1)
	go Worker2(100,&wg)
	
	fmt.Println("all workers completed")
	wg.Wait()
}