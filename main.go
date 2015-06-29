package main

import (
	"fmt"
	"log"
	"os"

	"github.com/letterj/pub/pages"
)

// Main module
func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}
	filename := os.Args[1]

	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)
}
