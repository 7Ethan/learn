package main

import (
	"fmt"
	"log"

	// "net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], " server")
		os.Exit(1)
	}

	serverAddr := os.Args[1]

	// client, err := rpc.Dial("tcp", serverAddr+":1234")
	client, err := jsonrpc.Dial("tcp", serverAddr+":1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Math.Mutiply", args, &reply)
	if err != nil {
		log.Fatal("Math Error : ", err)
	}
	fmt.Printf("Math Mutiply : %d * %d  = %d \n", args.A, args.B, reply)

	var quo Quotient
	err = client.Call("Math.Divide", args, &quo)
	if err != nil {
		log.Fatal("Math Error : ", err)
	}
	fmt.Printf("Math Divide : %d * %d  = %d reminder %d\n", args.A, args.B, quo.Quo, quo.Rem)

}
