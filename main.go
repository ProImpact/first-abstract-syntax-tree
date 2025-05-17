package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProImpact/first-ast/lexer"
)

func main() {
	f, err := os.Open("./parseText.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	parser, err := lexer.Init(f)
	if err != nil {
		log.Fatal(err)
	}
	parser.Parse()
	fmt.Println(parser)
}
