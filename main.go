package main

import (
	"fmt"
	"interpretor/lexer"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./test.ts")
	if err != nil {

	}
	var lex = lexer.New(string(input))
	for token := lex.NextToken(); token.Type != "EOF"; token = lex.NextToken() {
		fmt.Print(token)
		fmt.Print("\n")
	}
}
