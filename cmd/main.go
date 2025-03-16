package main

import (
	"fmt"
	"vavilon-library/internal/hexagon"
)

func main() {
	//Get hex
	hex := hexagon.GetHexagon(12345)
	fmt.Printf("Hexagon %s contains %d books\n", hex.ID, len(hex.Books))

	book := hex.Books[678]
	buf := make([]byte, hexagon.BookLength)
	content := book.GenerateContent(buf)
	fmt.Printf("Book %s starts with %s\n", book.ID, string(content[:20]))
}