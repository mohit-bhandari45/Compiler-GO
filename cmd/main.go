/* This is the entry point for the project. */

package main

import (
	"fmt"

	"github.com/mohit-bhandari45/Compiler-GO.git/internal/lexer"
)

func main() {
	input := `
	five
	`

	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
