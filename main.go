package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	source, err := ioutil.ReadFile("program.snk")

	if err != nil {
		log.Fatal(err)
	}

	input := string(source)

	tokens := Tokenize(input)

	for _, token := range tokens {
		fmt.Printf("%v -- %v: %v\n", token.Type, token.Name, token.Value)
	}
}
