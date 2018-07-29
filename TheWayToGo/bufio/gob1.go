package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.Encoder(&network)
	dec := gob.Decoder(&network)

	err := enc.Encode(P{1, 2, 3, "haha"})
	if err != nil {
		log.Fatal("encoder error : ", err)
	}

	var q Q
	err = dec.Decode(q)
	if err != nil {
		log.Fatal("decode error : ", err)
	}
	fmt.Printf(" %q : {%d,%d}\n", q.Name, q.X, q.Y)
}
