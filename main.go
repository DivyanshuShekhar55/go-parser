package main

import (
	"fmt"

	lexer "github.com/DivyanshuShekhar55/go-lexer"
)

func main(){
	fmt.Println("hello")
	source := "jjdj"
	tokens:= lexer.Tokenise(source)
	for _, item := range tokens{
		fmt.Println(item)
	}
}