package main

import (
	"fmt"
	"time"

	stress "github.com/7Ethan/stress/stress"
)

func main() {
	task := &stress.Task{
		Duration:   10, //Continuous request for 10 seconds
		Concurrent: 10,
		ReportHandler: func(results []*stress.Result, totalTime time.Duration) {
			//Processing result reporting function.
			//If the function is passed in, the incoming function is used to process the report,
			//otherwise the default function is used to process the report.
		},
	}
	err := task.Run(&stress.RequestConfig{
		URLStr: "http://localhost:8080/api/test",
		Method: "GET",
	})
	if err != nil {
		fmt.Println(err)
	}
}
