package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	//pa := &Address{"private", "Aartselaar", "Belgium"}
	//wa := &Address{"work", "Boom", "Belgium"}
	//vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	/*-------以上是写入 vcard.gob 的数据----*/
	var vc VCard
	file, err := os.Open("vcard.gob")
	if err != nil {
		log.Fatal("open file err", err)
	}
	defer file.Close()

	dec := gob.NewDecoder(file)
	decErr := dec.Decode(&vc)
	if decErr != nil {
		log.Fatal("decode error : ", decErr)
	}
	fmt.Println(vc)
}

//要传入一个指针,但定义vc的时候不定将vc定义成指针,要是值类型
//DecodeValue of unassignable value 没有价值的解压数据
