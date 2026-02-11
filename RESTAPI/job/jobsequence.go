package main

import "fmt"

func Worker(jobs chan string){
	for job := range jobs {
		fmt.Println("processing job ",job)
	}
}
func main(){
	jobs := make(chan string,5)

	go Worker(jobs)
	for k := 0 ;k< 5;k++ {
		job := fmt.Sprintf("job-%d",k)
		jobs <- job
	}
	close(jobs)
}