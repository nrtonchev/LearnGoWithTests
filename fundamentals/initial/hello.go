package main

import (
	"fmt"
)

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "spanish"
const frenchHelloPrefix = "Bonjour, "
const french = "french"

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	if lang == spanish {
		return spanishHelloPrefix + name
	}

	if lang == french {
		return frenchHelloPrefix + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
