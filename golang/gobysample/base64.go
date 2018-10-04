package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	stdEncode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(stdEncode)
	stdDecode,_ := base64.StdEncoding.DecodeString(stdEncode)
	fmt.Println(string(stdDecode))

	urlEncode := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEncode)
	urlDecode,_ := base64.URLEncoding.DecodeString(urlEncode)
	fmt.Println(string(urlDecode))

	/*
	YWJjMTIzIT8kKiYoKSctPUB+
	abc123!?$*&()'-=@~
	YWJjMTIzIT8kKiYoKSctPUB-
	abc123!?$*&()'-=@~
	*/
}
