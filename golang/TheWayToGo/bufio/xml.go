package main

import (
	"encoding/xml"
	"fmt"
	_ "net"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.Decoder(inputReader)

	for t, err := p.Token(); err == nil; t, err := p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name  : %s \n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attibute is : %s %s", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of element")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is content : %v\n", content)
		default:

		}

	}
}
