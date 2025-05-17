package lexer_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ProImpact/first-ast/lexer"
)

func TestParse_String(t *testing.T) {
	f, err := os.Open("../parseText.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	parser, err := lexer.Init(f)
	if err != nil {
		t.Fatal(err)
	}
	for tok := parser.Next(); tok.Type != lexer.EOF; tok = parser.Next() {
		fmt.Printf("%v\n", tok)
	}
}
