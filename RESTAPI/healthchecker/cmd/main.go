package main

import (
	"context"
	
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func checkHealth(){
	const url = "http://localhost:9000"
	client := &http.Client {
		Timeout: 3 * time.Second,
	}
	ctx,cancel := context.WithTimeout(context.Background(),2 * time.Second)
	defer cancel()

	req,err := http.NewRequestWithContext(ctx,"GET",url,nil)
	if err != nil {
		log.Panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	if resp.StatusCode == 200 {
		log.Println("Healthy Service, status code ",resp.StatusCode)
	}else {
		log.Println("Unhealthy Service, Status Code ",resp.StatusCode)
	}


	defer resp.Body.Close()
}
func main(){
	router := gin.Default()

	//make the ticker for checking health api at every 10s
	ticker := time.NewTicker(10 * time.Second)

	//sop the ticker when program exits
	defer ticker.Stop()
    //infinite loop for checking the service (/health api)
	
	go func(){
		for {
		<- ticker.C // it stops for 10 sec
		checkHealth()
	}
	}()
	
	router.Run(":9010")

	
	
}