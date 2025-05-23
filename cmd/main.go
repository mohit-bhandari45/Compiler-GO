package main

import (
	"fmt"
	"github.com/mohit-bhandari45/Compiler-GO/lexer"
)

func main() {
	input := `let five = "Mohit";`

	l := lexer.New(input);
	for tok := l.NextToken(); tok.Type!=lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok);
	}
}
