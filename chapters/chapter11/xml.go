package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	input := "<Person desc='Hello' encoding='UTF-8'><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	reader := strings.NewReader(input)
	decoder := xml.NewDecoder(reader)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:

		}
	}
}
