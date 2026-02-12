package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//function for fetching the api

func fetchApi(ctx context.Context , url string, result chan<-string){
	req,err := http.NewRequestWithContext(ctx,"GET",url,nil)
	if err != nil {
		result <- fmt.Sprintf("error occured with context",err)
	}
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		result <- fmt.Sprintf("error occured with request",err)
	}
	result <- fmt.Sprintf("sucessfully conclude the api with status code %v",resp.StatusCode)

}

func main(){
	ctx,cancel := context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://instagram.com",

	}
	result := make(chan string)
	for _,url := range urls {
		go fetchApi(ctx,url,result)
	}
	for range urls {
		fmt.Println(<-result)
	}
	
	
}