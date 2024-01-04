package main

import (
    "github.com/machinebox/graphql"
	//"log"
	"context"
	"fmt"
	"github.com/lonmarsDev/bpo-golang-grahpql/pkg/log"
//	"time"
)





func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		//time.Sleep(time.Second)
		task1(j)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	const numJobs = 100000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 1000 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 1000; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}






func task1(j int) {
    // create a client (safe to share across requests)
	client := graphql.NewClient("http://ec2-18-223-123-140.us-east-2.compute.amazonaws.com:6968/graphql")
	//client := graphql.NewClient("http://localhost:6969/graphql")

	// make a request
	req := graphql.NewRequest(`
		query AllPipeline {
			allPipeline(filter:{
			  limit:20
			  
			}){
			  totalCount
			  results{
				orderNumber
				address	
				authorId
				authorName
			  }
			}
		  }   
	`)

	// set any variables
	req.Var("key", "value")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQHR1cmJvLmNvbSIsInVzZXJJZCI6IjVkYzEzMThmNDc4YjZlNzZkODYxMWExNCIsIm5hbWUiOiJhZG1pbiB0dXJibyIsInJvbGUiOlsiQURNSU4iXSwiZXhwIjoxNTg3MTE0OTIxfQ.aFk0cayjnlV2qmgvdJaPL3_--c4oblbRamOraFuiz8g")
	// run it and capture the response
	var respData interface{}
	ctx :=  context.TODO()

	if err := client.Run(ctx, req, &respData); err != nil {
		log.Error("error %s", err)
	}
	log.Debug("job # %d", j)
	//fmt.Println("@debug output %s", respData)

}	