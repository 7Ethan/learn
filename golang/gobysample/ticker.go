package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(30 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	/*
		*$ go run ticker.go
			Tick at 2018-3-30 11:29:59.445736 +0800 CST m=+0.124074701
			Tick at 2018-3-30 11:29:59.5472635 +0800 CST m=+0.225602201
			Tick at 2018-3-30 11:29:59.6428382 +0800 CST m=+0.321176901
			Tick at 2018-3-30 11:29:59.7378911 +0800 CST m=+0.416229801
			Tick at 2018-3-30 11:29:59.8380687 +0800 CST m=+0.516407401
			Tick at 2018-3-30 11:29:59.9384321 +0800 CST m=+0.616770801
			Tick at 2018-3-30 11:30:00.0451632 +0800 CST m=+0.723501901
			Tick at 2018-3-30 11:30:00.1392279 +0800 CST m=+0.817566601
			Tick at 2018-3-30 11:30:00.2451477 +0800 CST m=+0.923486401
			Tick at 2018-3-30 11:30:00.3462849 +0800 CST m=+1.024623601
			Tick at 2018-3-30 11:30:00.44506 +0800 CST m=+1.123398701
			Tick at 2018-3-30 11:30:00.5377376 +0800 CST m=+1.216076301
			Tick at 2018-3-30 11:30:00.6390298 +0800 CST m=+1.317368501
			Tick at 2018-3-30 11:30:00.7455519 +0800 CST m=+1.423890601
			Tick at 2018-3-30 11:30:00.8395551 +0800 CST m=+1.517893801
			Tick at 2018-3-30 11:30:00.9449175 +0800 CST m=+1.623256201
			Ticker stopped
	*/
}
