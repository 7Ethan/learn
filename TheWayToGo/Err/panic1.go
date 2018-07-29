package main

import (
	_ "fmt"
	"log"
	"time"
)

func main() {
	//fmt.Println("---------")
	//panic("testing panic!")
	//fmt.Println("end of program")
	//defer func() {
	//	if err := recover();err != nil {
	//		log.Printf("runtime error : %v",err)
	//	}
	//}()
	//for  {
	//	gg()
	//}
	for {
		protect(G)
	}
}

func G() {
	panic("testing")
	time.Sleep(2 * time.Second)
	//defer recover()
}

func protect(g func()) {
	defer func() {
		log.Println("Done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("Start")
	g() //   possible runtime-error
}
