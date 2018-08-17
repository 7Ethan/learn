package main

import (
	"github.com/goinaction/code/chapter7/patterns/work"
	"log"
)

func main() {
}

func server(workchan <-chan *work.Worker) {
	for work := range workchan {
		go safelyDo(work)
	}
}

func safelyDo(worr *work.Worker) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Woek failed with %s in %v", err, worr)
		}
	}()
	do(worr)
}

func do(work work.Worker) {

}
